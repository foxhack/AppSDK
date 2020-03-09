package define

import (
	"time"
)

/*****************************************
mid Number 是 Message_id,在请求报文中该值为请求 ID；在应答报文中，
该值为应答所对应请求报文的 ID；
deviceId String 否 接收此消息的设备 ID，批处理设备时该字段无需携带
timestamp Number 是 消息发送的时间戳，UTC 时间,精度到秒
expire Number 否
此报文过期的相对时间（相对 timestamp），单位秒，不包
含该字段，或者该字段为-1 时表示永不过期；
type String 是 消息类型，全大写，单词中间用“_”隔开
param Object 否 报文内容，详见各报文定义

*****************************************/
type IOTReqMsgHeader struct {
	Mid         int64                       `json:"mid"`
	DeviceId    string                      `json:"deviceId"`
	Timestamp   int64                       `json:"timestamp"`
	Expire      int                         `json:"expire"`
	Type        string                      `json:"type"`
	Param       interface{}                 `json:"param"`
}

/*****************************************
mid Number 是 在请求报文中该值为请求 ID；在应答报文中，该值为应答
所对应请求报文的 ID；
deviceId String 否 发送此消息的设备 ID；
timestamp Number 是 消息发送的时间戳，UTC 时间，单位秒
type String 是 消息类型,全大写，单词中间用“_”隔开
param Object 否 报文内容，详见各报文定义
code Number 是 标识应答的返回码
msg String 否 应答结果描述，字符串，最大 256 字符。
*****************************************/
type IOTAckMsgHeader struct {
	Mid         int64                       `json:"mid"`
	DeviceId    string                      `json:"deviceId"`
	Timestamp   int64                       `json:"timestamp"`
	Type        string                      `json:"type"`
	Param       interface{}                 `json:"param"`
	Code        int                         `json:"code"`
	Msg         string                      `json:"msg"`
}

func NewIotReqTop(MsgId int64, DeviceId string, MsgType string)IOTReqMsgHeader{
	reqMsg := IOTReqMsgHeader{}
	reqMsg.Mid = MsgId
	reqMsg.DeviceId = DeviceId
	reqMsg.Type = MsgType  //"EVENT_DEVICE_ADD"
	reqMsg.Expire = -1
	reqMsg.Timestamp = time.Now().UnixNano() / 1e6  ///time.Now().Unix()
	///reqMsg.Param = MsgBody   //// interface{}, 使用时 直接 赋值

	return reqMsg
}

func NewIotAckTop(MsgId int64, DeviceId string, MsgType string, Code int, Msg string)IOTAckMsgHeader{
	reqMsg := IOTAckMsgHeader{}
	reqMsg.Mid = MsgId
	reqMsg.DeviceId = DeviceId
	reqMsg.Type = MsgType  //"EVENT_DEVICE_ADD"
	reqMsg.Timestamp = time.Now().Unix()
	reqMsg.Code = Code
	reqMsg.Msg = Msg
	///reqMsg.Param = MsgBody   //// interface{}, 使用时 直接 赋值

	return reqMsg
}
