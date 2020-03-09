//
// Copyright (c) 2017 Cavium
//
// SPDX-License-Identifier: Apache-2.0
//

package nsplus

import (
	"encoding/json"
	"fmt"
	"github.com/edgexfoundry/nsplussdk/sdk/nrprivate"

	"io/ioutil"
	"log"
	"strings"
)

const (

	ContainerCtrlDBTABLE = "appctrltable.json"
	//ContainerCtrlDBTABLE = "ctrlinfo.json"
)


type CtrlDockerGuarderStruct struct {
	MQTTServer      interface{}            `json:"MQTTServer"`
	Container       ContainerConfig        `json:"Container"`
	AppCtrlInfo     interface{}            `json:"AppCtrlInfo"`
}

///Container info
type ContainerConfig struct{
	ContainerName       string    `json:"ContainerName"`
}

///读取原有App设置后，写入 APPCtrlTABLE
func LoadContainerNameFromAppCtrlTableFile()(string, error) {
	var AppCtrlTable CtrlDockerGuarderStruct

	path := "./private/" //nari.GetIotAppPrivatePath(ContainerName)
	_, err1 := ioutil.ReadDir(path)
	if err1 != nil {
		log.Printf("LoadAppCtrlTableFile: couldn't read directory: %s; %v", path, err1.Error())
		return "", err1
	}

	fName := APPCtrlTABLE
	err := nrprivate.LoadJsonFromFile(path, fName , &AppCtrlTable) ///+ ".profile"
	if err != nil {
		errMsg := fmt.Sprintf("LoadAppCtrlTableFile: container path [%v] fail， error[%v]", path, err.Error())
		log.Printf(errMsg)
		return "", err
	}

	lfName := strings.ToLower(fName)
	//if (lfName == APPDBTABLE) { ///strings.HasSuffix(lfName, "data") || strings.HasSuffix(lfName, "dat") {
	fullPath := path + lfName
	appFile, err := ioutil.ReadFile(fullPath)
	if err != nil {
		log.Printf("LoadAppCtrlTableFile: couldn't read file: %s; %v", fullPath, err.Error())
		return "", err
	}

	err = json.Unmarshal(appFile, &AppCtrlTable)
	if err != nil {
		log.Printf("invalid APPDBTABLEFILE: %s; %v", fullPath, err.Error())
		return "", err
	}
	log.Printf("LoadAppCtrlTableFile: %v \n", AppCtrlTable)
	//}
	return AppCtrlTable.Container.ContainerName, nil
}

