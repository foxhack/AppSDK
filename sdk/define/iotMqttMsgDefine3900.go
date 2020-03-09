package define



type IOTMsgHeadTop struct{
	Mid          int64                `json:"mid"`
	Method       string               `json:"method"`
	DevId        string               `json:"devId"`
	Timestamp    int64                `json:"timestamp"`
	Body         interface{}          `json:"body"`
	Expire       int                  `json:"expire"`
}

/******************************************************************
设备-->平台
	topic:/v1/{edgeId}/topo/request
	内容:{
        		"method":"EVENT_ADD",
"mid": 1569582513983943814,
		        "expire":"-1",
        		"timestamp":1574930415426,
        		"body":{
            			"devices":[{
                			"sn":"sn011101",
                			"manufacture":"NARI",
               	 			"module":"M003"

           			 }{
                			"sn":"sn011102",
                			"manufacture":"NARI",
               	 			"module":"M004"
           			 }]
       			 }
    		}
******************************************************************/

type DeviceRegMsgHead3900 struct {
	Method          string               `json:"method"`
	Mid             int64                `json:"mid"`
	Expire          int                  `json:"expire"`
	Timestamp       int64                `json:"timestamp"`
	Body            DeviceRegMsg         `json:"body"`

}
type DeviceRegMsg struct {
	Devices          []DeviceRegBody     `json:"devices"`
}
type DeviceRegBody struct {
	Sn          string                `json:"sn"`
	Manufacture string                `json:"manufacture"`
	Module      string                `json:"module"`
	//Name        string                `json:"name"`
}

/*****************************
响应:平台-->设备
	topic:/v1/{edgeId}/topo/response
	内容:{
        		"method":"EVENT_ADD",
"mid": 1569582513983943814,
				"devId": "1123123",
				"timestamp": 1562746993127,
				"expire": -1,
        		"body":{
            			"devices":[{
                			"sn":"sn011101",
				          "devId":0102030405,
                			"code":0,
               	 			"msg":"成功"

           			 }{
                			"sn":"sn011102",
				"devId":0102030406
                			"code":0,
               	 			"msg":"成功"
           			 }]
       			 }
    		}
***************************/
type DeviceRegAckHead3900 struct {
	Method          string               `json:"method"`
	Mid             int64                `json:"mid"`
	DevId           string               `json:"devId"`
	Expire          int                  `json:"expire"`
	Timestamp       int64                `json:"timestamp"`
	Body            DeviceRegAckBody3900 `json:"body"`

}


type DeviceRegAckBody3900 struct{
	Devices     []DeviceRegAck        `json:"devices"`
}
type DeviceRegAck struct {
	Sn          string                `json:"sn"`
	DevId       int64                 `json:"devId"`
	Code        int                   `json:"code"`
	Msg         string                `json:"msg"`
}

/******************************************************************
devId Number 是 主站为设备分配的唯一 ID
isOnline Number 是 设备上线或者下线,1 表示上线;0 表示下线
*******************************************************************/
type DevOnlineStatusBody3900 struct {
	DevId           int                   `json:"devId"`
	IsOnline        int64                 `json:"isOnline"`
}

/******************************************************************
设备-->平台
	topic;/v1/2/4217058101079148192/2/service/telemetry/format
	内容:{
		"mid":111111111,
        		"method":"DATA_FORMAT",
        		"timestamp":1574930415426,
		"devId":"4217058101079148192",
        		"body":{
            			"devId":"4217058101079148192",
            			"data":{
                			"Ua":220.3,
                			"Ub":221,
               	 			"Ia":23.4    //属性  key-value结构
           			 }
       			 }
    		}
*****************************************************/

type DataReportHead3900 struct {
	Method          string               `json:"method"`
	Mid             int64                `json:"mid"`
	DevId           string               `json:"devId"`
	Timestamp       int64                `json:"timestamp"`
	Body            DataReportMsg         `json:"body"`

}
type DataReportMsg struct {
	DevId           string                  `json:"devId"`
	Data            map[string]interface{}  `json:"data"`
}

/*****************************************************
下发:平台-->设备
	topic:/v1/${devType}/4217058101079148192/service/control
	内容:{
        		"method":"CMD_CALL_DATA",
"mid": 1569582513983943814,
				"devId": "1123123",
				"timestamp": 1562746993127,
				"expire": -1,
        		"body":{
			"devId":"2323232323",
			"type":"analog",
            			"params":["Ua",”Ub”]
       			 }
    		}
*****************************************************/

type DataCallCmdHead3900 struct {
	Method          string               `json:"method"`
	Mid             int64                `json:"mid"`
	DevId           string               `json:"devId"`
	Timestamp       int64                `json:"timestamp"`
	Expire          int                  `json:"expire"`
	Body            DataCallCmdMsg         `json:"body"`

}
type DataCallCmdMsg struct {
	DevId           string                  `json:"devId"`
	Type            string                  `json:"type"`
	Params          []string                `json:"params"`
}

/*****************************************************
响应:设备-->平台
	topic:/v1/2/4217058101079148192/service/response/format
	内容:{
		"mid":11111111,
		"devId":"4217058101079148192",
        		"method":"CMD_CALL_DATA",
        		"timestamp":1574930415425,
		"code":200,
        		"body":{
            			"devId":"4217058101079148192",
            			"data":{
                			"Ua":220.3,
                			"Ub":221
           			 }
        		}
   	}
*****************************************************/

type DataCallAckHead3900 struct {
	Method          string               `json:"method"`
	Mid             int64                `json:"mid"`
	DevId           string               `json:"devId"`
	Timestamp       int64                `json:"timestamp"`
	Code            int                  `json:"code"`
	Body            DataReportMsg        `json:"body"`
}

/*****************************************************
设备-->平台
	topic:/v1/2/4217058101079148192/device/telemetry
	内容:{
		"mid":111111111,
        		"method":"EVENT_SYS_ALARM",
        		"timestamp":1574930415425,
        		"devId":4217058101079148192,
        		"body":{
            			"event":"101",  //典型事件类型  101告警事件,102安全事件,103故障事件
            			"msg":"cpu 利用率超过阀值"      //事件描述
        			}
    		}
*****************************************************/
type DeviceEventReportHead3900 struct {
	Method          string               `json:"method"`
	Mid             int64                `json:"mid"`
	DevId           string               `json:"devId"`
	Timestamp       int64                `json:"timestamp"`
	Body            DeviceEventMsg       `json:"body"`

}
type DeviceEventMsg struct {
	Event           string                  `json:"event"`
	Msg             string                  `json:"msg"`
}


/*****************************************************
5. 设备配置查询
命令:平台-->设备
	topic:/v1/${edgeId}/device/control
	内容:{
		"mid":111111111,
        		"method":"CMD_CONFIG_QUERY",
        		"timestamp":1574930415426,
		"devId":"4217058101079148192",
        		"body":{
            			"items":[]//全部查询
}
    		}
*****************************************************/
type DeviceConfigCmdHead3900 struct {
	Method          string               `json:"method"`
	Mid             int64                `json:"mid"`
	DevId           string               `json:"devId"`
	Timestamp       int64                `json:"timestamp"`
	Body            DeviceEventMsg       `json:"body"`

}
type DeviceConfigCmdMsg struct {
	Items           []string             `json:"items"`
}

/*****************************************************
命令应答:设备-->平台
	topic:/v1/${edgeId}/device/response
	内容:{
		"mid":111111111,
        		"method":"CMD_CONFIG_QUERY",
        		"timestamp":1574930415426,
		"devId":"4217058101079148192",
        		"body":[
			{
			"cpu":{
				"cpunum":4,
				"frequency":2.9,
				"cache":512
			}},
			{
			"mem":{
				"phy":32,
				"virt":4
			}},
			{
			"link":[{
				"type":"Ethernet",
				"name":"eth1"
			},{
				"type":"Ethernet",
				"name":"eth1"
			}]
			},{
			"os":{
				"distro":"Ubunut",
				"version":"18.10",
				"kernel":"3.10-17"
			}
			},{
			"disk":1024
			}
			]
    		}
*****************************************************/
type DeviceConfigAckHead3900 struct {
	Method          string               `json:"method"`
	Mid             int64                `json:"mid"`
	DevId           string               `json:"devId"`
	Timestamp       int64                `json:"timestamp"`
	Body            []interface{}        `json:"body"`

}
/*****************************************************/
type CpuAckMsg struct {
	Items           CpuResourceMsg           `json:"cpu"`
}
type CpuResourceMsg struct {
	Cpunum          int                  `json:"cpunum"`
	Frequency       int                  `json:"frequency"`
	Cache           int                  `json:"cache"`
}
/*****************************************************/
type MemAckMsg struct {
	Items           MemResourceMsg           `json:"mem"`
}
type MemResourceMsg struct {
	Phy             int                  `json:"phy"`
	Virt            int                  `json:"virt"`
	//Cache           int                  `json:"cache"`
}
/*****************************************************/
type LinkAckMsg struct {
	Items           []LinkResourceMsg           `json:"link"`
}
type LinkResourceMsg struct {
	Type            string                  `json:"type"`
	Name            string                  `json:"name"`
	//Cache           int                  `json:"cache"`
}
/*****************************************************/
type OsAckMsg struct {
	Items           OSResourceMsg           `json:"os"`
}
type OSResourceMsg struct {
	Distro          string                  `json:"distro"`
	Version         string                  `json:"version"`
	Kernel          string                  `json:"kernel"`
}
/*****************************************************/
type DiskAckMsg struct {
	Disk            int                      `json:"disk"`
}

/*****************************************************
6. 设备状态查询
查询命令 平台-->设备
	topic:/v1/${edgeId}/device/control
	内容:{
		"mid":111111111,
        		"method":"CMD_SYS_STATUS",
        		"timestamp":1574930415426,
		"devId":"4217058101079148192",
        		"body":{

       			 }
    		}
*****************************************************/
type DeviceStatusReqHead3900 struct {
	Method          string               `json:"method"`
	Mid             int64                `json:"mid"`
	DevId           string               `json:"devId"`
	Timestamp       int64                `json:"timestamp"`
	Body            NullBodyMsg          `json:"body"`
}
type NullBodyMsg struct {
	//Items           []string             `json:"items"`
}

/*****************************************************
查询应答   设备-->平台
	topic:/v1/${edgeId}/device/response
	内容:{
		"mid":111111111,
        		"method":"CMD_SYS_STATUS",
        		"timestamp":1574930415426,
		"devId":"4217058101079148192",
        		"body":[
			{
			"cpu":"cpu信息"},
			{
			"mem":{
				"phy":32,
				"virt":4
			}},
			{
			"link":[{
				"status":"up",
				"name":"eth1"
			},{
				"status":"down",
				"name":"eth1"
			}]
			},{
			"disk":1024
			}
			]
    		}
*****************************************************/
type DeviceStatusAckHead3900 struct {
	Method          string               `json:"method"`
	Mid             int64                `json:"mid"`
	DevId           string               `json:"devId"`
	Timestamp       int64                `json:"timestamp"`
	Body            []interface{}        `json:"body"`

}
/*****************************************************/
type CpuStatusMsg struct {
	Cpu            string           `json:"cpu"`
}

/*****************************************************/
type MemStatusMsg struct {
	Items           MemStatus           `json:"mem"`
}
type MemStatus struct {
	Phy             int                  `json:"phy"`
	Virt            int                  `json:"virt"`
	//Cache           int                  `json:"cache"`
}
/*****************************************************/
type LinkStatusMsg struct {
	Items           []LinkStatus           `json:"link"`
}
type LinkStatus struct {
	Status          string                  `json:"status"`
	Name            string                  `json:"name"`
	//Cache           int                  `json:"cache"`
}
/*****************************************************/
type DiskStatusMsg struct {
	Disk            int              `json:"disk"`
}

/*****************************************************
7. 设备状态主动上报
设备上报状态:  设备-->平台
	topic:/v1/${edgeId}/device /data
	内容:{
		"mid":111111111,
        		"method":"EVENT_SYS_STATUS_REAL_TIME",
        		"timestamp":1574930415426,
		"devId":"4217058101079148192",
        		"body":[
			{
			"cpu":"cpu信息"},
			{
			"mem":{
				"phy":32,
				"virt":4
			}},
			{
			"link":[{
				"status":"up",
				"name":"eth1"
			},{
				"status":"down",
				"name":"eth1"
			}]
			},{
			"disk":1024
			}
			]
    		}
同上 查询接口的应答
*****************************************************/


/*****************************************************
命令下发：平台-->装置
MQTT：
/v1/${edgeId}/device/control
{
	"mid": 1569582513983943814,
	"devId": "1123123",
	"method": "CMD_SYS_UPGRADE",
	"timestamp": 1562746993127,
	"expire": -1,
	"body": {
		"params": [
			{
				"protocol": 0, //传输方式,0默认表示HTTPS，1表示MQTT
				"version": "1.0.1",
				"duration": 5,
				"file": {
					"path": "http://XXX",
					"size": 12,
					"md5": "md5值",
					"name": "设备升级"
				},
				"jobId": 5,
				"command": "对升级文件执行的命令"
			}
		]
	},
	"sig": "body签名"
}
*****************************************************/
type DeviceUpdateHead3900 struct {
	Method          string               `json:"method"`
	Mid             int64                `json:"mid"`
	DevId           string               `json:"devId"`
	Timestamp       int64                `json:"timestamp"`
	Expire          int                  `json:"expire"`
	Body            DevUpdateBodyMsg     `json:"body"`
	Sig             string               `json:"sig"`
}
type DevUpdateBodyMsg struct {
	Params          []UpdateParams       `json:"params"`
}
type UpdateParams struct {
	Command         string               `json:"command"`
	Protocol        int                  `json:"protocol"`
	Version         string               `json:"version"`
	Duration        int                  `json:"duration"`
	File            Filejson             `json:"file"`
	JobId           int                  `json:"jobId"`
}
type Filejson struct {
	Name      string `json:"name"`     ///
	Addr      string `json:"path"`     ///文件在文件共享系统中的路径
	Size      string `json:"size"`     ///文件的大小，单位为字节
	Md5       string `json:"md5"`      ///文件的MD5值
}

/*****************************************************
状态上报:装置-->平台
MQTT：
/v1/${edgeId}/device/response
{
	"mid": 1569582513983943814,
	"devId": "1123123",
	"method": "CMD_SYS_UPGRADE",
	"timestamp": 1562746993127,
	"code": 500,
	"body": {
		"state": "200",
		"jobId": 200,
		"version": "V1.2"
	},
	"sig": "body签名" }


*****************************************************/
type DeviceUpdateAckHead3900 struct {
	Method          string               `json:"method"`
	Mid             int64                `json:"mid"`
	DevId           string               `json:"devId"`
	Timestamp       int64                `json:"timestamp"`
	//Expire          int                  `json:"expire"`
	Body            DevUpdateAckBodyMsg     `json:"body"`
	Sig             string               `json:"sig"`
}
type DevUpdateAckBodyMsg struct {

	State           string               `json:"state"`
	Version         string               `json:"version"`
	JobId           int                  `json:"jobId"`
}

/*****************************************************
9. 物模型下发
下发：平台-->装置
MQTT：
/v1/${edgeId}/device/control
{
	"mid": 1569582513983943814,
	"devId": "1123123",
	"method": "CMD_PROFILE_DOWN",
	"timestamp": 1562746993127,
	"expire": -1,
	"body": {
		"devId": 1123123,
		"file": {
			"path": "http://XXX",
			"size": 12,
			"md5": "md5值",
			"name": "XXXprofile"
		}
	},
	"sig": "body签名"
}
*****************************************************/
type ProfileDownCmdHead3900 struct {
	Method          string               `json:"method"`
	Mid             int64                `json:"mid"`
	DevId           string               `json:"devId"`
	Timestamp       int64                `json:"timestamp"`
	Expire          int                  `json:"expire"`
	Body            ProfileDownBodyMsg     `json:"body"`
	Sig             string               `json:"sig"`
}
type ProfileDownBodyMsg struct {
	DevId           int                  `json:"devId"`
	File            Filejson             `json:"file"`
}

/*****************************************************
命令响应：装置-->平台
MQTT：
/v1/${edgeId}/device/response

{
	"mid": 1569582513983943814,
	"devId": 1123123,
	"method":  "CMD_PROFILE_DOWN",
	"timestamp": 1562746993127,
"code": 500,
	"body": {
		"devId": 1123123,
		"code": "200",
		"msg": "消息请求成功"
	},
	"sig": "body签名"
}
*****************************************************/
type ProfileDownAckHead3900 struct {
	Method          string               `json:"method"`
	Mid             int64                `json:"mid"`
	DevId           string               `json:"devId"`
	Timestamp       int64                `json:"timestamp"`
	Code            int                  `json:"code"`
	Body            ProfileDownAckMsg     `json:"body"`
	Sig             string               `json:"sig"`
}
type ProfileDownAckMsg struct {
	DevId           int                  `json:"devId"`
	Code            int                  `json:"code"`
	Msg             string               `json:"msg"`
}


/*****************************************************
10. 应用安装
平台-->装置
MQTT：
/v1/${edgeId}/app/control
{
	"mid": 1569582513983943814,
	"devId": "1123123",
	"method": "CMD_FILE_DOWN",
	"timestamp": 1562746993127,
	"expire": -1,
	"body": {
		"file": {
			"path": "http://XXX",
			"size": 12,
			"md5": "md5值",
			"name": "XXXApp"
		},
		"policy": "1，立即下载:”1” 2.定时下载 “20191107131313”",
		"command":"给出下载后的执行命令（串）",
		"status":"“1”是开机自启动，“0”为开机不自启动"
	},
	"sig": "body签名"
}
*****************************************************/
type AppDownCmdHead3900 struct {
	Method          string               `json:"method"`
	Mid             int64                `json:"mid"`
	DevId           string               `json:"devId"`
	Timestamp       int64                `json:"timestamp"`
	Expire          int                  `json:"expire"`
	Body            AppDownBodyMsg       `json:"body"`
	Sig             string               `json:"sig"`
}
type AppDownBodyMsg struct {
	Policy          string               `json:"policy"`
	Command         string               `json:"command"`
	Status          string               `json:"status"`
	File            Filejson             `json:"file"`
}

/*****************************************************
状态上报：装置-->平台
MQTT：
/v1/${edgeId}/app/response
{
	"mid": 1569582513983943814,
	"devId": "1123123",
	"method": "CMD_FILE_DOWN",
	"timestamp": 1562746993127,
	"code": "500",
	"body": {
		"code": "200",
		"msg": "消息请求成功"
	},
	"sig": "body签名"
}
*****************************************************/
type AppDownAckHead3900 struct {
	Method          string               `json:"method"`
	Mid             int64                `json:"mid"`
	DevId           string               `json:"devId"`
	Timestamp       int64                `json:"timestamp"`
	Code            int                  `json:"code"`
	Body            AppDownAckMsg       `json:"body"`
	Sig             string               `json:"sig"`
}
type AppDownAckMsg struct {
	Code            int                  `json:"code"`
	Msg             string               `json:"msg"`
}

/*****************************************************
平台-->装置
MQTT：
/v1/${edgeId}/app/control
{
	"mid": 1569582513983943814,
	"devId": "1123123",
	"method": "CMD_APP_UPDATE",
	"timestamp": 1562746993127,
	"expire": -1,
	"body": {
		"name": "XXXApp",
		"file": {
			"path": "http://XXX",
			"size": 12,
			"md5": "md5值",
			"name": "XXXApp"
		},
		"command":"给出下载后的执行命令（串）",
		"status":"“1”是开机自启动，“0”为开机不自启动",
		"policy": "1，立即升级:”1” 2.定时升级 “20191107131313” "
	},
	"sig": "body签名"}
*****************************************************/
type AppUpdateCmdHead3900 struct {
	Method          string               `json:"method"`
	Mid             int64                `json:"mid"`
	DevId           string               `json:"devId"`
	Timestamp       int64                `json:"timestamp"`
	Expire          int                  `json:"expire"`
	Body            AppUpdateBodyMsg       `json:"body"`
	Sig             string               `json:"sig"`
}
type AppUpdateBodyMsg struct {
	Name            string               `json:"name"`
	Policy          string               `json:"policy"`
	Command         string               `json:"command"`
	Status          string               `json:"status"`
	File            Filejson             `json:"file"`
}

/*****************************************************
状态上报：装置-->平台
MQTT：
/v1/${edgeId}/app/response
{
	"mid": 1569582513983943814,
	"devId": "1123123",
	"method": "CMD_APP_UPDATE",
	"timestamp": 1562746993127,
	"code": "500",
	"body": {
		"code": "200",
		"msg": "消息请求成功"
	},
	"sig": "body签名"
}
*****************************************************/
type AppUpdateAckHead3900 struct {
	Method          string               `json:"method"`
	Mid             int64                `json:"mid"`
	DevId           string               `json:"devId"`
	Timestamp       int64                `json:"timestamp"`
	Code            int                  `json:"code"`
	Body            AppDownAckMsg       `json:"body"`
	Sig             string               `json:"sig"`
}

/*****************************************************
12. 应用控制
命令：平台-->装置
MQTT：
/v1/${edgeId}/app/control

{{
	"mid": 1569582513983943814,
	"devId": "1123123",
	"method": "CMD_APP_CONTROL",
	"timestamp": 1562746993127,
	"expire": -1,
	"body": {
		"app": "XXXApp1",
		"command": "start/stop/remove"
	},
	"sig": "body签名"
}
*****************************************************/
type AppControlCmdHead3900 struct {
	Method          string               `json:"method"`
	Mid             int64                `json:"mid"`
	DevId           string               `json:"devId"`
	Timestamp       int64                `json:"timestamp"`
	Expire          int                  `json:"expire"`
	Body            AppControlBodyMsg       `json:"body"`
	Sig             string               `json:"sig"`
}
type AppControlBodyMsg struct {
	app             string               `json:"app"`
	Command         string               `json:"command"`
}

/*****************************************************
命令上报：装置-->平台
MQTT：
/v1/${edgeId}/app/response
{
	"mid": 1569582513983943814,
	"devId": "1123123",
	"method": "CMD_APP_CONTROL",
	"timestamp": 1562746993127,
	"code": "100",
	"body": {
		"code": "200",
		"msg": "消息请求成功"
	},
	"sig": "body签名"
}
============   type AppUpdateAckHead3900 struct {
*****************************************************/

/*****************************************************
13. 应用状态查询
查询命令：平台-->装置
MQTT：
/v1/${edgeId}/app/control

{
	"mid": 1569582513983943814,
	"devId": "1123123",
	"method": "CMD_APP_STATUS",
	"timestamp": 1562746993127,
	"expire": -1,
	"body": {
		"name": "XXXApp"
	},
	"sig": "body签名"
}

*****************************************************/
type AppStatusCmdHead3900 struct {
	Method          string               `json:"method"`
	Mid             int64                `json:"mid"`
	DevId           string               `json:"devId"`
	Timestamp       int64                `json:"timestamp"`
	Expire          int                  `json:"expire"`
	Body            AppStatusBodyMsg       `json:"body"`
	Sig             string               `json:"sig"`
}
type AppStatusBodyMsg struct {
	Name            string               `json:"name"`
	///Command         string               `json:"command"`
}

/*****************************************************
查询结果上报：装置-->平台
MQTT：
/v1/${edgeId}/app/response
{
	"mid": 1569582513983943814,
	"devId": "1123123",
	"method": "CMD_APP_STATUS",
	"timestamp": 1562746993127,
	"code": "600",
	"body": {
		"code": "200",
		"msg": "消息请求成功",
		"app":{
			"name": "XXXApp",
			"version": "1.0.1",
			"profileversion": "1.2",
			"status": "on/off",
			"resource": "{xxx}"
		}

	},
	"sig": "body签名"
}
*****************************************************/
type AppStatusAckHead3900 struct {
	Method          string               `json:"method"`
	Mid             int64                `json:"mid"`
	DevId           string               `json:"devId"`
	Timestamp       int64                `json:"timestamp"`
	Code            int                  `json:"code"`
	Body            AppStatusAckMsg       `json:"body"`
	Sig             string               `json:"sig"`
}
type AppStatusAckMsg struct {
	Code            int                  `json:"code"`
	Msg             string               `json:"msg"`
	App             AppStatus            `json:"app"`
}
type AppStatus struct {
	Name            string               `json:"Name"`
	Version         string               `json:"version"`
	Profileversion  string               `json:"profileversion"`
	Status          string               `json:"status"`
	Resource        string               `json:"Resource"`
}
/*****************************************************
14. 应用配置下发
平台-->装置
MQTT：
/v1/${edgeId}/app/control

{
	"mid": 1569582513983943814,
	"devId": "1123123",
	"method": "CMD_CONFIG_DOWN",
	"timestamp": 1562746993127,
	"expire": -1,
	"body": {
		"name": "XXXApp配置文件",
		"file": {
			"path": "http://XXX",
			"size": 12,
			"md5": "md5值",
			"name": "XXXApp配置文件"
		},
		"policy": "1，立即下载:”1” 2.定时下载 “20191107131313”"
	},
	"sig": "body签名"
}
*****************************************************/
type AppConfigDownCmdHead3900 struct {
	Method          string               `json:"method"`
	Mid             int64                `json:"mid"`
	DevId           string               `json:"devId"`
	Timestamp       int64                `json:"timestamp"`
	Expire          int                  `json:"expire"`
	Body            AppConfigBodyMsg       `json:"body"`
	Sig             string               `json:"sig"`
}

type AppConfigBodyMsg struct {
	Policy          string               `json:"policy"`
	Name            string               `json:"name"`
	//Status          string               `json:"status"`
	File            Filejson             `json:"file"`
}
/*****************************************************
下发结果上报：装置-->平台
MQTT：
/v1/${edgeId}/app/response
{
	"mid": 1569582513983943814,
	"devId": "1123123",
	"method": "CMD_CONFIG_DOWN",
	"timestamp": 1562746993127,
	"code": "100",
	"body": {
		"code": "200",
		"msg": "消息请求成功"
	},
	"sig": "body签名"
}

type AppUpdateAckHead3900 struct {
*****************************************************/

/*****************************************************
15. 应用配置控制
平台-->装置
MQTT：
/v1/${edgeId}/app/control

{
	"mid": 1569582513983943814,
	"devId": "1123123",
	"method": "CMD_CONFIG_CONTROL",
	"timestamp": 1562746993127,
	"expire": -1,
	"body": {
		"name": "XXXApp配置文件",
		"action": "active/runback"
	},
	"sig": "body签名"
}
*****************************************************/
type AppConfigControlHead3900 struct {
	Method          string               `json:"method"`
	Mid             int64                `json:"mid"`
	DevId           string               `json:"devId"`
	Timestamp       int64                `json:"timestamp"`
	Expire          int                  `json:"expire"`
	Body            AppConfigCtrlMsg       `json:"body"`
	Sig             string               `json:"sig"`
}
type AppConfigCtrlMsg struct {
	Name            string               `json:"name"`
	Action          string               `json:"action"`
}

/*****************************************************
结果上报：装置-->平台
MQTT：
/v1/${edgeId}/app/response
{
	"mid": 1569582513983943814,
	"devId": "1123123",
	"method": "CMD_CONFIG_CONTROL",
	"timestamp": 1562746993127,
	"code": "500",
	"body": {
		"code": "200",
		"msg": "消息请求成功"
	},
	"sig": "body签名"
}
type AppUpdateAckHead3900 struct {
*****************************************************/

/*****************************************************

*****************************************************/

/*****************************************************

*****************************************************/

/*****************************************************

*****************************************************/


/*****************************************************

*****************************************************/

/*****************************************************

*****************************************************/

/*****************************************************

*****************************************************/

/*****************************************************

*****************************************************/

/*****************************************************

*****************************************************/


/*****************************************************

*****************************************************/


/*****************************************************

*****************************************************/
