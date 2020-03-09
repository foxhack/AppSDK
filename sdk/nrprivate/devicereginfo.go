package nrprivate

import (
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)



// 用于注册的界面使用的接口
type DeviceAutoRegInfo struct{
	ServiceName      string
	DeviceName       string
	DeviceModel      string
	ManufactureId    string
	Profile          string
	Description      string
	Labels           []string
	Protocols       map[string]models.ProtocolProperties
	RegCount         int32
	RegDateTime      int64
}

