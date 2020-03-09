package nsplus

import (
	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
)
/// init时赋值
func (sdk *NSPlusSdk)SetSdkLogClient(AppNameKey string, logTarget string, LogLevel string){
	//logTarget := setLoggingTarget()
	sdk.LoggingClient = logger.NewClient(AppNameKey, false, logTarget, LogLevel)
}


//func setLoggingTarget() string {
//	if Configuration.Logging.EnableRemote {
//		return Configuration.Clients["Logging"].Url() + clients.ApiLoggingRoute
//	}
//	return Configuration.Logging.File
//}
