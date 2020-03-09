package nsplus

import (
	"errors"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/edgexfoundry/nsplussdk/sdk/config"
	"github.com/edgexfoundry/nsplussdk/sdk/define"

	//"github.com/edgexfoundry/edgex-backend/internal/iot/define"
	//"github.com/edgexfoundry/edgex-backend/internal/pkg/config"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
	msgTypes "github.com/edgexfoundry/go-mod-messaging/pkg/types"
	"log"
	"net/http"
	"time"
)
/// init时赋值
//var messageClient messaging.MessageClient
var messageErrors chan error

type NSPlusSdk struct{
	LoggingClient     logger.LoggingClient
	LogTarget         string
	LogLevel          string

	HttpClient        *http.Client
	MqttClient        *mqttSender

	AppName           string  /// 初始化的AppNam, 容器形式时对应baseappname，当不在容器运行时作为App的key
	ContainerName     string  ///  d对应容器App的key，用于唯一标识一个容器app，当不在容器运行时取 初始化的AppNam

	Regfile           EdgeBaseAppRegInfo
	Profile           NsPlusProfile
	AppParamsCtrl     AppParamCtrlBuf

	MQConfig          config.MessageQueueInfo
	UserConfig        define.EdgeConfigStruct

	DataIndexMap      DataIndexForFilter

	MessageEnvelopes  chan msgTypes.MessageEnvelope
	MessageErrors     chan error
	ParamChange       chan AppParams

	MqttMsgFunc         MQTT.MessageHandler
	InitClientsCBFunc   InitClientsCallback
	InputRecvCBFunc     InputRecvCallback
	ParameterRecvCBfunc ParameterRecvCallback

}

/// ==== NSPlusSdk AppName和 DataIp DataPort是必选参数，日志设置可默认
func NewNsplusApp(appname string) (*NSPlusSdk, error) {

	if appname == "" {
		log.Printf("appname can not be null [%v]!!", appname)
		return nil, errors.New("appname can not be null")
	}

	nsSdk := NSPlusSdk{}
	err := nsSdk.GetUserConfig()
	if err != nil {
		log.Printf("GetUserConfig error [%v]!!", err.Error())
		return nil, err
	}
	log.Printf("load UserConfig [%v]", nsSdk.UserConfig)

	nsSdk.AppName = appname ///nsSdk.Regfile.AppBase.BaseAppName
	///nsSdk.LogTarget = LogTarget
	///nsSdk.LogLevel = LogLevel

	//// containerNAme 就是 appmgt 安装app时 输入的app名称，
	containerNAme, err := LoadContainerNameFromAppCtrlTableFile()
	if err != nil {
		log.Printf("LoadContainerNameFromAppCtrlTableFile error [%v]!!", err.Error())

		nsSdk.ContainerName = nsSdk.AppName
		//return nil, err
	}else {
		nsSdk.ContainerName = containerNAme
	}

	profileName := ""  ///nsSdk.AppName + ".profile"
	err = nsSdk.LoadAppRegFiles()
	if err != nil {
		log.Printf("LoadAppRegFiles error [%v]!!", err.Error())
		//return nil, err
	}else{
		if nsSdk.Regfile.Config.ProfileName != "" {
			profileName = nsSdk.Regfile.Config.ProfileName
		}
		if nsSdk.ContainerName == "" {
			nsSdk.ContainerName = nsSdk.Regfile.AppBase.BaseAppName
		}
	}
	log.Printf("load regfile [%v]", nsSdk.Regfile)
	/// 两次赋值为空的话，容器名称 取初始化时的AppName ， 比如App以单独程序运行的时候
	if nsSdk.ContainerName == "" {
		nsSdk.ContainerName = nsSdk.AppName
	}
	/// 为空的话，profile 名称 取初始化时的AppName
	if profileName == "" {
		nsSdk.ContainerName = nsSdk.AppName + ".profile"
	}

	/// ==== 日志相关设置 ====
	if nsSdk.LogTarget == "" {
		nsSdk.LogTarget = fmt.Sprintf("/edgex/logs/edgex-app-%v.log", nsSdk.ContainerName) ///nsSdk.AppName)
	}
	if nsSdk.LogLevel == "" {
		nsSdk.LogLevel = "INFO"
	}
	nsSdk.SetSdkLogClient(nsSdk.ContainerName, nsSdk.LogTarget, nsSdk.LogLevel)

	nsSdk.MQConfig.Host = nsSdk.UserConfig.MQTTServer.Host
	nsSdk.MQConfig.Port = ZERO_Port
	nsSdk.MQConfig.Protocol = ZERO_Protocol
	nsSdk.MQConfig.Type = ZERO_Type
	nsSdk.MQConfig.Topic = ZERO_Topic

	nsSdk.InitHttpClient()


	nsSdk.ParamChange = make(chan AppParams, 1)


	//initDockerApiExecHttpClient()
	err = nsSdk.SdkDataIndexInit(profileName)
	for (err != nil) {
		//NsplusSdk.LoggingClient.Info(fmt.Sprintf("GetUserConfig error [%v], continue", erro.Error()))
		select {
		case e := <-messageErrors:
			nsSdk.LoggingClient.Error(fmt.Sprintf("exit msg: %s", e.Error()))
			if err != nil {
				nsSdk.LoggingClient.Error(fmt.Sprintf("with error: %s", err.Error()))
			}
			return nil, err
		case <-time.After(time.Second*5):
		}
		err = nsSdk.SdkDataIndexInit(profileName)
	}
	/////==================================
	///  nsSdk.ContainerName = nsSdk.Profile.Uuid   /// debug使用
	/// SdkDataIndexInit 函数加载了profile后，自动发起 数据集的创建和profile同步到nsfilter两个操作

	/// 自动创建数据集
	for _, oneInput := range nsSdk.Profile.Inputs {
		if oneInput.Type == AppType_DataSet {
			/// first, check the dataset is exist or not, if not, add it
			if !nsSdk.DoDatasetProc(oneInput.Name, oneInput.Description) {
				ErrMsg := fmt.Sprintf("配置App[%v]的数据集[%v]失败， ", nsSdk.ContainerName, oneInput)
				nsSdk.LoggingClient.Error(ErrMsg)
				///nari.HttpStrResult(w, r, http.StatusBadRequest, ErrMsg)
				continue
			}
			ErrMsg := fmt.Sprintf("配置App[%v]的数据集[%v] success ", nsSdk.ContainerName, oneInput)
			nsSdk.LoggingClient.Info(ErrMsg)
		}
	}

	err = nsSdk.SyncProfileToNsFilter(&nsSdk.Profile)
	for (err != nil) {
		//NsplusSdk.LoggingClient.Info(fmt.Sprintf("GetUserConfig error [%v], continue", erro.Error()))
		select {
		case e := <-messageErrors:
			nsSdk.LoggingClient.Error(fmt.Sprintf("exit msg: %s", e.Error()))
			if err != nil {
				nsSdk.LoggingClient.Error(fmt.Sprintf("with error: %s", err.Error()))
			}
			return nil, err
		case <-time.After(time.Second*5):
		}
		err = nsSdk.SyncProfileToNsFilter(&nsSdk.Profile)
	}

	nsSdk.MessageErrors, nsSdk.MessageEnvelopes, err = nsSdk.InitZMQClients()
	for (err != nil) {
		//NsplusSdk.LoggingClient.Info(fmt.Sprintf("GetUserConfig error [%v], continue", erro.Error()))
		select {
		case e := <-messageErrors:
			nsSdk.LoggingClient.Error(fmt.Sprintf("exit msg: %s", e.Error()))
			if err != nil {
				nsSdk.LoggingClient.Error(fmt.Sprintf("with error: %s", err.Error()))
			}
			return nil, err
		case <-time.After(time.Second*5):
		}
		nsSdk.MessageErrors, nsSdk.MessageEnvelopes, err = nsSdk.InitZMQClients()
	}

	nsSdk.MqttClient = nsSdk.newMqttSender()

	nsSdk.MqttMsgFunc = func(client MQTT.Client, msg MQTT.Message) {
		//LoggingClient.Error(fmt.Sprintf("TOPIC: [%s] \n", msg.Topic()))
		var Recvmsg mqttRecv
		Recvmsg.topic = msg.Topic()
		Recvmsg.msgBody = msg.Payload()
		nsSdk.LoggingClient.Info(fmt.Sprintf("recv TOPIC: [%s] Payload[%v] ", msg.Topic(), string(Recvmsg.msgBody)))

		////CmdRecvResponse <- &Recvmsg============  do response ================
		err := nsSdk.AppCmdProcesser(&Recvmsg)
		if err != nil {
			nsSdk.LoggingClient.Info(fmt.Sprintf("AppCmdProcesser for msg[%v] error [%v]!!", Recvmsg, err.Error()))
			//return nil
		}
	}

	///nsSdk.MqttClient.Recv(getTopic_InPut(nsSdk.ContainerName), nsSdk.MqttMsgFunc)
	////nsSdk.MqttClient.Recv(getTopic_OutPut(nsSdk.ContainerName), nsSdk.MqttMsgFunc)
	nsSdk.MqttClient.Recv(getTopic_Requtest(nsSdk.ContainerName), nsSdk.MqttMsgFunc)

	//err = nsSdk.SdkRun()
	//if err != nil {
	//	nsSdk.LoggingClient.Info(fmt.Sprintf("SdkRun fail, error [%v]!!", err.Error()))
	//	return nil, err
	//}

	///nsSdk.TimerGoRoutines()
	return &nsSdk, nil
}


func (sdk *NSPlusSdk)SdkRun() error {

	if sdk.InitClientsCBFunc != nil {
		err := sdk.InitClientsCBFunc()
		if err != nil{
			sdk.LoggingClient.Error(fmt.Sprintf("InitClientsCBFunc fail error=[%s] ", err.Error()))
			return err
		}
	}

	err := sdk.UpdateDataIndexMap()
	for (err != nil) {
		//NsplusSdk.LoggingClient.Info(fmt.Sprintf("GetUserConfig error [%v], continue", erro.Error()))
		select {
		case e := <-messageErrors:
			sdk.LoggingClient.Error(fmt.Sprintf("exit msg: %s", e.Error()))
			if err != nil {
				sdk.LoggingClient.Error(fmt.Sprintf("with error: %s", err.Error()))
			}
			return err
		case <-time.After(time.Second*5):
		}
		err = sdk.UpdateDataIndexMap()
	}
	sdk.LoggingClient.Info(fmt.Sprintf("load dataIndexMap [%v]", sdk.DataIndexMap))


	sdk.UpdateAppParamsCtrlMap()


	sdk.TimerGoRoutines()

	for {
		select {
		case e := <-sdk.MessageErrors:
			sdk.LoggingClient.Error(fmt.Sprintf("Get Error msg: %s", e.Error()))
			//return
		case msgEnvelope := <-sdk.MessageEnvelopes:
			sdk.LoggingClient.Info("Event received on message queue. %v ",  string(msgEnvelope.Payload))

			RecvData, err := sdk.RecvData(msgEnvelope)
			if err != nil{
				sdk.LoggingClient.Error(fmt.Sprintf("received on message queue fail from zmq[%s] ", msgEnvelope.CorrelationID))
				continue
			}

			sdk.LoggingClient.Info(fmt.Sprintf("Receive data [%v] ", RecvData))
			if RecvData == nil{
				sdk.LoggingClient.Error(fmt.Sprintf("received nil data for uuid[%s] ", sdk.Profile.Uuid))
				continue
			}
			/********** ==== 按需求添加数据输出流程 ====**********/

			if sdk.InputRecvCBFunc != nil {
				err = sdk.InputRecvCBFunc(RecvData)
				if err != nil {
					sdk.LoggingClient.Error(fmt.Sprintf("InputRecvCBFunc fail, from err[%s] ", err.Error()))
					///continue
				}
			}else {
				sdk.LoggingClient.Error(fmt.Sprintf("InputRecvCBFunc fail, func is null [%v] ", sdk.InputRecvCBFunc))
			}

		case ParamChan := <-sdk.ParamChange:
			sdk.LoggingClient.Info("Parameter received %v ",  ParamChan)


			err = sdk.ParameterRecvCBfunc(&ParamChan)
			if err != nil{
				sdk.LoggingClient.Error(fmt.Sprintf("Parameter change fail, from app error[%s] ", err.Error()))
				///continue
			}
		}
	}



	//return nil
}




func (sdk *NSPlusSdk)SdkDataIndexInit(profileName string) error {



	/**erro := GetUserConfig()
	for (erro != nil) {
		//LoggingClient.Info(fmt.Sprintf("GetUserConfig error [%v], continue", erro.Error()))
		select {
		case e := <-messageErrors:
			sdk.LoggingClient.Error(fmt.Sprintf("SdkInit: get gateway config fail, exit msg: %s", e.Error()))
			if erro != nil {
				sdk.LoggingClient.Error(fmt.Sprintf("SdkInit: get gateway config fail, with error: %s", erro.Error()))
			}
			//return
		case <-time.After(time.Second*10):
		}
		erro = GetUserConfig()
	}**/


	err := sdk.LoadProfiles(profileName) //// error
	if err != nil {
		sdk.LoggingClient.Info(fmt.Sprintf("SdkInit: LoadProfiles[%v] fail, errormsg = [%v] ", profileName, err.Error()))
		return err
	}

	sdk.InitDataIndexMap(sdk.ContainerName, sdk.Profile.TemplateVersion)

	return nil
}







