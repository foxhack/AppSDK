// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018 Canonical Ltd
// Copyright (C) 2018 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package nsplus

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/edgexfoundry/nsplussdk/sdk/nrprivate"
	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	msgTypes "github.com/edgexfoundry/go-mod-messaging/pkg/types"
)

type DataDistroWithIndex struct {
	Id 			string               `json:"id" `
	Created     int64                `json:"created" ` // When the reading was created
	Device      string               `json:"device" `
	DeviceName  string               `json:"devicename" `

	// 索引表名称，暂定用profileName，用于接收时验证数据是否当前App的数据，sdk提供验证接口
	IndexName   string               `json:"IndexName" `
	// dev 或 res 有修改后，更的索引表，取时间串作为版本号， 发数据时带版本号使用指定索引表、同步时验证版本号是否相同
	Version     string               `json:"version"`

	DevIndex    int                  `json:"devIndex"`/// DataIndexForFilter.Index[]的下标

	NrPayload   []ResValue           `json:"nrpayload"`
}

type  ResValue struct {
	Index       int                  `json:"resIndex"`/// DataIndexForFilter.Index[]的下标
	Value       nrprivate.NrValue    `json:"value" `
}

func NewDataDistroWithIndex() *DataDistroWithIndex {
	ployload := DataDistroWithIndex{}
	ployload.NrPayload = make([]ResValue,0)

	return &ployload
}

func (sdk *NSPlusSdk)RecvData(recData msgTypes.MessageEnvelope) (*DataDistroWithIndex, error){

	RecvData, err := sdk.ParseRecvData(recData)
	if err != nil {
		errMsg := fmt.Sprintf("RecvData: ParseRecvData[%v] fail, error[%v]", recData.Payload, err.Error())
		sdk.LoggingClient.Error(errMsg)
		return nil, err
	}

	if !sdk.CheckDataRecvForApp(RecvData) {
		errMsg := fmt.Sprintf("RecvData: CheckDataRecvForApp[%v] fail, error", RecvData)
		sdk.LoggingClient.Error(errMsg)

		err = sdk.UpdateDataIndexMap()
		if (err != nil) {
			errMsg := fmt.Sprintf("RecvData: CheckDataRecvForApp fail to update Indexs fail, error[%v]", err.Error())
			sdk.LoggingClient.Error(errMsg)
		}

		return nil, err
	}


	/// 更新收到的input 数据的值
	for _, ResV := range RecvData.NrPayload {
		/// find input name
		inkey, ok := sdk.AppParamsCtrl.IndexForInput["/"+RecvData.Device + "/" +ResV.Value.Name]
		if !ok {
			break
		}
		/// find input
		input, ok := sdk.AppParamsCtrl.Input[inkey]
		if !ok {
			break
		}
		/// find input-device
		inputDevVal, ok := input.DevResValues[RecvData.Device]
		if !ok {
			break
		}
		/// find input-device-res
		inputDevResVal, ok := inputDevVal.ResValues[ResV.Value.Name]
		if !ok {
			break
		}
		/// save to buf  map
		inputDevResVal.Value = ResV.Value.Value
		inputDevResVal.Type = ResV.Value.DataType

		inputDevVal.ResValues[ResV.Value.Name] = inputDevResVal
		input.DevResValues[RecvData.Device] = inputDevVal
		sdk.AppParamsCtrl.Input[input.Name] = input
	}

	return RecvData, nil
}

func (sdk *NSPlusSdk)ParseRecvData(recvData msgTypes.MessageEnvelope) (*DataDistroWithIndex, error){

	if recvData.ContentType != clients.ContentTypeJSON {
		errMsg := fmt.Sprintf("Incorrect content type: Received: %s, Expected: %s", recvData.ContentType, clients.ContentTypeJSON)
		sdk.LoggingClient.Error(errMsg)
		return nil,errors.New(errMsg)
	}
	str := string(recvData.Payload)
	sdk.LoggingClient.Error(fmt.Sprintf("Receive data %s ", str))

	data := DataDistroWithIndex{}

	if err := json.Unmarshal(recvData.Payload, &data); err != nil {
		sdk.LoggingClient.Error(err.Error())
		return nil, nil
	}

	if len(data.NrPayload) <= 0 {
		errMsg := fmt.Sprintf("Received: %s, but value count is 0", data)
		sdk.LoggingClient.Error(errMsg)
		return nil,errors.New(errMsg)
	}

	return &data, nil
}

///==== 获取数据后，判断是否当前App的数据，判断版本是否对应，判断Index值是否有效 ====
func (sdk *NSPlusSdk)CheckDataRecvForApp(recvData *DataDistroWithIndex) bool {

	if recvData.IndexName != sdk.DataIndexMap.Name {
		errMsg := fmt.Sprintf("CheckDataRecvForApp: receive Data [%v]，but not for App [%v]", recvData, sdk.Profile.Uuid)
		sdk.LoggingClient.Error(errMsg)
		return false
	}

	if recvData.Version != sdk.DataIndexMap.Version {
		errMsg := fmt.Sprintf("CheckDataRecvForApp: receive Data [%v]，but Index version not valid[%v], get IndexMap again now", recvData, sdk.Profile.TemplateVersion)
		sdk.LoggingClient.Error(errMsg)
		return false
	}

	if recvData.DevIndex >= len(sdk.DataIndexMap.Index) {
		errMsg := fmt.Sprintf("CheckDataRecvForApp: receive devIndex[%v]，but Index Count[%v], get IndexMap again now", recvData, len(sdk.DataIndexMap.Index) )
		sdk.LoggingClient.Error(errMsg)
		return false
	}

	for _, ResV := range recvData.NrPayload {
		if ResV.Index >= len(sdk.DataIndexMap.Index[recvData.DevIndex].Resource) {
			errMsg := fmt.Sprintf("CheckDataRecvForApp: receive devIndex[%v] ResV.Index[%v]，but resIndex Count[%v], get IndexMap again now", recvData, ResV.Index, len(sdk.DataIndexMap.Index[recvData.DevIndex].Resource))
			sdk.LoggingClient.Error(errMsg)
			return false
		}
	}

	return true
}