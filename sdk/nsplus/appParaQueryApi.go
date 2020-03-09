// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018 Canonical Ltd
// Copyright (C) 2018 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package nsplus

import (
	"log"
)

type AppInputs struct {
	Param        InputsInfo           `json:"param"`        // input\output parameter define
	DevResValues map[string]ResValues `json:"devResValues"` // key devname, value res name and values
}

type AppOutputs struct {
	Param  OutputsInfo `json:"param"`  //  output define
	Values ValueDefine `json:"values"` //  value  and type
}

type AppParams struct {
	Param  ParameterInfo `json:"param"`  //  parameter define
	Values ValueDefine   `json:"values"` //  value  and type
}

//type ResValues struct{
//	ResValues        map[string]ValueDefine   `json:"resValues"`     // key resname, value res type and values
//}
//type ValueDefine struct{
//	Value            string                   `json:"value"`
//	Type             string                   `json:"type"`
//}

///========= 获取App parameter 定义和值， pName为空的时候取所有的parameters =========
func (sdk *NSPlusSdk) GetInput(pName string) []AppInputs {
	log.Printf("GetInput for input[%v] ", pName)
	var appinputs []AppInputs
	///=== 	 inputs
	bFind := false
	for _, input := range sdk.Profile.Inputs {
		if pName != "" {
			if bFind {
				break
			}
			if input.Name != pName {
				continue
			} else {
				bFind = true
			}
		}

		var oneinput AppInputs
		oneinput.Param = input
		oneinput.DevResValues = make(map[string]ResValues, 0)
		/// find input
		inputVal, ok := sdk.AppParamsCtrl.Input[input.Name]
		if !ok {
			continue
		}
		for devKey, inputResVal := range inputVal.DevResValues {
			log.Printf("GetInput for input get Dev[%v] ResValues[%v] ", devKey, inputResVal)
			if devKey != "" {
				oneinput.DevResValues[devKey] = inputResVal
			}
		}

		appinputs = append(appinputs, oneinput)
	}
	log.Printf("GetInput for input get one input[%v] ", appinputs)
	return appinputs
}

///========= 获取App parameter 定义和值， pName为空的时候取所有的parameters =========
func (sdk *NSPlusSdk) GetOutput(pName string) []AppOutputs {

	log.Printf("GetOutput for Output[%v] ", pName)
	var apppara []AppOutputs
	///=== 	 inputs
	bFind := false
	for _, para := range sdk.Profile.Outputs {
		if pName != "" {
			if bFind {
				break
			}
			if para.Name != pName {
				continue
			} else {
				bFind = true
			}
		}
		var onepara AppOutputs
		onepara.Param = para
		/// find input
		paraVal, ok := sdk.AppParamsCtrl.Output[para.Name]
		if !ok {
			continue
		}

		log.Printf("GetOutput for input get Dev[%v] ResValues[%v] ", para.Name, paraVal)
		onepara.Values = paraVal.Values
		apppara = append(apppara, onepara)
	}

	return apppara
}

///========= 获取App parameter 定义和值， pName为空的时候取所有的parameters =========
func (sdk *NSPlusSdk) GetParams(pName string) []AppParams {

	log.Printf("GetParams for Output[%v] ", pName)
	var apppara []AppParams
	///=== 	 inputs
	bFind := false
	for _, para := range sdk.Profile.Parameter {
		if pName != "" {
			if bFind {
				break
			}
			if para.Name != pName {
				continue
			} else {
				bFind = true
			}
		}
		var onepara AppParams
		onepara.Param = para
		log.Printf("GetParams for parameter get Dev[%v] para[%v] ", para.Name, para)
		/// find input
		paraVal, ok := sdk.AppParamsCtrl.Parameter[para.Name]
		if !ok {
			continue
		}

		log.Printf("GetParams for parameter get Dev[%v] ResValues[%v] ", para.Name, paraVal)
		onepara.Values = paraVal.Values
		apppara = append(apppara, onepara)
	}

	return apppara
}

///========= 获取App Port 定义和值 =========
func (sdk *NSPlusSdk) GetPortConfig() int {
	log.Printf("GetPortConfig port[%v] ", sdk.Profile.Config)
	return sdk.Profile.Config.Port
}

///========= 获取App 业务路径 =========
func (sdk *NSPlusSdk) GetDirConfig() []string {
	log.Printf("GetDirConfig port[%v] ", sdk.Profile.Config)
	return sdk.Profile.Config.Mount
}
