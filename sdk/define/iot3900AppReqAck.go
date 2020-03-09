package define
/********************* 互联网部规范 2020.01 *********************/
/*****************************************
webserver：网关
Url: /api/v1/iot/cmd/app/install
Method: PUT

18、App安装流程    平台到设备   	topic:/v1/${edgeId}/app/command
	内容：
	{
		"deviceId": "1577063490614240253",
		"type": "CMD_APP_INSTALL",
		"expire": -1,
		"mid": 1569582513983943816,
		"param": {
			"jobId": 1569582513983943816,
			"container": "容器名称，写死",
			"version": "V1.2",
			"file": {
				"url": "http://26.47.73.135:23004/v2/iot/files/1231/up/134E1BB7-061D-4F89-9A86-FD7BB953AA55.tar",
				"size": 12,
				"name": "TestApp1",
				"md5": "d41d8cd98f00b204e9800998ecf8427e"
			},
			"cpu": {					//规范 未明确
				"cpus": 4,
				"cpuLmt": 50
			},
			"mem": {
				"memory": 12,
				"memLmt": 50
			}
			"enable": "1",
			"policy": 1
		},
		"timestamp": 1577096316
	}

type IOTCmdAppInstall struct {
	Mid                int64                           `json:"mid"`
	DeviceId           string                          `json:"deviceId"`
	Timestamp          int64                           `json:"timestamp"`
	Expire             int                             `json:"expire"`
	Type               string                          `json:"type"`
	Param              IOTAppInstall                   `json:"param"`
}*****************************************/
type IOTAppInstall struct {
	JobId              int64                           `json:"jobId"`
	Container          string                          `json:"container"`
	Version            string                          `json:"version"`
	File               IOTFileDefine                   `json:"file"`
	Cpu                IOTAppCpuDefine                 `json:"cpu"`
	Mem                IOTAppMemDskDefine              `json:"mem"`
	Enable             string                          `json:"enable"`
	Policy             int                             `json:"policy"`
}
/*****************************************
如果该命令可以执行，则 应答该命令是否被接受。
19、App安装命令响应  设备到平台   	topic:/v1/${edgeId}/app/reply
	内容:
	{
		"mid": 1569582513983943816,
		"deviceId": "1577063490614240253",
		"type": "CMD_APP_INSTALL",
		"timestamp": 1577073563000,
		"code": 202,
		"msg": "SUCCESS"
	}

IOTAckMsgHeader
 *****************************************/

/*********************************
20、App安装结果上报  设备到平台 	topic:/v1/${edgeId}/app/data
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

type IOTCmdAppInstallStatusAck struct {
	Mid                int64                           `json:"mid"`
	DeviceId           string                          `json:"deviceId"`
	Timestamp          int64                           `json:"timestamp"`
	Expire             int                             `json:"expire"`
	Type               string                          `json:"type"`
	Param              IOTAppInstallStatusAck          `json:"param"`
	Code               int                             `json:"code"`
	Msg                string                          `json:"msg"`
}*********************************/
type IOTAppInstallStatusAck struct {
	JobId              int64                           `json:"jobId"`
	Result             int                             `json:"result"`
	Info               string                          `json:"info"`
}

/*********************************
21、App升级流程    平台到设备  	topic:/v1/${edgeId}/app/command
	内容：
	{
		"deviceId": "1577063490614240253",
		"type": "CMD_APP_UPGRADE",
		"expire": -1,
		"mid": 1569582513983943816,
		"param": {
			"jobId": 1569582513983943816,
			"container": "容器名称，写死",
			"version": "V1.2",
			"file": {
				"url": "http://26.47.73.135:23004/v2/iot/files/1231/up/134E1BB7-061D-4F89-9A86-FD7BB953AA55.tar",
				"size": 12,
				"name": "TestApp1",
				"md5": "d41d8cd98f00b204e9800998ecf8427e"
			},
			"policy": 1
		},
		"timestamp": 1577096316
	}
*********************************/
type IOTAppUpdate struct {
	JobId              int64                           `json:"jobId"`
	Container          string                          `json:"container"`
	Version            string                          `json:"version"`
	File               IOTFileDefine                   `json:"file"`
	Policy             int                             `json:"policy"`
}

/*****************************************
22、App升级命令响应  设备到平台 	topic:/v1/${edgeId}/app/reply
	内容:
	{
		"mid": 1569582513983943816,
		"deviceId": "1577063490614240253",
		"type": "CMD_APP_UPGRADE",
		"timestamp": 1577073563000,
		"code": 202,
		"msg": "SUCCESS"
	}
IOTAckMsgHeader
 *****************************************/

/*********************************
23、App升级结果上报  设备到平台  	topic:/v1/${edgeId}/app/data
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
  同 20、App安装结果上报  设备到平台 	topic:/v1/${edgeId}/app/data
   IOTAppInstallStatusAck
*********************************/

/*********************************
24、应用查询状态  平台到设备  	topic:/v1/${edgeId}/app/command
	内容:
	{
		"deviceId": "1577123434420863986",
		"type": "CMD_APP_STATUS",
		"expire": -1,
		"mid": 777677455,
		"param": {
			"container": "容器名称"
		},
		"timestamp": 1577106435
	}
*********************************/
type IOTAppStatus struct {
	Container          string                          `json:"container"`
}
/*********************************
25、应用查询状态应答  设备到平台  	topic:/v1/${edgeId}/app/reply
	内容:
	{
		"mid": 777677455,
		"deviceId": "1577123434420863986",
		"type": "CMD_APP_STATUS",
		"timestamp": 1562746993127,
		"code": 200,
		"msg": "SUCCESS",
		"param": {
			"container": "容器名称",
			"apps": [{
				"app": "下发测试",
				"version": "1.0.1",
				"appHash": "APP 的哈希值",
				"srvNumber": 12,
				"process": [{
					"srvIndex": 123,
					"srvName": "进程名称",
					"srvEnable": "on",
					"srvStatus": "running",
					"cpuLmt": 50,
					"cpuRate": 10,
					"memLmt": 60,
					"memUsage": 4.8,
					"startTime": "服务启动时间"
				}]
			}]
		}
	}
*********************************/
type IOTAppStatusAck struct {
	Container          string                          `json:"container"`
	Apps               []AppStatusInfo                 `json:"apps"`
}
type AppStatusInfo struct {
	App                string                          `json:"app"`
	Version            string                          `json:"version"`
	AppHash            string                          `json:"appHash"`
	SrvNumber          string                          `json:"srvNumber"`
	Process            []AppProcessInfo                `json:"srvNumber"`
}
type AppProcessInfo struct {
	SrvIndex           int                             `json:"srvIndex"`
	SrvName            string                          `json:"srvName"`
	SrvEnable          string                          `json:"srvEnable"`
	SrvStatus          string                          `json:"srvStatus"`
	CpuLmt             int                             `json:"cpuLmt"`
	CpuRate            int                             `json:"cpuRate"`
	MemLmt             int                             `json:"memLmt"`
	MemUsage           int                             `json:"memUsage"`
	StartTime          string                          `json:"startTime"`
}
/*********************************
26、应用状态自动上报  设备到平台  	topic:/v1/${edgeId}/app/data
	内容:
	{
		"mid": 777677455,
		"deviceId": "1577123434420863986",
		"type": "CMD_APP_STATUS",
		"timestamp": 1562746993127,
		"code": 200,
		"msg": "SUCCESS",
		"param": {
			"container": "容器名称",
			"apps": [{
				"app": "下发测试",
				"version": "1.0.1",
				"appHash": "APP 的哈希值",
				"srvNumber": 12,
				"process": [{
					"srvIndex": 123,
					"srvName": "进程名称",
					"srvEnable": "on",
					"srvStatus": "running",
					"cpuLmt": 50,
					"cpuRate": 10,
					"memLmt": 60,
					"memUsage": 4.8,
					"startTime": "服务启动时间"
				}]
			}]
		}
	}
   同25    IOTAppStatusAck
*********************************/


/*********************************
27、应用启动命令	平台到设备  	topic:/v1/${edgeId}/app/command
	内容:
	{
		"deviceId": "1577123434420863986",
		"type": "CMD_APP_START",
		"expire": -1,
		"mid": 667645741,
		"param": {
			"container": "容器名称",
			"app": "app名称"
		},
		"timestamp": 1577106555
	}
*********************************/
type IOTAppCommand struct {
	Container          string                          `json:"container"`
	App                string                          `json:"app"`
}

/*********************************
28、应用启动命令应答	设备到平台
	topic:/v1/${edgeId}/app/reply
	内容:
	{
		"mid": 667645741,
		"deviceId": "1577123434420863986",
		"type": "CMD_APP_START",
		"timestamp": 1562746993127,
		"code": 200,
		"msg": "SUCCESS"
	}
    IOTAckMsgHeader
*********************************/


/*********************************
29、应用停止命令	平台到设备  	topic:/v1/${edgeId}/app/command
	内容:
	{
		"deviceId": "1577123434420863986",
		"type": "CMD_APP_STOP",
		"expire": -1,
		"mid": 667645741,
		"param": {
			"container": "容器名称",
			"app": "app名称"
		},
		"timestamp": 1577106555
	}
    IOTAppCommand
*********************************/


/*********************************
30、应用停止命令应答	设备到平台 	topic:/v1/${edgeId}/app/reply
	内容:
	{
		"mid": 667645741,
		"deviceId": "1577123434420863986",
		"type": "CMD_APP_STOP",
		"timestamp": 1562746993127,
		"code": 200,
		"msg": "SUCCESS"
	}
    IOTAckMsgHeader
*********************************/

/*********************************
31、应用卸载命令：	平台到设备  	topic:/v1/${edgeId}/app/command
	内容:
	{
		"deviceId": "1577123434420863986",
		"type": "CMD_APP_REMOVE",
		"expire": -1,
		"mid": 667645741,
		"param": {
			"container": "容器名称",
			"app": "app名称"
		},
		"timestamp": 1577106555
	}
    IOTAppCommand
*********************************/


/*********************************
32、应用卸载命令应答：	设备到平台  	topic:/v1/${edgeId}/app/reply
	内容:
	{
		"mid": 667645741,
		"deviceId": "1577123434420863986",
		"type": "CMD_APP_REMOVE",
		"timestamp": 1562746993127,
		"code": 200,
		"msg": "SUCCESS"
	}
    IOTAckMsgHeader
*********************************/


/*********************************
33、应用使能命令：	平台到设备  	topic:/v1/${edgeId}/app/command
	内容:
	{
		"deviceId": "1577123434420863986",
		"type": "CMD_APP_ENABLE",
		"expire": -1,
		"mid": 667645741,
		"param": {
			"container": "容器名称",
			"app": "app名称"
		},
		"timestamp": 1577106555
	}
    IOTAppCommand
*********************************/


/*********************************
34、应用使能命令应答	设备到平台 	topic:/v1/${edgeId}/app/reply
	内容:
	{
		"mid": 667645741,
		"deviceId": "1577123434420863986",
		"type": "CMD_APP_ENABLE",
		"timestamp": 1562746993127,
		"code": 200,
		"msg": "SUCCESS"
	}
    IOTAckMsgHeader
*********************************/


/*********************************
35、应用去使能命令：	平台到设备 	topic:/v1/${edgeId}/app/command
	内容:
	{
		"deviceId": "1577123434420863986",
		"type": "CMD_APP_UNENABLE",
		"expire": -1,
		"mid": 667645741,
		"param": {
			"container": "容器名称",
			"app": "app名称"
		},
		"timestamp": 1577106555
	}
    IOTAppCommand
*********************************/


/*********************************
36、应用去使能命令应答	设备到平台  	topic:/v1/${edgeId}/app/reply
	内容:
	{
		"mid": 667645741,
		"deviceId": "1577123434420863986",
		"type": "CMD_APP_UNENABLE",
		"timestamp": 1562746993127,
		"code": 200,
		"msg": "SUCCESS"
	}
    IOTAckMsgHeader
*********************************/


/*********************************
37、应用配置修改命令 	平台到设备   	topic:/v1/${edgeId}/app/command
	内容:
	{
		"deviceId": "1577123434420863986",
		"type": "CMD_APP_SET_CONFIG",
		"expire": -1,
		"mid": 667645741,
		"param": {
			"container": "容器名称",
			"app": "app名称",
			"cpu": {				//规范 未明确
				"cpus": 4,
				"cpuLmt": 50
			},
			"mem": {
				"memory": 12,
				"memLmt": 50
			}
		},
		"timestamp": 1577106555
	}
*********************************/
type IOTAppContainerConfig struct {
	Container          string                          `json:"container"`
	App                string                          `json:"app"`
	Cpu                IOTAppCpuDefine                 `json:"cpu"`
	Mem                IOTAppMemDskDefine              `json:"mem"`
}
/*********************************
38、应用配置修改命令应答  	设备到平台  	topic:/v1/${edgeId}/app/reply
	内容:
	{
		"mid": 667645741,
		"deviceId": "1577123434420863986",
		"type": "CMD_APP_SET_CONFIG",
		"timestamp": 1562746993127,
		"code": 200,
		"msg": "SUCCESS"
	}
    IOTAckMsgHeader
*********************************/

/*********************************
39、应用配置查询命令	平台到设备   	topic:/v1/${edgeId}/app/command
	内容:
	{
		"deviceId": "1577123434420863986",
		"type": "CMD_APP_GET_CONFIG",
		"expire": -1,
		"mid": 667645741,
		"param": {
			"container": "容器名称"
		},
		"timestamp": 1577106555
	}
    IOTAppStatus
*********************************/

/*********************************
40、应用配置查询命令应答	设备到平台  	topic:/v1/${edgeId}/app/reply
	内容:
	{
		"mid": 667645741,
		"deviceId": "1577123434420863986",
		"type": "CMD_APP_GET_CONFIG",
		"timestamp": 1562746993127,
		"code": 200,
		"msg": "SUCCESS",
		"param": {
			"container": "容器名称",
			"appCfgs": [
				{
					"app": "app名称",
					"cpuLmt": 50,
					"memLmt": 50
				}
			]
		}
	}
*********************************/
type IOTAppContConfigGetAck struct {
	Container          string                          `json:"container"`
	AppCfgs            []IotAppCfgs                    `json:"appCfgs"`
}
type IotAppCfgs struct {
	App                string                          `json:"app"`
	CpuLmt             int                             `json:"cpuLmt"`
	MemLmt             int                             `json:"memLmt"`
}

/*********************************
41、应用日志召回	平台到设备   	topic:/v1/${edgeId}/app/command
	内容:
	{
		"deviceId": "1577123434420863986",
		"type": "CMD_APP_LOG",
		"expire": -1,
		"mid": 667645741,
		"param": {
			"container": "容器名称",
			"url": "http:XXX"
		},
		"timestamp": 1577106555
	}
*********************************/
type IOTAppCallLog struct {
	Container          string                          `json:"container"`
	Url                string                          `json:"url"`
}

/*********************************
42、应用日志召回应答	设备到平台   	topic:/v1/${edgeId}/app/reply
	内容:
	{
		"mid": 667645741,
		"deviceId": "1577123434420863986",
		"type": "CMD_APP_LOG",
		"timestamp": 1562746993127,
		"code": 200,
		"msg": "SUCCESS",
		"param": {
			"file": {
				"url": "http://26.47.73.135:23004/v2/iot/files/1231/up/134E1BB7-061D-4F89-9A86-FD7BB953AA55.tar",
				"size": 12,
				"name": "TestApp1",
				"md5": "d41d8cd98f00b204e9800998ecf8427e"
			}
		}
	}
*********************************/
type IOTAppCallLogAck struct {
	File               IOTFileDefine                   `json:"file"`
}
/*********************************
43、应用控制范围等业务配置命令下发  	平台到设备    ----客户新增协议----	topic:/v1/${edgeId}/app/command
	内容:
	{
		"mid": 1569582513983943814,
		"deviceId": "1123123",
		"type": "CMD_APP_REMOTE_CONFIG_DOWN",
		"timestamp": 1562746993127,
		"expire": -1,
		"param": {
			"container": "容器名称",
			"appName": "XXXApp名称",
			"name": "XXXApp配置文件名称",
			"configFile": "配置命令内容Base64编码内容"
		},
	}
*********************************/
type IOTAppParameterSet struct {
	Container          string                          `json:"container"`
	AppName            string                          `json:"appName"`
	Name               string                          `json:"name"`
	ConfigFile         string                          `json:"configFile"`
}

/*********************************
44、应用控制范围等业务配置命令下发应答  设备到平台   	topic:/v1/${edgeId}/app/reply
	内容:
	{
		"mid": 1569582513983943814,
		"deviceId": "1123123",
		"type": "CMD_APP_REMOTE_CONFIG_DOWN",
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


/*********************************

*********************************/
