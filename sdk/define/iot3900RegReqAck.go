package define

import "time"

/********************* 互联网部规范 2020.01 *********************/
/******************************************************************
1、边缘设备接入请求:	设备到平台 	topic:/v1/${edgeId}/device/request
	内容:
	{
	"mid":111111111,
	"type ":"EVENT_LINKUP",
	"timestamp":1574930415425,
	"deviceId":1576633766105801053,
	"param":{
			"dev":{
				"devSN":"00000000002121",
				"devType":"TTU",
				"devName":"TTU01",
				"mfgInfo":"NARI",
				"devStatus":"ONLINE"
			},
			"cpu":{
				"cpus":8,
				"frequency":2.94,
				"cache":1024,
				"arch":"",
				"cpuLmt":60
			},
			"mem":{
				"phy":16,
				"virt":2,
				"memLmt":50
			},
			"disk":{
				"volume":1024,
				"volumeLmt":60
			},
			"os":{
				"distro":"Ubunut",
				"version":"18.10",
				"kernel":"3.10-17",
				"softVersion":"V01.024"
			},
			"links"[
				{
					"type":"Ethernet",
					"id":"000000000211",
					"name":"eth1",
					"mac":"B8-85-74-15-A5-3E"
				},{
					"type":"4G",
					"id":"000000000212",
					"name":"ppp-0",
					"mac":"B8-85-74-15-A5-3D"
				}
			]
		}
	}
******************************************************************/

type EdgeRegMsgPara3900 struct{
	Dev                    EdgeDevInfo   `json:"dev"`
	Cpu                    CpuInfo       `json:"cpu"`
	Mem                    MemInfo       `json:"mem"`
	Disk                   DiskInfo      `json:"disk"`
	Os                     OsInfo        `json:"os"`
	Links                  []LinkInfo    `json:"links"`
}

type EdgeDevInfo struct {
	DevSN                  string        `json:"devSN"`
	DevType                string        `json:"devType"`
	DevName                string        `json:"devName"`
	MfgInfo                string        `json:"mfgInfo"`
	DevStatus              string        `json:"devStatus"`
}
type CpuInfo struct {
	Cpus                   int           `json:"cpus"`
	Frequency              float32       `json:"frequency"`
	Cache                  int           `json:"cache"`
	Arch                   string        `json:"arch"`
	CpuLmt                 string        `json:"cpuLmt"`
}
type MemInfo struct {
	Phy                    int           `json:"phy"`
	Virt                   int           `json:"virt"`
	MemLmt                 int           `json:"memLmt"`
}

type DiskInfo struct {
	Volume                 int           `json:"volume"`
	VolumeLmt              int           `json:"volumeLmt"`
}
type OsInfo struct {
	Distro                 string        `json:"distro"`   /// centos7
	Version                string        `json:"version"`
	Kernel                 string        `json:"kernel"`
	SoftVersion            string        `json:"softVersion"`
}
type LinkInfo struct {
	Type                   string        `json:"type"`
	Id                     string        `json:"id"`
	Name                   string        `json:"name"`
	Mac                    string        `json:"mac"`
}

/********************************************
2、设备接入应答:  平台到设备 	topic:/v1/${edgeId}/device/response
	内容:
	{
		"mid":111111111,
		"type ":"EVENT_ LINKUP",
		"timestamp":1574930415425,
		"deviceId":1576633766105801053,
		"code": 200,
		"msg": "SUCCESS",
		"param":{}
	}
 *******************************************/


/********************************************
3、子设备接入认证（子设备添加问题）  设备到平台 	topic:/v1/${edgeId}/topo/request
	内容:	{
		"mid":123456,
		"deviceId":"415211",
		"timestamp":15212345678900,
		"type":"CMD_TOPO_ADD",
		"param":{
			"nodeInfos":[
				{
					"nodeId":"testSN001",
					"name":"test001",
					"description":"test",
					"mfgInfo":"NARI",
					"nodeModel":"nari"
				}
			]
		}
	}
 *******************************************/

type DevRegMsgPara3900 struct{
	NodeInfos              []DevRegInfos `json:"nodeInfos"`
}

type DevRegInfos struct {
	NodeId                 string        `json:"nodeId"`   /// DevSN
	Name                   string        `json:"name"`
	Description            string        `json:"description"`
	MfgInfo                string        `json:"mfgInfo"`
	NodeModel              string        `json:"nodeModel"`
}

/************************************************
	4、子设备接入应答  平台到设备
	topic:/v1/${edgeId}/topo/response
	内容:
	{
		"mid":123456,
		"deviceId":"415211",
		"timestamp":15212345678900,
		"type":"CMD_TOPO_ADD",
		"code":200,
		"msg":"SUCCESS!",
		"param":{
			"result":[
				{
					"statusCode":200,
					"statusDesc":"SUCCESS!",
					"nodeId":"testSN001",
					"deviceId":"152223495866321"
				},
				{
					"statusCode":200,
					"statusDesc":"SUCCESS!",
					"nodeId":"testSN002",
					"deviceId":"152223495866322"
				}
			]
		}
	}
 ***************************************************/
type IOTDevRegAckPara3900 struct {
	Mid         int64                       `json:"mid"`
	DeviceId    string                      `json:"deviceId"`
	Timestamp   int64                       `json:"timestamp"`
	Type        string                      `json:"type"`
	Param       DevRegAckPara3900           `json:"param"`
	Code        int                         `json:"code"`
	Msg         string                      `json:"msg"`
}

type DevRegAckPara3900 struct{
	Result                 []DevRegAck   `json:"result"`
}

type DevRegAck struct {
	StatusCode             int           `json:"statusCode"`
	StatusDesc             string        `json:"statusDesc"`
	NodeId                 string        `json:"nodeId"`   /// DevSN
	DeviceId               string        `json:"deviceId"`
}

/************************************************
	5、子设备删除 设备到平台
	topic:/v1/${edgeId}/topo/request
	内容:
	{
		"mid":123456,
		"deviceId":"415211",
		"timestamp":15212345678900,
		"type":"CMD_TOPO_DEL",
		"param":{
			"nodeIds":["152223495866321","152223495866322"]
		}
	}
 ***************************************************/

type DevDelMsgPara3900 struct{
	NodeIds                []string      `json:"nodeIds"`
}

/************************************************
   6、子设备删除应答 平台到设备 	topic:/v1/${edgeId}/topo/response
	内容:
	{
		"mid":123456,
		"deviceId":"415211",
		"timestamp":15212345678900,
		"type":"CMD_TOPO_DEL",
		"code":200,
		"msg":"SUCCESS!",
		"param":{
			"result":[
				{
					"statusCode":200,
					"statusDesc":"SUCCESS!",
					"deviceId":"152223495866321"
				},
				{
					"statusCode":200,
					"statusDesc":"SUCCESS!",
					"deviceId":"152223495866322"
				}
			]
		}
	}
***************************************************/
type DevDelAckPara3900 struct{
	Result                 []DevOprAck   `json:"result"`
}

type DevOprAck struct {
	StatusCode             int           `json:"statusCode"`
	StatusDesc             string        `json:"statusDesc"`
	//NodeId                 string        `json:"nodeId"`   /// DevSN
	DeviceId               string        `json:"deviceId"`
}

/************************************************
7、子设备状态更新 设备到平台   	topic:/v1/${edgeId}/topo/request
	内容:
	{
		"mid":123456,
		"deviceId":"415211",
		"timestamp":15212345678900,
		"type":"CMD_TOPO_UPDATE",
		"param":{
			"nodeStatuses":[
  				{
  					"deviceId":"152223495866321",
  					"status":"ONLINE"
  				},
  				{
  					"deviceId":"152223495866322",
  					"status":"OFFLINE"
  				}
  			]
		}
	}
 ***************************************************/
type IOTReqDevUpdateMsgPara3900 struct {
	Mid         int64                       `json:"mid"`
	DeviceId    string                      `json:"deviceId"`
	Timestamp   int64                       `json:"timestamp"`
	Expire      int                         `json:"expire"`
	Type        string                      `json:"type"`
	Param       DevUpdateMsgPara3900        `json:"param"`
}

type DevUpdateMsgPara3900 struct{
	NodeStatuses           []DevUpdateMsg `json:"nodeStatuses"`
}

type DevUpdateMsg struct {
	DeviceId               string         `json:"deviceId"`
	Status                 string         `json:"status"`
}


func NewIotReqDevUpdateMsgPara3900(MsgId int64, DeviceId string, MsgType string)IOTReqDevUpdateMsgPara3900{
	reqMsg := IOTReqDevUpdateMsgPara3900{}
	reqMsg.Mid = MsgId
	reqMsg.DeviceId = DeviceId
	reqMsg.Type = MsgType  //"EVENT_DEVICE_ADD"
	reqMsg.Expire = -1
	reqMsg.Timestamp = time.Now().UnixNano() / 1e6  ///time.Now().Unix()
	///reqMsg.Param = MsgBody   //// interface{}, 使用时 直接 赋值

	return reqMsg
}

/************************************************
8、子设备状态更新应答 平台到设备 	topic:/v1/${edgeId}/topo/response
	内容:
	{
		"mid":123456,
		"deviceId":"415211",
		"timestamp":15212345678900,
		"type":"CMD_TOPO_UPDATE",
		"code":200,
		"msg":"SUCCESS!",
		"param":{
			"result":[
				{
					"statusCode":200,
					"statusDesc":"SUCCESS!",
					"deviceId":"152223495866321"
				},
				{
					"statusCode":200,
					"statusDesc":"SUCCESS!",
					"deviceId":"152223495866322"
				}
			]
		}
	}
***************************************************/
type DevUpdateAckPara3900 struct{
	Result                 []DevOprAck    `json:"result"`
}

/************************************************
9、设备事件主动上报  设备到平台			//协议中未指定     	topic:/v1/{edgeId}/service/data
	内容:
	{
		"mid":111111111,
		"type ":"EVENT_DATA_ALARM",		//协议中未指定
		"timestamp":1574930415425,
		"deviceId":1576633766105801053,
		"param":{
				"deviceId":1576633766105801053,
				"cmd":"B_phs_alert",		//数据转发使用
				"data":{
					"B_phs":"12",
					"B_phs_alert_val":"34"
				}
		}
	}
***************************************************/


/************************************************
10、设备数据主动上报   设备到平台  	topic:/v1/{edgeId}/service/data
	内容:
	{
		"mid":111111111,
		"type ":"CMD_REPORTDATA",
		"timestamp":1574930415425,
		"deviceId":1576633766105801053,
		"param":{
			"deviceId":1576633766105801053,
			"cmd":"B_phs_alert",		//数据转发使用    //协议中未定义
			"data":{
					"Ua":220.3,
					"Ub":221,
					"Ia":23.4    //属性  key-value结构
			 }
		}
	}
***************************************************/
type DevDataReportMsgPara3900 struct{
	DeviceId               string         `json:"deviceId"`
	Cmd                    string         `json:"cmd"`
	Data                   map[string]interface{}   `json:"data"`
}

/************************************************
11、业务命令下发 平台到设备 	topic:/v1/${edgeId}/service/command
	内容:
	{
		"mid":111111111,
		"type ":"CMD_SERVICE",
		"timestamp":1574930415425,
		"param":{
			"cmd":"analog_Get",
			"paras":{
				"body":
				{
					"PhV_phsA":"",
					"PhV_phsB":""
				}
			}
		}
	}
***************************************************/

/************************************************
12、业务命令响应:	设备到平台 	topic:/v1/${edgeId}/service/reply
	内容:
	{
		"mid":111111111,
		"type ":"CMD_SERVICE",
		"timestamp":1574930415425,
		"deviceId":4217058101079148192,
		"code": 200,
		"msg": "SUCCESS",
		"param":{
			"cmd":"analog_Get",
			"paras":{
				"body":
				{
					"PhV_phsA":"220.4",
					"PhV_phsB":"220.7"
				}
			}
		}
	}
***************************************************/
type PDZCmdMsgPara3900 struct{
	DeviceId               string         `json:"deviceId"`
	Cmd                    string         `json:"cmd"`
	paras                  ParasInfo      `json:"paras"`
}
type ParasInfo struct{
	Body                   map[string]interface{}   `json:"body"`
}


/************************************************
13、物模型下发 平台到设备
	topic:/v1/${edgeId}/service/command
	内容：
	{
	"deviceId": "1577063490614240253",
	"type": "CMD_PROFILE",
	"expire": -1,
	"mid": 1569582513983943814,
	"param": {
		deviceIds:["1577063490614240253","1577063490614240253"],
		"file": {
			"url": "http://26.47.73.135:23004/v2/iot/files/23010/up/4404ABE3-A0A6-4C91-A954-62C981813E19.json",
			"size": 4398,
			"name": "TTU_V2.1_1577066057788.json",
			"md5": "fdecc0af975477cf363295f8717475e3"
		}
	},
	"timestamp": 1577094947
}
***************************************************/
type DevProfileDownMsgPara3900 struct{
	DeviceIds              []string       `json:"deviceIds"`
	File                   FileInfo       `json:"file"`
}

/************************************************
14、物模型下发响应 	设备到平台
	topic:/v1/${edgeId}/service/reply
	内容:
	{
		"mid": 1569582513983943814,
		"deviceId": "1577063490614240253",
		"type": "CMD_PROFILE",
		"timestamp": 1577065651000,
		"code": 200,
		"msg": "SUCCESS",
		"param":{}
	}
***************************************************/


/************************************************

15、固件下发（设备升级）流程    平台到设备 	topic:/v1/${edgeId}/device/command
	内容：
	{
		"deviceId": "1577063490614240253",
		"type": "CMD_SYS_UPGRADE",
		"expire": -1,
		"mid": 1569582513983943815,
		"param": {
			"jobId": 1569582513983943816,
			"version": "V1.2",
			"policy": 1,
			"upgradeType": 0,
			"file": {
				"url": "http://26.47.73.135:23004/v2/iot/files/1231/up/C8849216-A006-4CED-B9BF-C369538B9848.tar",
				"size": 22,
				"name": "固件1",
				"md5": "76cdb2bad9582d23c1f6f4d868218d6c"
			}
		},
		"timestamp": 1577101833
	}
***************************************************/
type DevFirmDownMsgPara3900 struct{
	JobId                  int64          `json:"jobId"`
	Version                string         `json:"version"`
	Policy                 int            `json:"policy"`
	UpgradeType            int            `json:"upgradeType"`
	File                   FileInfo       `json:"file"`
}


/************************************************
16、固件下发（设备升级）命令响应  设备到平台

	topic:/v1/${edgeId}/device/reply
	内容:
	{
		"mid": 1569582513983943815,
		"deviceId": "1577063490614240253",
		"type": "CMD_SYS_UPGRADE",
		"timestamp": 1577073563000,
		"code": 202,
		"msg": "SUCCESS",
		"param":{}
	}
***************************************************/


/************************************************
17、固件下发（设备升级）结果上报  设备到平台 	topic:/v1/${edgeId}/device/data
	内容:
	{
		"mid": 1569582513983943815,
		"deviceId": "1577063490614240253",
		"type": "REP_JOB_RESULT",
		"timestamp": 1577065651000,
		"code": 200,
		"param": {
			"jobId": 1569582513983943816,
			"result": 200,
			"info": "SUCCESS"
		},
		"msg": "SUCCESS"
	}
***************************************************/
type DevFirmDownAckPara3900 struct{
	JobId                  int64          `json:"jobId"`
	Result                 int            `json:"result"`
	Info                   string         `json:"upgradeType"`
}


