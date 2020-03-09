package define

/// ==== 上下线的 chan
type DeviceOnOffChan struct{
	DevId             string          `json:"devId"`
	OnOff             int             `json:"OnOff"`

	RemoteId          string          `json:"remoteId"`
}


