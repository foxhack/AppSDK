package define

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

type EdgeRegMsgHead3900 struct{
	Id          int64                 `json:"id"`
	Version     string                `json:"version"`
	Function    string                `json:"function"`
	Caller      string                `json:"caller"`
	Body        EdgeRegMsgBody3900        `json:"body"`
}

type EdgeRegMsgBody3900 struct {
	DevSN       string                `json:"sn"`
	Model       string                `json:"module"`
	Factory     string                `json:"manufacture"`
	AlgId       int                   `json:"algId"`
	CheckID     string                `json:"checkID"`
	///HttpPath string  `json:"path"`   /// 测试用， 平台接入接口全地址
	//// ==== NS3910 需要带地址和端口
	///EdgeAddr    string                `json:"edgeAddr"`
	///EdgePort    int                   `json:"edgePort"`
}


/*********************** 网关接入返回信息 *******************/
type EdgeRegAckHead3900 struct {
	Code       int                   `json:"code"`
	ErrMsg     string                `json:"errMsg"`
	Id         string                `json:"id"`
	Value      EdgeRegInfo3900       `json:"value"`
}




