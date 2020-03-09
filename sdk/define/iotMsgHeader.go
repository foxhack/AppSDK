package define

import (
	"fmt"
	"time"
)

/*********************************
*********************************/

type IOTReqMsgHeader_DevReg struct {
	Mid         int64                       `json:"mid"`
	DeviceId    string                      `json:"deviceId"`
	Timestamp   int64                       `json:"timestamp"`
	Expire      int                         `json:"expire"`
	Type        string                      `json:"type"`
	Param       DeviceRegMsgHead            `json:"param"`
}

func NewReqMsgHeader_DevReg(MsgId int64, edgeId int64, MsgType string, MsgBody DeviceRegMsgHead)IOTReqMsgHeader_DevReg{
	reqMsg := IOTReqMsgHeader_DevReg{}
	reqMsg.Mid = MsgId
	reqMsg.DeviceId = fmt.Sprintf("%v", edgeId)
	reqMsg.Type = MsgType  //"EVENT_DEVICE_ADD"
	reqMsg.Expire = -1
	reqMsg.Timestamp = time.Now().Unix()
	reqMsg.Param = MsgBody

	return reqMsg
}

/*********************************
*********************************/
type IOTReqMsgHeader_OnOffStatus struct {
	Mid         int64                       `json:"mid"`
	DeviceId    string                         `json:"deviceId"`
	Timestamp   int64                       `json:"timestamp"`
	Expire      int                         `json:"expire"`
	Type        string                      `json:"type"`
	Param       DeviceOnlineStatusBody      `json:"param"`
}

func NewReqMsgHeader_OnOffStatus(MsgId int64, edgeId string, MsgType string, MsgBody DeviceOnlineStatusBody)IOTReqMsgHeader_OnOffStatus{
	reqMsg := IOTReqMsgHeader_OnOffStatus{}
	reqMsg.Mid = MsgId
	reqMsg.DeviceId = fmt.Sprintf("%v", edgeId) /// edgeId
	reqMsg.Type = MsgType  //"EVENT_DEVICE_ADD"
	reqMsg.Expire = -1
	reqMsg.Timestamp = time.Now().Unix()
	reqMsg.Param = MsgBody

	return reqMsg
}


/*********************************
*********************************/
type IOTReqMsgHeader_Headbeat struct {
	Mid         int64                       `json:"mid"`
	DeviceId    string                      `json:"deviceId"`
	Timestamp   int64                       `json:"timestamp"`
	Expire      int                         `json:"expire"`
	Type        string                      `json:"type"`
	Param       EdgeHeartbeatAckBody      `json:"param"`
}

func NewReqMsgHeader_Headbeat(MsgId int64, edgeId int64, MsgType string, MsgBody EdgeHeartbeatAckBody)IOTReqMsgHeader_Headbeat{
	reqMsg := IOTReqMsgHeader_Headbeat{}
	reqMsg.Mid = MsgId
	reqMsg.DeviceId = fmt.Sprintf("%v", edgeId) /// edgeId
	reqMsg.Type = MsgType  //"EVENT_DEVICE_ADD"
	reqMsg.Expire = -1
	reqMsg.Timestamp = time.Now().Unix()
	reqMsg.Param = MsgBody

	return reqMsg
}

/*********************************
*********************************/

type IOTAckMsgHeader_devAuth struct {
	Mid         int64                       `json:"mid"`
	DeviceId    string                      `json:"deviceId"`
	Timestamp   int64                       `json:"timestamp"`
	Type        string                      `json:"type"`
	Param       DeviceRegAckHead            `json:"param"`
	Code        int                         `json:"code"`
	Msg         string                      `json:"msg"`
}


func NewAckMsgHeader_devAuth(MsgId int64, devId string, MsgType string)IOTAckMsgHeader_devAuth{
	reqMsg := IOTAckMsgHeader_devAuth{}
	reqMsg.Mid = MsgId
	reqMsg.DeviceId = devId
	reqMsg.Type = MsgType  //"EVENT_DEVICE_ADD"
	reqMsg.Msg = ""
	reqMsg.Timestamp = time.Now().Unix()

	return reqMsg
}

