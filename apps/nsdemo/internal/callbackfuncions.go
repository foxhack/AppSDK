//
// Copyright (c) 2018 Tencent
// Copyright (c) 2019 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package internal

import (
	"fmt"
	"github.com/edgexfoundry/nsplussdk/sdk/nsplus"
	"time"
)


/****************SDK里定义一个函数指针类型
type InputRecvCallback func(recData *DataDistroWithIndex) error

type ParameterRecvCallback func(recvParam *AppParams) error

type InitClientsCallback func() error
********************************************************/


func InputCBFuns(recvParam *nsplus.DataDistroWithIndex) error {

	NsplusSdk.LoggingClient.Error(fmt.Sprintf("nnnnnnnnnnn InputCBFuns recv data[%v]", recvParam))
	//close(errChan)
	return nil
}

func ParameterCBFuns(recvParam *nsplus.AppParams) error {

	NsplusSdk.LoggingClient.Error(fmt.Sprintf("mmmmmmmmmmm ParameterCBFuns recv data[%v]", recvParam))
	//close(errChan)
	return nil
}

func InitClientsCBFuns() error {

	NsplusSdk.LoggingClient.Error(fmt.Sprintf("kkkkkkkkkk InitClientsCBFuns now time=[%v]", time.Now()))
	//close(errChan)
	return nil
}

