// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2017-2018 Canonical Ltd
// Copyright (C) 2018 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package nsplus

import (
	"errors"
	"fmt"
	"github.com/edgexfoundry/nsplussdk/sdk/nrprivate"
	"log"
)

const(
	NULL_UUID = "00000000000000000000"
)
//var profileMap []nari.AppProfile
/// 返回 uuid
func (sdk *NSPlusSdk)LoadProfiles(profileName string) error  {
	///adding  add dockerapp directly; updating, delete the old and then add the new

	/// ==== 应从Config Svr上查询App对应的profile文件， HTTP接口====


	//var profile NsPlusProfile
	err := nrprivate.LoadJsonFromFile("./res/", profileName, &sdk.Profile) ///StationReg
	if err != nil {
		errMsg := fmt.Sprintf("LoadProfiles: 加载nsplus profile=[%v] fail， error[%v]", profileName, err.Error())
		sdk.LoggingClient.Error(errMsg)
		return err
	}

	return nil
}


func LoadAppProfile(path string, appName string) (*NsPlusProfile, error)  {
	///adding  add dockerapp directly; updating, delete the old and then add the new
	log.Printf("LoadAppProfile with path[%v] appname[%]", path, appName)
	if path == "" {
		return nil, errors.New("path cannot be null!")
	}
	if appName == "" {
		return nil, errors.New("appname cannot be null!")
	}

	var nsprofile NsPlusProfile
	err := nrprivate.LoadJsonFromFile(path + "/res/", appName , &nsprofile) ///+ ".profile"
	if err != nil {
		errMsg := fmt.Sprintf("LoadProfiles: 加载nsplus App [%v] profile失败， error[%v]", appName, err.Error())
		log.Printf(errMsg)
		return nil, err
	}

	return &nsprofile, nil
}


