// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018 Canonical Ltd
// Copyright (C) 2018 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package nsplus

import (
	"encoding/json"
	"fmt"
)

type DataIndexForFilter struct{
	Name             string                   `json:"name"`          // Unique name for identifying a filter ， use profile uuid now， eg. export-iotmqtt
	Version          string                   `json:"version"`       // dev 或 res 有修改后，更新索引表，取时间串作为版本号， 发数据时带版本号使用指定索引表、同步时验证版本号是否相同
	IsValid          string                   `json:"isvalid"`       // 版本号变动时置0，同步确认后置1，约定置1前使用原索引关系，置1后使用新索引关系

	SyncTime         string                   `json:"synctime"`      // 同步标识，定时每30s同步一次索引，以同步时记录的时间戳为标识 另加通知驱动同步流程
	//Service        DeviceService            `json:"service"`       // Associated Device Service - One per device
	Index            []DataIndexByDevRes      `json:"index"`         /// Index 数组，所有dev和res组成的数组，索引取两个数值，devindex和resindex
}

type DataIndexByDevRes struct{
	DevName          string                   `json:"devname"`          // Unique name for identifying a device ， use as deviceSN
	Parent           string                   `json:"parent"`          // 父节点网关SN，网关SN默认全网唯一，与devSN一起唯一标识一个dev
	//Service        DeviceService            `json:"service"`       // Associated Device Service - One per device
	Resource         []ResourceInfo           `json:"resource"`      /// 参数列表，使用数组，取下标作为index使用
}

type ResourceInfo struct{
	Name             string                   `json:"name"`
	Type             string                   `json:"type"`
	MeasureType      string                   `json:"measuretype"`
	///ResourceIotName       string          `json:"ResourceIotName"`
	///BeNeedReport          bool            `json:"BeNeedReport"`
	Value            string                   `json:"type"`  /// 由数据发布方管理

	InputPara        string                   `json:"inputPara"` //// 所属的input数据集或设备资源参数
}


///========= 获取Appprofile后，初始化DataIndexMap，然后从filterSvr获取索引表详细内容 =========
func (sdk *NSPlusSdk)InitDataIndexMap(AppName string, version string){

	sdk.DataIndexMap = DataIndexForFilter{}
	sdk.DataIndexMap.Name = AppName
	sdk.DataIndexMap.Version = version
	sdk.DataIndexMap.Index = make([]DataIndexByDevRes, 0)

}

///========= 获取Appprofile后，初始化DataIndexMap，然后从filterSvr获取索引表详细内容 =========
///========= url := fmt.Sprintf("%s://%s:%v", c.Protocol, c.Host, c.Port)
func (sdk *NSPlusSdk)UpdateDataIndexMap() error {

	//return nil

	url := fmt.Sprintf("%s://%s:%v", Protocol_HTTP, sdk.UserConfig.MQTTServer.Host, PORT_FILTERMicSvr)
	sdk.LoggingClient.Error(fmt.Sprintf("UpdateDataIndexMap ： http url %v", url))

	res, err := sdk.HttpClient.Get(url + Api_Get_IndexMap + sdk.ContainerName)
	if err != nil {
		sdk.LoggingClient.Error(fmt.Sprintf("http.Get err %v", err.Error()))
		return err
	}

	if res.StatusCode != 200 {
		sdk.LoggingClient.Error(fmt.Sprintf("UpdateDataIndexMap http.Get StatusCode[%v], errorMsg[%v]", res.StatusCode, err.Error()))
		return err
	}

	var ack DataIndexForFilter
	err = json.NewDecoder(res.Body).Decode(&ack)
	if err != nil {
		sdk.LoggingClient.Error(fmt.Sprintf("Failed to get response for Device, error[%s]", err.Error()))
		//http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	sdk.LoggingClient.Error(fmt.Sprintf("get response for df[%v]， sdk.DataIndexMap[%v]", ack, sdk.DataIndexMap))

	//if ack.Version != sdk.DataIndexMap.Version {
		sdk.DataIndexMap = ack
		sdk.LoggingClient.Debug(fmt.Sprintf("GET DataIndexForFilter[%v]  ", sdk.DataIndexMap))
	//}else{
		//sdk.LoggingClient.Debug(fmt.Sprintf("GET the same version[%v] of DataIndexForFilter ", ack.Version))
	//}

	//LoggingClient.Debug(fmt.Sprintf(" GetNrProfileFromBackEnd get CurDevInfo.DevicePorfile count[%v] \n", len(CurDevInfo.DeviceList)))
	return nil
}


/*******************************************************
func GetNrProfileFromBackEnd() error {
	fmt.Fprintf(os.Stdout, "Configuration.Clients[ BackEnd ].Url()  %v ", Configuration.Clients[nari.ClientBackEnd].Url()+ nari.APIIOTDeviceProfile )
	res, err := http.Get(Configuration.Clients[nari.ClientBackEnd].Url() + nari.APIIOTDeviceProfile)
	if err != nil {
		fmt.Fprintf(os.Stdout, "http.Get err %v", err.Error())
		return err
	}
	defer res.Body.Close()

	var ack []contract.DeviceProfile
	err = json.NewDecoder(res.Body).Decode(&ack)
	if err != nil {
		LoggingClient.Error(fmt.Sprintf("Failed to get response for Device, error[%s]", err.Error()))
		//http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}

	LoggingClient.Debug(fmt.Sprintf("http.get DeviceProfile count[%v] \n", len(ack)))

	for _, Ied := range ack {
		index := 1 //getCurDevBufByProfileName(Ied.Name)
		if (len(index) > 0) {
			for _, ind := range index {
				//CurDevInfo.AppDeviceList = append(CurDevInfo.AppDeviceList, Ied.Name) ////.DeviceList[ind].ProfileForIOT = Ied
				log.Printf(" in getCurDevBufByProfileName set CurDevInfo.DeviceList[%v] ", ind)
			}
		}
	}
	//LoggingClient.Debug(fmt.Sprintf(" GetNrProfileFromBackEnd get CurDevInfo.DevicePorfile count[%v] \n", len(CurDevInfo.DeviceList)))
	return nil
}


func GetDevicesFromMetadata() (error){
	//fmt.Fprintf(os.Stdout, "Configuration.Clients[ Metadata ].Url()  %v ", Configuration.Clients[nari.ClientMetadata].Url() + nari.APIMETADATADevice)
	res, err := http.Get(Configuration.Clients[nari.ClientBackEnd].Url() + nari.APIMETADATADevice)
	if err != nil {
		LoggingClient.Error(fmt.Sprintf("http.Get  APIMETADATADevice err %v", err.Error()))
		return err
	}
	defer res.Body.Close()

	var devicelist []contract.Device

	err = json.NewDecoder(res.Body).Decode(&devicelist)
	if err != nil {
		LoggingClient.Error(fmt.Sprintf("Failed to get response for Device, error[%s]", err.Error()))
		//http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	//LoggingClient.Debug(fmt.Sprintf("get Devices count[%v]", len(devicelist)))

	//for _, dev := range devicelist {
		//bExist := false
		//if (getCurDevBufByDevSN(dev.Name) > -1){
		//	bExist = true
		//}
		//for _, dresin := range CurDevInfo.DeviceList {
		//	if(dresin.DeviceDefineInfo.Name == dev.Name){
		//		bExist = true
		//		break
		//	}
		//}
		***if(!bExist) {
			var devDefine DeviceInfo
			devDefine.DeviceDefineInfo.Name = dev.Name
			devDefine.DeviceDefineInfo.Port = dev.Service.Addressable.Port
			devDefine.DeviceDefineInfo.Address = dev.Service.Addressable.Address
			devDefine.DeviceDefineInfo.ProfileName = dev.Profile.Name

			devDefine.DeviceOnOff = "0"

			devDefine.ResourceNameType = make(map[string]ResourceNameTypeInfo)

			for _, dres := range dev.Profile.DeviceResources {
				var RsNameType ResourceNameTypeInfo
				RsNameType.ResourceName = dres.Name
				RsNameType.ResourceType = dres.Properties.Value.Type
				RsNameType.BeNeedReport = false
				RsNameType.ResourceIotName = dres.Name

				Type := nari.IOT_ATTR_YC //"analog"
				if(dres.Attributes["type"] == nari.Attr_YX_L) || (dres.Attributes["type"] == nari.Attr_YX_U) {
					Type = nari.IOT_ATTR_YX  //"discrete"
				}else if(dres.Attributes["type"] == nari.Attr_EVENT_L) || (dres.Attributes["type"] == nari.Attr_EVENT_U){
					Type = nari.IOT_ATTR_EVENT ///"discrete"
				}
				RsNameType.AttrType_serverId = Type

				///devDefine.DeviceDefineInfo.ResourceNameType = append(devDefine.DeviceDefineInfo.ResourceNameType, RsNameType)
				devDefine.ResourceNameType[dres.Name] = RsNameType
			}

			//log.Printf(" IOTMqtt: get CurDevInfo [%v] \n", devDefine)
			CurDevInfo.DeviceList = append(CurDevInfo.DeviceList, devDefine) /// first get the device define info

			////// 同时将core-backend生成的profile 拷贝到 iotconfig目录
			//

		}***
	//}

	////UpdateIndexMaps()

	return nil
}

*******************************************************/
