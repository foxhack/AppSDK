// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018 Canonical Ltd
// Copyright (C) 2018 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package nsplus

import (
	"fmt"
	"log"
)

func (sdk *NSPlusSdk) DoDatasetProc(datasetName string, desc string) bool {

	if sdk.checkDatasetByName(datasetName) {
		sdk.LoggingClient.Error(fmt.Sprintf("DoDatasetProc check for dataset[%v] existed", datasetName))
		return true
	}

	err := sdk.AddNilDatasetByName(sdk.ContainerName, datasetName, desc)

	if err != nil {
		sdk.LoggingClient.Error(fmt.Sprintf("DoDatasetProc AddNilDatasetByName for app[%v] dataset[%v] fail, error=[%v]", sdk.ContainerName, datasetName, err.Error()))
		return false
	}
	log.Printf("DoDatasetProc: add dataSet[%v] for app[%v] succ", datasetName, sdk.ContainerName)

	return true
}

func (sdk *NSPlusSdk) checkDatasetByName(datasetName string) bool {
	dataSet, err := sdk.GetDatasetByName(datasetName)
	if err != nil {
		sdk.LoggingClient.Error(fmt.Sprintf("checkDatasetByName for dataset[%v], error[%v]", datasetName, err.Error()))
		return false
	}
	if dataSet == nil {
		sdk.LoggingClient.Error(fmt.Sprintf("checkDatasetByName for dataset[%v] get nil set", datasetName))
		return false
	}
	log.Printf("checkDatasetByName === dataSet[%v]", dataSet)

	return true
}

/// ====================================================================
