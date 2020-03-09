package define
const(
	EVENT_LINKUP      = "EVENT_LINKUP"        /// 1、边缘设备接入请求   设备到平台
	CMD_TOPO_ADD      = "CMD_TOPO_ADD"        /// 3、子设备接入认证（子设备添加问题）  设备到平台
	CMD_TOPO_DEL      = "CMD_TOPO_DEL"        /// 5、子设备删除  设备到平台
	CMD_TOPO_UPDATE   = "CMD_TOPO_UPDATE"     /// 7、子设备状态更新 设备到平台

	EVENT_DATA_ALARM  = "EVENT_DATA_ALARM"    /// 9、设备事件主动上报  设备到平台
	CMD_REPORTDATA    = "CMD_REPORTDATA"      /// 10、设备数据主动上报   设备到平台
	CMD_SERVICE       = "CMD_SERVICE"         /// 11、业务命令下发 平台到设备
	CMD_PROFILE       = "CMD_PROFILE"         /// 13、物模型下发 平台到设备
	CMD_SYS_UPGRADE   = "CMD_SYS_UPGRADE"     /// 15、固件下发（设备升级）流程    平台到设备
	REP_JOB_RESULT    = "REP_JOB_RESULT"      /// 17、固件下发（设备升级）结果上报  设备到平台

	CMD_APP_INSTALL   = "CMD_APP_INSTALL"     /// 18、App安装流程    平台到设备
	//// REP_JOB_RESULT    = "REP_JOB_RESULT"      /// 20、App安装结果上报  设备到平台    同17   23、App升级结果上报  设备到平台
	CMD_APP_UPGRADE   = "CMD_APP_UPGRADE"    /// 21、App升级流程    平台到设备
	CMD_APP_STATUS    = "CMD_APP_STATUS"     /// 24、应用查询状态  平台到设备   25 应用查询状态应答 、 26、应用状态自动上报  设备到平台
	CMD_APP_START     = "CMD_APP_START"      /// 27、应用启动命令	平台到设备
	CMD_APP_STOP      = "CMD_APP_STOP"       /// 29、应用停止命令	平台到设备
	CMD_APP_REMOVE    = "CMD_APP_REMOVE"     /// 31、应用卸载命令：	平台到设备
	CMD_APP_ENABLE    = "CMD_APP_ENABLE"     /// 33、应用使能命令：	平台到设备
	CMD_APP_UNENABLE  = "CMD_APP_UNENABLE"   /// 35、应用去使能命令：	平台到设备
	CMD_APP_SET_CONFIG= "CMD_APP_SET_CONFIG" /// 37、应用配置修改命令 	平台到设备
	CMD_APP_GET_CONFIG= "CMD_APP_GET_CONFIG" /// 39、应用配置查询命令	平台到设备
	CMD_APP_LOG       = "CMD_APP_LOG"        /// 41、应用日志召回	平台到设备

	CMD_APP_REMOTE_CONFIG_DOWN  = "CMD_APP_REMOTE_CONFIG_DOWN"  /// 43、应用控制范围等业务配置命令下发	平台到设备    ----客户新增协议


	CMD_CTRL          = "CMD_CTRL"
	CMD_WEBSSH        = "CMD_WEBSSH"



	///==========================
	Edge_Status_ONLINE= "ONLINE"
	Edge_Status_OfLINE= "OFFLINE"

	Dev_Status_OK     = 200
)


type FileInfo struct{
	Url                    string         `json:"url"`
	Size                   int64          `json:"size"`
	Name                   string         `json:"name"`
	Md5                    string         `json:"md5"`
}



