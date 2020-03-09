//
// Copyright (c) 2018 Tencent
// Copyright (c) 2019 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package nsplus

import (
	"fmt"
	"log"
	"time"
)

func (sdk *NSPlusSdk)TimerGoRoutines(){
	///checkEdgeAuthStatusLoop()

	sdk.debugParameterChangeSendLoop()

	sdk.debugOutputSendLoop()

	///CheckDevicesOnOffStatusLoog()

}



/// ==== 定时检查devices是否接入 ==== 达到即插即用需求
func (sdk *NSPlusSdk)debugParameterChangeSendLoop() {
	go func() {
		for {
			/// 判断网关是否已接入

			//var Outdata OutputInterface
			//Outdata.Sqid = sdk.GetSqid(OUTPUTTopic)
			//Outdata.Command = OutputWriteOutputCmd
			//Outdata.Timedate = FormatStringForNow(time.Now())

			///var apppara []AppParams
			///=== 	 inputs
			//bFind := false
			for _, para := range sdk.Profile.Parameter {

				var onepara AppParams
				onepara.Param = para
				log.Printf("GetParams for parameter get Dev[%v] para[%v] ", para.Name, para)
				/// find input
				paraVal, ok := sdk.AppParamsCtrl.Parameter[para.Name]
				if !ok {
					continue
				}


				paraVal.Values.Value = fmt.Sprintf("%v", RandInt(0,10))
				onepara.Values = paraVal.Values
				//apppara = append(apppara, onepara)
				log.Printf("GetParams for parameter get Dev[%v] ResValues[%v] ", para.Name, paraVal)

				sdk.ParamChange <- onepara
			}



			////////////////////////////////
			now := time.Now() // 计算下一个零点
			next := now.Add(time.Second * 30)   //// time.Second * 180
			next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), next.Second(), 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}



/// ==== 定时检查devices是否接入 ==== 达到即插即用需求
func (sdk *NSPlusSdk)debugOutputSendLoop() {
	go func() {
		for {
			/// 判断网关是否已接入

			//var Outdata OutputInterface
			//Outdata.Sqid = sdk.GetSqid(OUTPUTTopic)
			//Outdata.Command = OutputWriteOutputCmd
			//Outdata.Timedate = FormatStringForNow(time.Now())
			Outputs := make([]NameValues, 0)
			for _, param := range sdk.AppParamsCtrl.Output {
				var onepara NameValues
				onepara.Name = param.Name
				onepara.Value = fmt.Sprintf("%v_%v", "outValue", sdk.AppParamsCtrl.OutputSqid)

				log.Printf("debugOutputSendLoop for Output [%v] ", onepara)
				Outputs = append(Outputs, onepara)
			}

			err := sdk.SendAppOutput(Outputs)
			if err != nil {
				log.Printf("SendAppOutput for Output [%v] fail, error=[%v] ", Outputs, err.Error())
			}

			////////////////////////////////
			now := time.Now() // 计算下一个零点
			next := now.Add(time.Second * 60)   //// time.Second * 180
			next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), next.Second(), 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}
