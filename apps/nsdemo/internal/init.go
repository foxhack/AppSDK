//
// Copyright (c) 2018 Tencent
// Copyright (c) 2019 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package internal

import (
	"github.com/edgexfoundry/nsplussdk/sdk/nsplus"
	"github.com/edgexfoundry/nsplussdk/sdk/telemetry"
)

/// var LoggingClient logger.LoggingClient
/// var ec coredata.EventClient
//var Configuration *ConfigurationStruct
//var registryClient registry.Client
//var registryErrors chan error        //A channel for "config wait errors" sourced from Registry
//var registryUpdates chan interface{} //A channel for "config updates" sourced from Registry

///var messageErrors chan error
//var messageEnvelopes chan msgTypes.MessageEnvelope   //chan *msgTypes.MessageEnvelope

var AppNameKey string

var NsplusSdk *nsplus.NSPlusSdk



func Init() bool {

	go telemetry.StartCpuUsageAverage()

	return true
}

func Destruct() {

	if NsplusSdk != nil {
		close(NsplusSdk.MessageEnvelopes)
		close(NsplusSdk.MessageErrors)

		NsplusSdk = nil
	}
}


