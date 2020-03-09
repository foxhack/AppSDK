package nsplus

const (
	BootTimeoutDefault   = 30000
	ClientMonitorDefault = 15000
	ConfigFileName       = "configuration.toml"
	ConfigRegistryStem   = "edgex/core/1.0/"
	LogDurationKey       = "duration"
)

const(

	// ==== For nsplus profile  data type
	AppType_Int        = "Int"
	AppType_Double     = "Double"
	AppType_String     = "String"
	AppType_Bool       = "Bool"
	AppType_Enum       = "Enum"
	AppType_Time       = "Time"
	AppType_Array      = "Array"
	AppType_File       = "File"
	AppType_Stream     = "Stream"

	AppType_Event      = "Event"

	AppType_DataSet    = "DataSet"


	/// ====

	ZERO_Protocol       = "tcp"
	ZERO_Port           = 5577
	ZERO_Type           = "zero"
	ZERO_Topic          = "events"

	/// '/edgex/logs/edgex-app-nsplusdemo.log'


	APPCtrlTABLE       = "appctrltable.json"


	ProfileExt             = ".profile"
	RegfileExt             = ".reginfo"

	App_Log_Target      = "/edgex/logs/"
	// ====
	Protocol_HTTP       = "http"
	PORT_ConfigMicSvr   = "20000"


	Backend_Port           = 5555

	PORT_FILTERMicSvr   = "23000"
	Api_Get_IndexMap    = "/api/v1/df/cmd/sync/uuid/"  ///"/api/v1/df/cmd/sync/uuid/{uuid}"



	Timer_Format_UCT        = "2006-01-02 15:04:05"
	Time_Format_dot001      = "2006-01-02 15:04:05.009"

)
