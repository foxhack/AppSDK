package define


type EdgeAuthInfo struct {
	DevSN       string                       `json:"devSn"`
	EdgeId      int64                          `json:"edgeId"`

	Devices     map[string]DeviceAuthInfo    `json:"devices"`

	AppList     map[string]string            `json:"appList"`
}


type DeviceAuthInfo struct {
	DevId       string                       `json:"devId"`

	OnOff       int                          `json:"onOff"`
	ForceOff    bool                         `json:"forceOff"`
	ServiceInfo DevServiceInfo               `json:"serviceInfo"`
	//DevSN       string                       `json:"devSn"`
	//Manufacture string                       `json:"manufacture"`
	//Model       string                       `json:"model"`
	//Name        string                       `json:"name"`

	DevInfo     DeviceRegMsgBody             `json:"devInfo"`
}



type DevServiceInfo struct{
	Name           string           `json:"name"`          // Unique name for identifying a device ， use as deviceSN
	///LastConnected  int64         `json:"lastConnected"` // Time (milliseconds) that the device last provided any feedback or responded to any request
	///LastReported   int64         `json:"lastReported"`  // Time (milliseconds) that the device reported data to the core microservice
	//Service        DeviceService  `json:"service"`       // Associated Device Service - One per device
	Address        string            `json:"address"`      // Address of the addressable  "address": "192.168.43.150",
	Port           int               `json:"port,Number"`  // Port for the address        "port": 49982,
	//Profile      DeviceProfile     `json:"profile"`      // Associated Device Profile - Describes the device
	ProfileName    string            `json:"profilename"`  // Non-database identifier (must be unique)

	///ResourceNameType   []ResourceNameTypeInfo   `json:"resourcenametype"`
}


/******************************************************************
Url: /v2/iot/devices/direct/auth
Method: POST
Request:
参数	必选	类型	说明
id	是	long	请求和应答一一对应，便于调用者分析请求和响应关系
version	是	String	对REST接口的版本限制
function	是	String	调用Controller里面的那个函数名称
caller	是	String	调用发起者的名称
body	是	JSONObject	请求报文体内容

body 定义
参数	必选	类型	说明
devSn	是	String	设备唯一标识
model	否	String	设备型号
factory	否	String	设备厂商
algId	否	String	算法ID，表示使用哪个算法加密
checkID	否	String	检测值，该值为使用algId算法以及设备的密钥对SN加密后的字符串

******************************************************************/

type EdgeRegMsgHead struct{
	Id          int64                 `json:"id"`
	Version     string                `json:"version"`
	Function    string                `json:"function"`
	Caller      string                `json:"caller"`
	Body        EdgeRegMsgBody        `json:"body"`
}

type EdgeRegMsgBody struct {
	DevSN       string                `json:"devSn"`
	Model       string                `json:"model"`
	Factory     string                `json:"factory"`
	AlgId       string                   `json:"algId"`
	CheckID     string                `json:"checkID"`
	///HttpPath string  `json:"path"`   /// 测试用， 平台接入接口全地址
	//// ==== NS3910 需要带地址和端口
	EdgeAddr    string                `json:"edgeAddr"`
	EdgePort    int                   `json:"edgePort"`
}

type EdgeRegAckHead struct{
	Code        int                   `json:"code"`
	ErrMsg      string                `json:"errMsg"`
	EdgeId      int64                   `json:"edgeId"`
}


/******************************************************************
Cer String 否 证书请求文件（后期迭代版本考虑增加证书加密接入）
devices []Devices 是  设备 Sn 列表
Devices 定义如下： 表 3-9 Devices 对象定义 字段 类型 是否必选 描述
sn String 是 设备的唯一身份标识
manufacture String 是 设备供应商（供应商字段长度最长不超过 10
字符，字母开头，可包含数字，不支持特殊
字符）
module String 是 设备型号
name String 否 设备名字，非必选
******************************************************************/

type DeviceRegMsgHead struct {
	Cer          string               `json:"Cer"`   /// 暂定为空

	Devices      []DeviceRegMsgBody   `json:"devices"`

	EdgeSN       string                  `json:"edgeSN"`
}
type DeviceRegMsgBody struct {
	Sn          string                `json:"sn"`
	Manufacture string                `json:"manufacture"`
	Module      string                `json:"module"`
	Name        string                `json:"name"`
}


type DeviceRegAckHead struct{
	Devices     []DeviceRegAckBody    `json:"devices"`
}
type DeviceRegAckBody struct {
	Sn          string                `json:"sn"`
	DevId       int64                 `json:"devId"`
	Code        int                   `json:"code"`
	Msg         string                `json:"msg"`
}


/******************************************************************
devId Number 是 主站为设备分配的唯一 ID
isOnline Number 是 设备上线或者下线,1 表示上线;0 表示下线
*******************************************************************/
type DeviceOnlineStatusBody struct {
	DevId           string                `json:"devId"`
	IsOnline        int64                 `json:"isOnline"`
}

/******************************************************************
edgeId Number 是 网关在主站平台的内部 Id
time String 是 当前时间的时间戳，以秒 s 为单位
*******************************************************************/
type EdgeHeartbeatAckBody struct {
	DevId       int64                   `json:"edgeId"`
	Time        int64                 `json:"time"`
}

/******************************************************************
edgeConfig	EdgeConfiguration	网关配置	（1）和（3）接口时有值
devConfig	[]DevConfiguration	设备配置	（2）（3）（4）接口时有值

其中， EdgeConfiguration为网关配置，定义如下：
表 3 2 EdgeConfiguration定义
字段名	类型	说明	备注
stationName	String	厂站名称
voltageLevel	String	厂站电压等级，单位kV
posLongitude	Number	厂站所在位置经度数值, float64	地图绘图
posLatitude	Number	厂站所在位置纬度数值, float64	地图绘图
edgeAddr	String	网关地址
edgePort	Number	网关端口,int
edgeSn	String	网关SN，唯一标识
edgeModel	String	网关型号
edgeManufacture	String	网关制造厂家
edgeLocation	Location	网关安装位置

其中， DevConfiguration为返回二次设备和边缘APP配置，定义如下：
表 3 3 DevConfiguration定义
字段名	类型	说明	备注
devSn	String	二次设备SN，唯一标识
devType	String	二次设备类型，测控装置、同步时钟、同步向量测量装置、相量数据集中器……
devModel	String	设备型号
devManufacture	String	设备制造厂家
devLocation	Location	设备安装位置
appBindDevice	String	APP关联的设备	对APP有效
capacity	Array	设备具备的能力列表，String数组，比如虚拟液晶显示能力-VD
devName	String	设备名称

其中Location定义如下：
表 3 4 Location 定义
字段名	类型	说明	备注
room	String	安装小室
cabinet	String	屏柜
rock	String	安装槽位
******************************************************************/
/*****************************************************
type EdgeDevicesConfiguration struct{
	//Sn          string                `json:"sn"`
	EdgeConfig	    EdgeConfiguration     `json:"edgeConfig"`
	DevConfig	    []DevConfiguration    `json:"devConfig"`
}
type EdgeConfiguration struct{
	//Sn          string                 `json:"sn"`
	StationName	    string                `json:"stationName"`
	VoltageLevel	string                `json:"voltageLevel"`
	PosLongitude	float64               `json:"posLongitude"`
	PosLatitude	    float64               `json:"posLatitude"`
	EdgeAddr	    string                `json:"edgeAddr"`
	EdgePort	    int                   `json:"edgePort"`
	EdgeSn	        string                `json:"edgeSn"`
	EdgeModel       string                `json:"edgeModel"`
	EdgeManufacture string                `json:"edgeManufacture"`
	EdgeLocation    DevLocation           `json:"edgeLocation"`
}
type DevConfiguration struct{
	//Sn          string                `json:"sn"`
	DevSn	        string                `json:"devSn"`  ///	二次设备SN，唯一标识
	DevType	        string                `json:"devType"`  ///	二次设备类型，测控装置、同步时钟、同步向量测量装置、相量数据集中器……
	DevModel	    string                `json:"devModel"`  ///	设备型号
	DevManufacture	string                `json:"devManufacture"`  ///	设备制造厂家
	DevLocation	    DevLocation           `json:"devLocation"`  ///	设备安装位置
	AppBindDevice	string                `json:"appBindDevice"`  ///	APP关联的设备	对APP有效
	Capacity	    []string              `json:"capacity"`  ///	设备具备的能力列表，String数组，比如虚拟液晶显示能力-VD
	DevName	        string                `json:"devName"`  ///	设备名称
}
type DevLocation struct{
	//Sn          string                `json:"sn"`
	Room	        string                `json:"room"`	///安装小室
	Cabinet	        string                `json:"cabinet"`	///屏柜
	Rock	        string                `json:"rock"`	///安装槽位
}
*****************************************************/


