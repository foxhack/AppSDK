package define


///// ==================   NS3900  =======================
const(


	APIIOTEdgexAuth3900_V2            = "/v2/iot/devices/direct/auth"
	//APIIOTDeviceAuth_V2           = "/v2/iot/devices/auth/indirect"
	//APIIOTDeviceOnOff_V2          = "/v2/iot/devices/onoffstatus"

)
const(
	TopicID_DEV_CTL            = "DEV_CTL"   /// down
	TopicID_DEV_INIT           = "DEV_INIT"
	TopicID_DEV_REQ            = "DEV_REQ"
	TopicID_DEV_RES            = "DEV_RES"
	TopicID_DEV_ATTR           = "DEV_ATTR"
	TopicID_DEV_TLM            = "DEV_TLM"


	TopicID_SER_TLM_FMT        = "SER_TLM_FMT"
	TopicID_SER_TLM_RAW        = "SER_TLM_RAW"
	TopicID_SER_CTL            = "SER_CTL"       /// down
	TopicID_SER_CTR_RAW        = "SER_CTR_RAW"  /// down
	TopicID_SER_RES            = "SER_RES"
	TopicID_SER_RES_RAW        = "SER_RES_RAW"

	TopicID_SER_EMG            = "SER_EMG"


	TopicID_APP_CTL            = "APP_CTL"   /// down
	TopicID_APP_RES            = "APP_RES"
	TopicID_APP_TLM            = "APP_TLM"

)

///// ==================   NS3910  =======================
const(


	APIIOTEdgexAuth_V2            = "/v2/iot/devices/auth/direct"
	APIIOTDeviceAuth_V2           = "/v2/iot/devices/auth/indirect"
	APIIOTDeviceOnOff_V2          = "/v2/iot/devices/onoffstatus"

)

const(


	IOT_ACK_Code_Success          = 0
	IOT_ACK_Code_Fail             = 1

	IOT_DEVICE_ONOFF_STATUS_OFF   = 0
	IOT_DEVICE_ONOFF_STATUS_ON    = 1

)