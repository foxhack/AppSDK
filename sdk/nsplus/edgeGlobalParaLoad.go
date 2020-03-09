// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018 Canonical Ltd
// Copyright (C) 2018 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package nsplus

import (
	"github.com/edgexfoundry/nsplussdk/sdk/define"
	"github.com/edgexfoundry/nsplussdk/sdk/nrprivate"
)


///var UserConfig define.EdgeConfigStruct // nari.UserConfigStruct

func (sdk *NSPlusSdk)GetUserConfig() error {
	var err error


	ack := define.EdgeConfigStruct{}
	err =  nrprivate.LoadJsonFromFile(nrprivate.RootPath_Edge_global, define.EdgeConfigFILE, &ack) ///StationReg
	if err != nil {
		//LoggingClient.Error(fmt.Sprintf("加载站点注册信息失败， error[%v]", err.Error()))
		///errMsg := fmt.Sprintf("加载网关全局配置失败， error[%v]", err.Error())
		//LoggingClient.Error(errMsg)
		//w.WriteHeader(http.StatusBadRequest)
		//io.WriteString(w, errMsg)
		return err
	}

	sdk.UserConfig = ack
	/**UserConfig, err = nari.LoadUserConfig()
	if err != nil {
		return err
	}
	if err == nil {

		if (UserConfig.IOTSetting.Host == "") || (UserConfig.IOTSetting.Port == "") {
			//UserConfig.IOTSetting.Host = "127.0.0.1"
			//UserConfig.IOTSetting.Port = "48073"
			return errors.New(fmt.Sprintf("加载网关全局配置异常， IOT 连接设置无效[%v]", UserConfig.IOTSetting))
		}
		if (UserConfig.MQTTServer.Host == "") || (UserConfig.MQTTServer.Port == "") {
			//UserConfig.MQTTServer.Host = "127.0.0.1"
			//UserConfig.MQTTServer.Port = "1883"
			return errors.New(fmt.Sprintf("加载网关全局配置异常， 内部MQTT连接设置无效[%v]", UserConfig.MQTTServer))
		}
	}**/

	return nil
}



