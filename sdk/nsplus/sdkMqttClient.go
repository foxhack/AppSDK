package nsplus

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"log"
)


type mqttSender struct {
	client MQTT.Client
	topic  string
}


func (sdk *NSPlusSdk)GetMqttBroker() string {
	return fmt.Sprintf("tcp://%v:%v", sdk.UserConfig.MQTTServer.Host, sdk.UserConfig.MQTTServer.Port) ////" + addr.Address + ":" + strconv.Itoa(addr.Port) + addr.Path
}
// newMqttSender - create new mqtt sender
func (sdk *NSPlusSdk)newMqttSender() *mqttSender {
	//protocol := strings.ToLower(addr.Protocol)

	opts := MQTT.NewClientOptions()
	broker := sdk.GetMqttBroker()  ///protocol + "://" + addr.Address + ":" + strconv.Itoa(addr.Port) + addr.Path
	opts.AddBroker(broker)
	opts.SetClientID("nsplus" + sdk.Profile.Uuid)
	opts.SetUsername(sdk.UserConfig.MQTTServer.User) ///Configuration.MQTTServer.User)
	opts.SetPassword(sdk.UserConfig.MQTTServer.Password) ///Configuration.MQTTServer.Password)
	opts.SetAutoReconnect(false)
	sdk.LoggingClient.Error(fmt.Sprintf("addr[%v] opts[%v]", broker, opts))

	/*if protocol == "tcps" || protocol == "ssl" || protocol == "tls" {
		cert, err := tls.LoadX509KeyPair(cert, key)

		if err != nil {
			LoggingClient.Error("Failed loading x509 data")
			return nil
		}

		tlsConfig := &tls.Config{
			ClientCAs:          nil,
			InsecureSkipVerify: true,
			Certificates:       []tls.Certificate{cert},
		}

		opts.SetTLSConfig(tlsConfig)

	}*/

	sender := &mqttSender{
		client: MQTT.NewClient(opts),
		topic:  getTopic_OutPut(sdk.ContainerName),
	}

	return sender
}

func (sender *mqttSender) Send(data []byte, topic string) bool {
	if !sender.client.IsConnected() {
		//LoggingClient.Info("Connecting to mqtt server")

		if token := sender.client.Connect(); token.Wait() && token.Error() != nil {
			//LoggingClient.Error(fmt.Sprintf("server: %v [%v]", sender, sender.client))
			log.Printf("Could not connect to mqtt server[%v], drop event. Error: %s", sender.client, token.Error().Error())
			return false
		}
	}
	if (topic == ""){
		topic = sender.topic
	}
	log.Printf("ready to publish to server: %v topic: %v \n", sender, topic)
	token := sender.client.Publish(topic, 0, false, data)
	// FIXME: could be removed? set of tokens?
	token.Wait()
	if token.Error() != nil {
		log.Printf("token.Wait( error: %v \n",token.Error().Error())
		return false
	} else {
		log.Printf("Sent data[%v]", string(data))
		return true
	}
}


func (sdk *NSPlusSdk)procSendData(topic string, req interface{}) error {

	log.Printf("procSendData [%v] to topic [%v]", req, topic)

	if (topic == "") {
		errMSg := fmt.Sprintf("procSendData topic=[%s] is not valid!!", topic)
		sdk.LoggingClient.Error(errMSg)
		//http.Error(w, err.Error(), http.StatusBadRequest)
		return errors.New(errMSg)
	}
	/////send
	formatted, err := json.Marshal(req)
	if err != nil {
		errMSg := fmt.Sprintf("Marshal [%v] to JSON faile. Error: %s", req, err.Error())
		sdk.LoggingClient.Error(errMSg)
		return err
	}


	if ! sdk.MqttClient.Send(formatted, topic) {
		sdk.LoggingClient.Debug(fmt.Sprintf("Failed to send date[%v]", string(formatted)))
		return err
	}

	sdk.LoggingClient.Error(fmt.Sprintf("send data msg[%v] succ.", req))
	///log.Printf("Get responseMSG from IOT [%v]\n ", ack)

	return nil
}

///// -------------receive topics of app response --------------------
type mqttRecv struct {
	topic  string
	msgBody    []byte
}


func (sender *mqttSender) Recv(topic string, f MQTT.MessageHandler) bool {
	if !sender.client.IsConnected() {
		//LoggingClient.Info("Connecting to mqtt server")
		//LoggingClient.Error(fmt.Sprintf("server: %v \n", sender))
		if token := sender.client.Connect(); token.Wait() && token.Error() != nil {
			//LoggingClient.Error(fmt.Sprintf("server: %v [[%v\n", sender, sender.client))
			log.Printf("Could not connect to mqtt server[%v], drop event. Error: %s", sender.client, token.Error().Error())
			return false
		}
	}
	///LoggingClient.Error(fmt.Sprintf("ready to Subscribe from server: %v \n", sender))
	token := sender.client.Subscribe(topic, 0, f)
	// FIXME: could be removed? set of tokens?
	token.Wait()
	if token.Error() != nil {
		log.Printf("Recv token.Wait( error: %v \n",token.Error().Error())
		return false
	} else {
		log.Printf("recv with token: %v \n", token)
		return true
	}
}

///========================= up =========================
func getTopic_InPut(AppKey string) string{
	topic := fmt.Sprintf("/%s/%s",INPUTTopic, AppKey)
	return topic
}

func getTopic_OutPut(AppKey string) string{
	topic := fmt.Sprintf("/%s/%s",, AppKey)
	return topic
}

func getTopic_Requtest(AppKey string) string{
	topic := fmt.Sprintf("/%s/%s",REQUESTTopic, AppKey)
	return topic
}

func getTopic_Response(AppKey string) string{
	topic := fmt.Sprintf("/%s/%s",RESPONSETopic, AppKey)
	return topic
}


func (sdk *NSPlusSdk)AppCmdProcesser(Recvmsg *mqttRecv) error {
	/// ================ 目前只有3个topic ==============
	//nsSdk.MqttClient.Recv(getTopic_InPut(nsSdk.Profile.Uuid))
	//nsSdk.MqttClient.Recv(getTopic_OutPut(nsSdk.Profile.Uuid))
	//nsSdk.MqttClient.Recv(getTopic_Requtest(nsSdk.Profile.Uuid))
	sdk.LoggingClient.Error(fmt.Sprintf("AppCmdProcesser: mqttRecv string[%v]", string(Recvmsg.msgBody)))

	var inteer interface{}
	errunw := json.Unmarshal(Recvmsg.msgBody, &inteer)
	if (errunw != nil) {
		sdk.LoggingClient.Error(fmt.Sprintf("Get mqttRecv err:[%v]", errunw.Error()))
	}
	sdk.LoggingClient.Error(fmt.Sprintf("Get mqttRecv msg[%v]", inteer))

	switch Recvmsg.topic {
	case getTopic_InPut(sdk.ContainerName):
		/// 新框架暂定以0mq收发数据， input 暂不处理
		///return fmt.Errorf("delete update not processed")
	case getTopic_OutPut(sdk.ContainerName):
		var outputPara OutputInterface

		errun := json.Unmarshal(Recvmsg.msgBody, &outputPara)
		if (errun != nil) {
			sdk.LoggingClient.Error(fmt.Sprintf("Get OutputInterface err:[%v]", errun.Error()))
		}
		sdk.LoggingClient.Error(fmt.Sprintf("Get OutputInterface outputPara[%v]", outputPara))
		/// add command readInputs process
		if(outputPara.Command == "readInputs"){
			sdk.LoggingClient.Info(fmt.Sprintf("recv command readInputs for[%v]", outputPara))





			var data []NameValues
			sdk.SendAppOutput(data)
		}
	case getTopic_Requtest(sdk.ContainerName):
		var resquestData requestWriteInterface //ResponseInterface
		err := json.Unmarshal(Recvmsg.msgBody, &resquestData)
		if (err != nil) {
			sdk.LoggingClient.Error(fmt.Sprintf("Unmarshal resquestData err:[%v]", err.Error()))
			return err
		}
		sdk.LoggingClient.Error(fmt.Sprintf("Get resquestData[%v]", resquestData))


		sdk.SendAppResponse(&resquestData)
			//LoggingClient.Error(fmt.Sprintf("Convert responseInterface.Sqid: [%v] => LastSqid[%v] in map [%v]", responseData.Sqid, appResp.LastSqid, TimeOutMap[appId]))


	case getTopic_Response(sdk.ContainerName):
		/// do nothing

	default:
		return fmt.Errorf(fmt.Sprintf("Invalid topic =[%v]", Recvmsg.topic))
	}

	return nil
}


func (sdk *NSPlusSdk)SendAppOutput(data []NameValues) error {
	/// ================ 目前只有3个topic ==============
	//nsSdk.MqttClient.Recv(getTopic_InPut(nsSdk.Profile.Uuid))
	//nsSdk.MqttClient.Recv(getTopic_OutPut(nsSdk.Profile.Uuid))
	//nsSdk.MqttClient.Recv(getTopic_Requtest(nsSdk.Profile.Uuid))

	var Outdata OutputInterface
	Outdata.Sqid = sdk.GetSqid()
	Outdata.Command = OutputWriteOutputCmd
	Outdata.Timedate = FormatStringForNow(time.Now())
	///Outdata.Outputs = data
	for _, onedat := range data {
		Outdata.Outputs = append(Outdata.Outputs, onedat)

		/// find output-device-res and save
		ResVal, ok := sdk.AppParamsCtrl.Output[onedat.Name]
		if ok {
			ResVal.Name = onedat.Name
			ResVal.Sqid = Outdata.Sqid
			ResVal.Values.Value = fmt.Sprintf("%v", onedat.Value)

			sdk.AppParamsCtrl.Output[onedat.Name] = ResVal
		}
	}


	err := sdk.procSendData(getTopic_OutPut(sdk.ContainerName), &Outdata) //error
	if (err != nil) {
		sdk.LoggingClient.Error(fmt.Sprintf("sdk.procSendData for Output data[%v] fail, err:[%v]", Outdata, err.Error()))
		return err
	}

	return nil
}


func (sdk *NSPlusSdk)SendAppResponse(req *requestWriteInterface) error {
	/// ================ 目前只有3个topic ==============
	//nsSdk.MqttClient.Recv(getTopic_InPut(nsSdk.Profile.Uuid))
	//nsSdk.MqttClient.Recv(getTopic_OutPut(nsSdk.Profile.Uuid))
	//nsSdk.MqttClient.Recv(getTopic_Requtest(nsSdk.Profile.Uuid))
	var responseData ResponseInterface

	responseData.Sqid = req.Sqid
	responseData.Status = "ACK"
	responseData.Timedate = FormatStringForNow(time.Now())

	if req.Command == RequestReadParamCmd {
		///  “readParameter”:需要读取的条目 name 数组，若为空， 表示全部读取
		if len(req.Arguments) == 0 {
			for _, param := range sdk.AppParamsCtrl.Parameter {
				var onepara NameValues
				onepara.Name = param.Name
				onepara.Value = param.Values.Value //fmt.Sprintf("%v", param.Values.Value)

				log.Printf("RequestReadParamCmd for responese [%v] ", onepara)
				responseData.Results = append(responseData.Results, onepara)
			}
		}else{
			for _, reqparam := range req.Arguments {
				/// find output-device-res and save
				ResVal, ok := sdk.AppParamsCtrl.Parameter[reqparam.Name]
				if ok {
					var onepara NameValues
					onepara.Name = reqparam.Name
					onepara.Value = ResVal.Values.Value

					log.Printf("RequestReadParamCmd for responese [%v] ", onepara)
					responseData.Results = append(responseData.Results, onepara)
				}
			}
		}

	}else if req.Command == RequestWriteParamCmd {
		for _, reqparam := range req.Arguments {
			/// find output-device-res and save
			ResVal, ok := sdk.AppParamsCtrl.Parameter[reqparam.Name]
			if ok {
				var onepara ValueDefine
				onepara.Value = fmt.Sprintf("%v", reqparam.Value)
				onepara.Type = ResVal.Values.Type

				ResVal.Values = onepara
				sdk.AppParamsCtrl.Parameter[reqparam.Name] = ResVal
				log.Printf("RequestReadParamCmd for responese [%v] into [%v]", onepara, ResVal)
			}
		}
	}else if req.Command == RequestReadOutputCmd {
		/// “readOutputs”:需要读取的条目 name 数组，若为空， 表示全部读
		if len(req.Arguments) == 0 {
			for _, param := range sdk.AppParamsCtrl.Output {
				var onepara NameValues
				onepara.Name = param.Name
				onepara.Value = param.Values.Value //fmt.Sprintf("%v", param.Values.Value)

				log.Printf("RequestReadParamCmd for Output [%v] ", onepara)
				responseData.Results = append(responseData.Results, onepara)
			}
		}else{
			for _, reqparam := range req.Arguments {
				/// find output-device-res and save
				ResVal, ok := sdk.AppParamsCtrl.Output[reqparam.Name]
				if ok {
					var onepara NameValues
					onepara.Name = reqparam.Name
					onepara.Value = ResVal.Values.Value

					log.Printf("RequestReadParamCmd for Output [%v] ", onepara)
					responseData.Results = append(responseData.Results, onepara)
				}
			}
		}
	}else if req.Command == ReloadProfileCmd {



	}


	err := sdk.procSendData(getTopic_Response(sdk.ContainerName), &responseData) //error
	if (err != nil) {
		sdk.LoggingClient.Error(fmt.Sprintf("sdk.procSendData for Output data[%v] fail, err:[%v]", responseData, err.Error()))
		return err
	}

	return nil
}



