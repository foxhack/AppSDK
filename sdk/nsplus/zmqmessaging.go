//
// Copyright (c) 2019 Intel Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package nsplus

import (
	"fmt"
	"github.com/edgexfoundry/go-mod-messaging/messaging"
	msgTypes "github.com/edgexfoundry/go-mod-messaging/pkg/types"
)
var MsgClient messaging.MessageClient

func (sdk *NSPlusSdk)InitZMQClients()(chan error, chan msgTypes.MessageEnvelope, error) {
	//var client messaging.MessageClient
	// Create the messaging client
	var err error
	MsgClient, err = messaging.NewMessageClient(msgTypes.MessageBusConfig{
		SubscribeHost: msgTypes.HostInfo{
			Host:     sdk.UserConfig.MQTTServer.Host, ///..MQConfig.Host, //Configuration.MessageQueue.Host,
			Port:     sdk.MQConfig.Port, //Configuration.MessageQueue.Port,
			Protocol: sdk.MQConfig.Protocol,  //Configuration.MessageQueue.Protocol,
		},
		Type: sdk.MQConfig.Type, //Configuration.MessageQueue.Type,
	})

	if err != nil {
		sdk.LoggingClient.Error("failed to create messaging client: " + err.Error())
		return nil, nil, fmt.Errorf("failed to create messaging client:: %s", err.Error())
	}
	sdk.LoggingClient.Info(fmt.Sprintf("create messaging client[%v]", sdk.MQConfig))

	if err := MsgClient.Connect(); err != nil {
		return nil, nil, fmt.Errorf("failed to connect to message bus: %s ", err.Error())
	}

	errs := make(chan error, 2)
	messages := make(chan msgTypes.MessageEnvelope, 1)

	topics := []msgTypes.TopicChannel{
		{
			Topic:    sdk.MQConfig.Topic, // Configuration.MessageQueue.Topic,
			Messages: messages,
		},
	}

	sdk.LoggingClient.Info("Connecting to incoming message bus at: " + sdk.MQConfig.Uri() ) ///Configuration.MessageQueue.Uri())

	err = MsgClient.Subscribe(topics, errs)
	if err != nil {
		//close(errs)
		close(messages)
		return nil, nil, fmt.Errorf("failed to subscribe for event messages: %s", err.Error())
	}

	sdk.LoggingClient.Info("Connected to inbound event messages for topic: " + sdk.MQConfig.Topic)

	return errs, messages, nil
}
/********8
func initMessaging(Uri string, topic string, client messaging.MessageClient) (chan error, chan msgTypes.MessageEnvelope, error) {
	if err := client.Connect(); err != nil {
		return nil, nil, fmt.Errorf("failed to connect to message bus: %s ", err.Error())
	}

	errs := make(chan error, 2)
	messages := make(chan msgTypes.MessageEnvelope, 10)

	topics := []msgTypes.TopicChannel{
		{
			Topic:    topic, // Configuration.MessageQueue.Topic,
			Messages: messages,
		},
	}

	LoggingClient.Info("Connecting to incoming message bus at: " + Uri ) ///Configuration.MessageQueue.Uri())

	err := client.Subscribe(topics, errs)
	if err != nil {
		close(errs)
		close(messages)
		return nil, nil, fmt.Errorf("failed to subscribe for event messages: %s", err.Error())
	}

	LoggingClient.Info("Connected to inbound event messages for topic: " + topic)

	return errs, messages, nil
}**************/
