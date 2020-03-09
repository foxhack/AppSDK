package nsplus

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/edgexfoundry/nsplussdk/sdk/define"
	"github.com/edgexfoundry/nsplussdk/sdk/nrprivate"
	"io/ioutil"
	"net/http"
)


func (sdk *NSPlusSdk)InitHttpClient() {
	///
	sdk.HttpClient = nrprivate.NewHttpClient()
}


func (sdk *NSPlusSdk)CallAllDataForIndexTable() []DataDistroWithIndex {
	///http://192.168.43.129:5555/api/v1/readvalue/device/Random05
	if len(sdk.DataIndexMap.Index) == 0 {
		errMsg := fmt.Sprintf("CallAllDataForIndexTable: sdk.DataIndexMap.Index is null， [%v]", sdk.DataIndexMap.Index)
		sdk.LoggingClient.Error(errMsg)
		return nil
	}
	var DataWithIndex []DataDistroWithIndex
	for devIndex, dev := range sdk.DataIndexMap.Index {
		if dev.DevName == "" {
			errMsg := fmt.Sprintf("for sdk.DataIndexMap.Index dev[%v] devname is null ", dev)
			sdk.LoggingClient.Error(errMsg)
			continue
		}
		DevRes, err := sdk.GetDevResReadings(dev.DevName)
		if err != nil {
			errMsg := fmt.Sprintf("CallAllDataForIndexTable: 获取设备[%v]的readings失败， error[%v]", dev.DevName, err.Error())
			sdk.LoggingClient.Error(errMsg)
			continue
		}
		errMsg := fmt.Sprintf("CallAllDataForIndexTable: 获取设备[%v]的readings， DevRes[%v]", dev.DevName, DevRes)
		sdk.LoggingClient.Error(errMsg)

		var devDataWithIndex DataDistroWithIndex
		devDataWithIndex.DevIndex = devIndex
		devDataWithIndex.Device = dev.DevName
		devDataWithIndex.IndexName = sdk.DataIndexMap.Name
		devDataWithIndex.Version = sdk.DataIndexMap.Version


		for resindex, readings := range dev.Resource {
			var ResReading ResValue
			ResReading.Index = resindex

			Data, ok := DevRes[readings.Name]
			if !ok {
				errMsg := fmt.Sprintf("CallAllDataForIndexTable: not find reading for dev[%v] res[%v]", dev.DevName, readings.Name)
				sdk.LoggingClient.Error(errMsg)
				continue
			}

			ResReading.Value.Name = Data.Name
			ResReading.Value.Value = Data.Value
			ResReading.Value.Description = Data.Description
			ResReading.Value.Measuretype = Data.Measuretype
			ResReading.Value.DataType = Data.DataType
			ResReading.Value.Unit = Data.Unit

			devDataWithIndex.NrPayload = append(devDataWithIndex.NrPayload, ResReading)
		}

		DataWithIndex = append(DataWithIndex, devDataWithIndex)

	}

	return DataWithIndex
}


/// =========== 取设备 reading 列表 ===========
func (sdk *NSPlusSdk)GetDevResReadings(DevName string) (map[string]*nrprivate.DeviceValue, error){
	//fmt.Fprintf(os.Stdout, "Configuration.Clients[ Metadata ].Url()  %v ", Configuration.Clients[nari.ClientMetadata].Url() + nari.APIMETADATADevice)
	httpPath := fmt.Sprintf("http://%v:%v/api/v1/readvalue/device/%v", sdk.UserConfig.MQTTServer.Host, Backend_Port, DevName)

		res, err := sdk.HttpClient.Get(httpPath)
		if err != nil {
			errMsg := fmt.Sprintf("GetDevResReadings: 获取设备[%v]的readings失败， error[%v]", DevName, err.Error())
			sdk.LoggingClient.Error(errMsg)

			return nil, err
		}
		defer res.Body.Close()

		var devReslist map[string]*nrprivate.DeviceValue

		err = json.NewDecoder(res.Body).Decode(&devReslist)
		if err != nil {
			sdk.LoggingClient.Error(fmt.Sprintf("Failed to get response for Device readings, error[%s]", err.Error()))
			//http.Error(w, err.Error(), http.StatusBadRequest)
			return nil, err
		}

	//LoggingClient.Debug(fmt.Sprintf("get Devices count[%v]", len(devicelist)))
	///////////////UpdateIndexMaps()

	return devReslist, nil
}



/// =========== 取数据过滤集列表 ===========
func (sdk *NSPlusSdk)GetDatasetByName(datasetName string) (*define.DataSetConfig, error){
	//fmt.Fprintf(os.Stdout, "Configuration.Clients[ Metadata ].Url()  %v ", Configuration.Clients[nari.ClientMetadata].Url() + nari.APIMETADATADevice)
	httpPath := fmt.Sprintf("http://%v:%v/api/v1/nsapp/dataset/name/%v", sdk.UserConfig.MQTTServer.Host, PORT_ConfigMicSvr, datasetName)

	res, err := sdk.HttpClient.Get(httpPath)
	if err != nil {
		sdk.LoggingClient.Error(fmt.Sprintf("http.Get filter error =[%v]", err.Error()))
		return nil, err
	}
	defer res.Body.Close()

	var dataset define.DataSetConfig

	err = json.NewDecoder(res.Body).Decode(&dataset)
	if err != nil {
		sdk.LoggingClient.Error(fmt.Sprintf("Failed to get response for Device, error[%s]", err.Error()))
		//http.Error(w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	sdk.LoggingClient.Debug(fmt.Sprintf("get dataset[%v]", dataset))
	///////////////UpdateIndexMaps()

	return &dataset, nil
}

func (sdk *NSPlusSdk)AddNilDatasetByName(AppName string, datasetName string, desc string)error {
	///   数据集名称改为 appname.inputname
	var req define.DataSetConfig
	req.Name = fmt.Sprintf("%v.%v", AppName, datasetName)
	req.Description = desc

	sdk.LoggingClient.Error(fmt.Sprintf("AddNilDatasetByName req[%v] ", req))
	//reqOut.CheckID = "xxxxxxxxx"  /// 此处需要增加 对 SN 的加密处理  encodeing(Algid, deviceKey, reqIn.DevSN)

	var b []byte
	b, err := json.Marshal(req)
	if err != nil {
		sdk.LoggingClient.Error(fmt.Sprintf("Failed to Marshal request parameters to json, Error[%v]", err.Error()))
		return err
	}

	bodytoIOT := bytes.NewBuffer([]byte(b))
	//调试时使用
	//res, err := http.Post("http://192.168.43.129:50000/api/v1/iot/devices/auth", "application/json;charset=utf-8", bodytoIOT)
	//res, err := http.Post("http://192.168.43.129:50000/iot/devices/auth/direct", "application/json;charset=utf-8", bodytoIOT)
	HttpPath := fmt.Sprintf("http://%v:%v/api/v1/nsapp/dataset/add", sdk.UserConfig.MQTTServer.Host, PORT_ConfigMicSvr)

	sdk.LoggingClient.Info(fmt.Sprintf("Get reg http path[%v] for dataset ", HttpPath))
	///res, err := http.Post(HttpPath, "application/json;charset=utf-8", bodytoIOT)
	res, err := sdk.HttpClient.Post(HttpPath, "application/json", bodytoIOT)
	if err != nil {
		sdk.LoggingClient.Error(fmt.Sprintf("Failed to Post reg request to IOT, Error[%v]", err.Error()))
		return err
	}

	if res.StatusCode != http.StatusOK {
		sdk.LoggingClient.Error(fmt.Sprintf("add dataset失败，返回错误码[%v]", res.StatusCode))
		return err
	}

	sdk.LoggingClient.Info(fmt.Sprintf("add dataset succuss，返回码[%v]", res.StatusCode))
	result, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		sdk.LoggingClient.Error(fmt.Sprintf("Failed to read reponse msg, Error[%v]", err.Error()))
		return err
	}
	//ack结构已知，解析到结构体


	var ack define.HttpRetCode
	err = json.Unmarshal([]byte(result), &ack)
	if err != nil {
		sdk.LoggingClient.Error(fmt.Sprintf("AddNilDatasetByName : Failed to Unmarshal http parameters, Error[%v]", err.Error()))
		return err
	}

	if(ack.Code != 0){
		sdk.LoggingClient.Error(fmt.Sprintf("Get dataset Ack success for req=[%v], ack=[%v]", req, ack))
		return errors.New(ack.ErrMsg)
	}
	sdk.LoggingClient.Info(fmt.Sprintf("AddNilDatasetByName for req[%v], ack=[%v]", req, ack))

	return nil
}


func (sdk *NSPlusSdk)SyncProfileToNsFilter(req *NsPlusProfile)error {
	/// second 边缘网关的接入， 然后在处理 device的接入， 接入后才能处理数据

	sdk.LoggingClient.Info(fmt.Sprintf("SyncProfileToNsFilter req[%v] ", req))
	//reqOut.CheckID = "xxxxxxxxx"  /// 此处需要增加 对 SN 的加密处理  encodeing(Algid, deviceKey, reqIn.DevSN)

	var b []byte
	b, err := json.Marshal(req)
	if err != nil {
		sdk.LoggingClient.Error(fmt.Sprintf("Failed to Marshal request parameters to json, Error[%v]", err.Error()))
		return err
	}

	bodytoIOT := bytes.NewBuffer([]byte(b))
	//调试时使用
	//res, err := http.Post("http://192.168.43.129:50000/api/v1/iot/devices/auth", "application/json;charset=utf-8", bodytoIOT)
	//res, err := http.Post("http://192.168.43.129:50000/iot/devices/auth/direct", "application/json;charset=utf-8", bodytoIOT)
	HttpPath := fmt.Sprintf("http://%v:%v/api/v1/df/cmd/sync/profile/app/%v", sdk.UserConfig.MQTTServer.Host, PORT_FILTERMicSvr, sdk.ContainerName)
	///Configuration.Clients[nari.ClientNsFilter].Url() + "/api/v1/df/cmd/sync/profile"

	sdk.LoggingClient.Info(fmt.Sprintf("Get reg http path[%v] for profile ", HttpPath))
	///res, err := http.Post(HttpPath, "application/json;charset=utf-8", bodytoIOT)
	res, err := sdk.HttpClient.Post(HttpPath, "application/json", bodytoIOT)
	if err != nil {
		sdk.LoggingClient.Error(fmt.Sprintf("Failed to Post reg request to IOT, Error[%v]", err.Error()))
		return err
	}

	if res.StatusCode != http.StatusOK {
		sdk.LoggingClient.Error(fmt.Sprintf("sync/profile失败，返回错误码[%v]", res.StatusCode))
		return err
	}

	sdk.LoggingClient.Info(fmt.Sprintf("sync/profile succuss，返回码[%v]", res.StatusCode))
	result, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		sdk.LoggingClient.Error(fmt.Sprintf("Failed to read reponse msg, Error[%v]", err.Error()))
		return err
	}
	//ack结构已知，解析到结构体


	var ack HttpRetCode
	err = json.Unmarshal([]byte(result), &ack)
	if err != nil {
		sdk.LoggingClient.Error(fmt.Sprintf("SyncProfileToNsFilter : Failed to Unmarshal http parameters, Error[%v]", err.Error()))
		return err
	}

	if(ack.Code != 0){
		sdk.LoggingClient.Error(fmt.Sprintf("Get dataset Ack success for req=[%v], ack=[%v]", req, ack))
		return errors.New(ack.ErrMsg)
	}
	sdk.LoggingClient.Info(fmt.Sprintf("SyncProfileToNsFilter for req[%v], ack=[%v]", req, ack))

	return nil
}




