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

type AppParamCtrlBuf struct{
	Input            map[string]AppInput           `json:"input"`
	IndexForInput    map[string]string             `json:"indexForInput"`
	Output           map[string]AppOutputParams    `json:"output"`
	Parameter        map[string]AppOutputParams    `json:"parameter"`

	InputSqid        int                           `json:"inputSqid"`
	OutputSqid       int                           `json:"outputSqid"`
	//OutputSqid       int                           `json:"outputSqid"`
}

type IndexForInputDef struct{
	input           string                   `json:"input"`
}

type AppInput struct{
	Name             string                   `json:"name"`          // input\output parameter name
	Sqid             int                      `json:"sqid"`
	DevResValues     map[string]ResValues     `json:"devResValues"`  // key devname, value res name and values
}

type ResValues struct{
	ResValues        map[string]ValueDefine   `json:"resValues"`     // key resname, value res type and values
}

type ValueDefine struct{
	Value            string                   `json:"value"`
	Type             string                   `json:"type"`
}

type AppOutputParams struct{
	Name             string                   `json:"name"`          // input\output parameter name
	Sqid             int                      `json:"sqid"`
	Values           ValueDefine              `json:"values"`  // key devname, value res name and values
}


///========= 获取Appprofile后，初始化DataIndexMap，然后从filterSvr获取索引表详细内容 =========
func (sdk *NSPlusSdk)UpdateAppParamsCtrlMap() {

	sdk.AppParamsCtrl = AppParamCtrlBuf{}
	sdk.AppParamsCtrl.Input = make(map[string]AppInput)
	sdk.AppParamsCtrl.IndexForInput =make(map[string]string)
	sdk.AppParamsCtrl.Output = make(map[string]AppOutputParams)
	sdk.AppParamsCtrl.Parameter = make(map[string]AppOutputParams)

	///=== 	 inputs
	for _, input := range sdk.Profile.Inputs {
		var oneinput AppInput
		oneinput.Name = input.Name
		oneinput.Sqid = 0
		oneinput.DevResValues = make(map[string]ResValues)
		log.Printf("UpdateAppParamsCtrlMap for input with input[%v] ", input.Name)
		for _, devs := range sdk.DataIndexMap.Index {
			//CurDevInfo.AppDeviceList = append(CurDevInfo.AppDeviceList, Ied.Name) ////.DeviceList[ind].ProfileForIOT = Ied
			log.Printf("UpdateAppParamsCtrlMap for input with dev[%v] ", devs)

			var ResVal ResValues
			ResVal.ResValues = make(map[string]ValueDefine)

			for _, res := range devs.Resource {

				var resDef ValueDefine
				resDef.Value = res.Value
				resDef.Type = res.Type

				ResVal.ResValues[res.Name] = resDef

				sdk.AppParamsCtrl.IndexForInput["/"+devs.DevName+"/"+res.Name] = input.Name
			}
			log.Printf("UpdateAppParamsCtrlMap for input get DevResValues[%v] ", ResVal)
			oneinput.DevResValues[devs.DevName] = ResVal
		}
		log.Printf("UpdateAppParamsCtrlMap for input get one input[%v] ", oneinput)
		sdk.AppParamsCtrl.Input[input.Name] = oneinput
	}


	// === output ===
	for _, param := range sdk.Profile.Outputs {
		var onepara AppOutputParams
		onepara.Name = param.Name
		onepara.Sqid = 0
		onepara.Values.Value = "-"
		onepara.Values.Type = param.Type

		log.Printf("UpdateAppParamsCtrlMap for output [%v] ", onepara)
		sdk.AppParamsCtrl.Output[param.Name] = onepara
	}
	// === parameter ===
	for _, param := range sdk.Profile.Parameter {
		var onepara AppOutputParams
		onepara.Name = param.Name
		onepara.Sqid = 0
		onepara.Values.Value = fmt.Sprintf("%v", param.DefaultValue)
		onepara.Values.Type = param.Type

		log.Printf("UpdateAppParamsCtrlMap for parameter [%v] ", onepara)
		sdk.AppParamsCtrl.Parameter[param.Name] = onepara
	}
}



func (sdk *NSPlusSdk)GetSqid(ptype string)int {
	if ptype == OUTPUTTopic {
		sdk.AppParamsCtrl.OutputSqid = sdk.AppParamsCtrl.OutputSqid + 1
		if sdk.AppParamsCtrl.OutputSqid > 65534 {
			sdk.AppParamsCtrl.OutputSqid = 1
		}
		return sdk.AppParamsCtrl.OutputSqid
	}else if ptype == INPUTTopic {
		sdk.AppParamsCtrl.OutputSqid = sdk.AppParamsCtrl.OutputSqid + 1
		return sdk.AppParamsCtrl.InputSqid
	}

	return 1010
}


