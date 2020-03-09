package define

import "github.com/edgexfoundry/nsplussdk/sdk/nrprivate"

const(
	EngineConfFILE          = "engineconf.json"
)
////======== 数据发布引擎配置文件 ========
type IOTAuthInterfaceDefine struct{
	//// 引擎名称， 用于标识配置文件，不可重复
	Name             string               `json:"name"`  /// === 字符串加数字，全局唯一关键字
	Description      string               `json:"description"`
	///  对应不同的数据发布引擎， 1、NS3800 二次设备维护平台， 2、IOT 主站平台
	IotType          string               `json:"iotType"`
	/// 接入设置， 说明接入的api
	// === 创建数据发布模板必须配置接入参数，以配合数据接入微服务；
	// === 修改操作、删除操作考虑做通知，告知接入微服务重新加载参数或停止数据接入
	AuthConfig       AuthConfigDefine     `json:"authConfig"`

	///  新增IOT模板时，默认置 0 ，不生效， 通过手动连接命令 改为1 生效
	EnableFlag       string               `json:"IsValid"`

	/// IOT参数设置， 分网关接入 和 设备接入， EdgeSN 作为全局变量提取
	EdgeConfig       nrprivate.AllConfigAckDefine   `json:"edgeConfig"`

	/// 数据过滤设置 === 数据发布时可用
	DataFilter       DataFilterDefine     `json:"dataFilter"`
	/// 数据转换定义 === 数据发布时可用
	DataTrans        DataTransConfig      `json:"dataTrans"`
}
////======== 数据发布引擎配置文件  ========


//// ======== 接入证书信息定义 ========
type CertificatesDefine struct {
	Cert string `json:"cert"` //// Cert = 'dummy.crt'
	Key  string `json:"key"`  ///  Key = 'dummy.key'
}
//// 考虑兼容 https和mqtt两种接入接口定义
type AuthInterfaceDefine struct {
	Protocol string             `json:"protocol"`
	Host     string             `json:"host"`
	Port     int                `json:"port"`
	Path     string             `json:"path"`
	Topic    string             `json:"topic"`
	UserName string             `json:"userName"`
	Password string             `json:"password"`
	Cert     CertificatesDefine `json:"cert"`     /// 加密证书
	CheckKey string             `json:"checkKey"` /// 接入的关键字，设备key之类的,备用，默认0
}
/// 接入区分 边和 端 分别定义
type AuthConfigDefine struct{
	/// 接入相关配置
	EdgeApi       AuthInterfaceDefine    `json:"edgeApi"`
	DeviceApi     AuthInterfaceDefine    `json:"deviceApi"`
}

//// ======== 数据过滤定义 ========
///  如果都设置了，则先按量测类型过滤，再按数据类型过滤；
/// ======== 通用过滤设置，暂定 过滤按 1-量测类型 、2-资源类型、3-profile名称（全部资源）、4-设备名（全部资源） 四种模式：========
type GeneralFilterDefine struct{
	Type              string                 `json:"type"`
	Key               map[string]string      `json:"key"`
}

////======== 详细过滤类型 type 10-按 profile + 资源名, 20-按 device + 资源名 ========
type DetailFilterDefine struct{
	Type              string                            `json:"type"` ///
	Name              string                            `json:"name"` /// 设备或profilename
	Key               map[string]string                 `json:"key"`  /// 资源名列表
}

type DataFilterDefine struct{
	GeneralFilter     map[string]GeneralFilterDefine    `json:"generalFilter"`
	DetailFilter      map[string]DetailFilterDefine     `json:"detailFilter"`
}

//// ======== 数据转换定义 ========
//// 包括 网关类型 到 平台类型的转换定义， 资源名称 到平台名称的转换定义
type DataTransConfig struct{
	/// 数据类型的转换，网关类型和平台类型定义不同，可配置转换，key为网关类型，value为平台类型
	Types       map[string]string                      `json:"types"`
	/// 数据名称的转换，按平台要求的名称进行转换，可配置按profile-资源名称转换或device-资源名称装换，
	// key为网关类型，value为平台类型
	Names       map[string]DetailFilterDefine          `json:"names"`
}


