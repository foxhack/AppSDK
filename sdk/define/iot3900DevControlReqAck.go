package define
/********************* 互联网部规范 2020.01 *********************/
/*****************************************
45、设备控制命令    	平台到设备   	topic:/v1/${edgeId}/device/command
	内容:	{
		"deviceId": "1577123434420863986",
		"type": "CMD_CTRL",
		"expire": -1,
		"mid": 667645741,
		"param": {
			"action": "os-reboot"
		},
		"timestamp": 1577106555
	}
*****************************************/
type IOTDevControlReq struct {
	Action              string                         `json:"action"`
}

/*****************************************
46、设备控制命令应答   设备到平台    	topic:/v1/${edgeId}/device/reply
	内容:	{
		"mid": 1569582513983943814,
		"deviceId": "1123123",
		"type": "CMD_CTRL",
		"timestamp": 1562746993127,
		"code": "200",
		"msg": "SUCCESS"
	}
    IOTAckMsgHeader
 *****************************************/

/*********************************
47、设备webSSh命令-	平台到设备 ------客户新增协议------   	topic：/v1/${edgeId}/device/command
	内容：  	{
		"deviceId": "1577123434420863986",
		"type": "CMD_WEBSSH",
		"expire": -1,
		"mid": 667645741,
		"param": {
			"webSshIp": "",
			"webSshPort": "",
			"webSshName": "",
			"webSshPassword": "",
			"localPort": "",
		},
		"timestamp": 1577106555
	}
*********************************/
type IOTDevWebSShConfig struct {
	WebSshIp           string                          `json:"webSshIp"`
	WebSshPort         string                          `json:"webSshPort"`
	WebSshName         string                          `json:"webSshName"`
	WebSshPassword     string                          `json:"webSshPassword"`
	LocalPort          string                          `json:"localPort"`
}

/*********************************
48、设备webSSh命令回应		 设备到平台
	topic：/v1/${edgeId}/device/reply
	内容：
	{
		"mid": 1569582513983943814,
		"deviceId": "1123123",
		"type": "CMD_WEBSSH",
		"timestamp": 1562746993127,
		"code": "200",
		"msg": "SUCCESS"
	}
    IOTAckMsgHeader
*********************************/


/*********************************

*********************************/


/*********************************

*********************************/


/*********************************

*********************************/
