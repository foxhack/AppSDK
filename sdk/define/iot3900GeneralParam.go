package define
/********************* 互联网部规范 2020.01 *********************/
/*****************************************
A.10 文件信息字段:file
file字段为object，描述一个文件的基本信息，其格式如下表A.10所示：
表 A.10 file 字段说明
字段 类型 是否必选 描述
name string 是 文件的名字
url string 否
文件的下载路径，即 url 地址，如
“https://123.123.123.123:6789”
size number 是 文件的大小，单位：k Bytes
md5 string 是 文件的 md5 值，用于校验文件
///sign object 否 文件的数字签证信息
*****************************************/
type IOTFileDefine struct {
	Name               string                          `json:"name"`
	Url                string                          `json:"url"`
	Size               int64                           `json:"size"`
	Md5                string                          `json:"md5"`
	//Sign               IOTFileSign                     `json:"sign"`
}

/*****************************************
A.11 数字签证信息字段:sign
表 A.11 sign 字段说明
字段 类型 是否必选 描述
name string 是 数字证书文件名字
url string 否
数字证书文件下载路径，即 url 地址，如
“https://123.123.123.123:6789”
md5 string 否 数字证书文件的 md5 值，用于校验文件
*****************************************/
type IOTFileSign struct {
	Name               string                          `json:"name"`
	Url                string                          `json:"url"`
	Md5                string                          `json:"md5"`
}

/*********************************
A.7 cpu 阈值信息字段:cpuLmt
参数cpuLmt给出了与cpu资源配置相关的参数，如下表A.7所示：
表 A.7 cpu 字段说明
字段 类型 是否必选 描述
cpus number 是 CPU 核数(例如值为 2，3，4)
Lmt number 是 CPU 监控阈值
*********************************/

type IOTAppCpuDefine struct {
	Cpus               int64                           `json:"cpus"`
	CpuLmt             int64                           `json:"cpuLmt"`
}

/*********************************
A.8 内存阈值信息字段:memLmt
参数memLmt给出了与memory资源配置相关的参数，如下表A.8所示：
表 A.8 mem 字段说明
字段 类型 是否必选 描述
unit number 是 内存限值,单位：M byte
Lmt number 是 内存监控阈值，百分数
*********************************/

/*********************************
A.9 硬盘阈值信息字段:diskLmt
参数diskLmt给出了与存储disk资源配置相关的参数，如下表A.9所示：
表 A.9 disk 字段说明
字段 类型 是否必选 描述
unit number 是 存储限值，单位：M byte
Lmt number 是 磁盘存储监控阈值，百分数
*********************************/
type IOTAppMemDskDefine struct {
	Memory             int64                           `json:"memory"`
	MemLmt             int64                           `json:"memLmt"`
}


/*********************************
A.1 设备信息字段：dev
dev字段的object中表示设备的基本信息，其各成员定义如下表A.1所示：
表 A.1 dev 字段说明
字段 类型 是否必选 描述
devSN string 是 边设备序列号
devType string 是 边设备类型
devName string 是 边设备名称
mfgInfo string 是 边设备厂商信息
devStatus string 是 边设备状态
*********************************/

type IOTDevDefine struct {
	DevSN               string                           `json:"devSN"`
	DevType             string                           `json:"devType"`
	DevName             string                           `json:"devName"`
	MfgInfo             string                           `json:"mfgInfo"`
	DevStatus           string                           `json:"devStatus"`
}

