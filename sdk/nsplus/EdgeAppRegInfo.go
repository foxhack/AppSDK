// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018 Canonical Ltd
// Copyright (C) 2018 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package nsplus

import (
	"fmt"
	"github.com/edgexfoundry/nsplussdk/sdk/nrprivate"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

//// AppMgt V2 方案：
/// 每个App基础包 对应一个app注册信息文件，该文件与app基础包文件同名，其中注明app基本信息，安装或升级时直接用来智能判断处理流程
/// 0、App基本包通过appmgt直接管理，作为版本，支持web查询并操作安装、升级等，可查询已有App基础包、名称和版本，中文名称，描述等信息等
/// 1、新增App基本包，则通过原appload接口上传，gz包，包含app的安装包和app基本信息注册文件
/// 2、App运行状态查询:app启停状态，app占用内存状态
/// 3、App启停的单独操作接口，以原容器启停接口为基础实现


/// 前置条件：
/// App包准备工作，每个App包里需要一个注册文件，制作app的时候写入相关信息，同时压缩到app包里， 操作的时候读入，并根据其中内容操作app

type EdgeBaseAppRegInfo struct {
	/// 基础App版本和名称，可以确定到唯一镜像， 应增加接口查询当前支持的基础镜像版本和名称
	AppBase          AppBaseInfo         `json:"appbase"`
	BaseDescription  AppDescriptionInfo  `json:"basedescription"`
	BaseImageID      string              `json:"baseimageid"`    /// imagename:version
	FileName         string              `json:"filename"`       /// 文件名称
	Size             string              `json:"size"`           /// 文件名称 "size": "2.6M"
	DateTime         string              `json:"datetime"`       ////"datetime": "2019-08-30 15:20:10.123"
	CtrlFlags        CtrlFlagDefine      `json:"ctrlflag"`       /// app的固有控制属性
	Config           AppConfigdefine     `json:"config"`       /// app运行的配置参数，内存、cpu、挂载目录等
}


type AppConfigdefine struct {
	ProfileName string                    `json:"profileName"`
	Cpu         IOTAppCpuDefine           `json:"cpu"`
	Mem         IOTAppMemDskDefine        `json:"mem"`
	Port        int                       `json:"port"`  // 改到profile中，只配置业务需要的挂载路径
	Mount       []string                  `json:"mount"`  // 改到profile中，只配置业务需要的挂载路径
}

/*********************************
A.7 cpu 阈值信息字段:cpuLmt
参数cpuLmt给出了与cpu资源配置相关的参数，如下表A.7所示：
表 A.7 cpu 字段说明
字段 类型 是否必选 描述
cpus number 是 CPU 核数(例如值为 2，3，4)
Lmt number 是 CPU 监控阈值
*********************************/

type IOTAppCpuDefine struct {
	Cpus               int64                           `json:"cpus"`
	CpuLmt             int64                           `json:"cpuLmt"`
}

/*********************************
A.8 内存阈值信息字段:memLmt
参数memLmt给出了与memory资源配置相关的参数，如下表A.8所示：
表 A.8 mem 字段说明
字段 类型 是否必选 描述
unit number 是 内存限值,单位：M byte
Lmt number 是 内存监控阈值，百分数
*********************************/

/*********************************
A.9 硬盘阈值信息字段:diskLmt
参数diskLmt给出了与存储disk资源配置相关的参数，如下表A.9所示：
表 A.9 disk 字段说明
字段 类型 是否必选 描述
unit number 是 存储限值，单位：M byte
Lmt number 是 磁盘存储监控阈值，百分数
*********************************/
type IOTAppMemDskDefine struct {
	Memory             int64                           `json:"memory"`
	MemLmt             int64                           `json:"memLmt"`
}


type CtrlFlagDefine struct {
	MultiEnable           bool    `json:"multienable"`
	RunOnceEnable         bool    `json:"runonce"`
}


type AppBaseInfo struct {
	/// 基本App版本和名称，当需要安装多个同名的app的情况，如主变1、主变2，基础app为zb，AppName可设定为zb1，zb2
	BaseAppName           string    `json:"baseappname"`
	BaseAppVersion        string    `json:"baseappversion"`
	/// App 开发的sdk名称和版本
	AppSdk                string    `json:"sdk"`      //// App SDK name and version, device-sdk-go,  edgex-sdk-c
	/// 10：App安装包， 20：App升级包, 30：基础空镜像包（升级）, 40: 封装App的镜像包，视作APP，不需基础镜像和基础App
	Type                  string    `json:"type"`
	/// App的描述，中文名称等应在界面显示，应在上架的时候就存储在web后台，安装app后查询显示
}

/// App的描述，中文名称等应在界面显示，应在上架的时候就存储在web后台，安装app后查询显示
type AppDescriptionInfo struct {
	/// App中文名称
	NameCHS               string              `json:"namechs"`
	/// App功能描述
	Description           string              `json:"description"`
	/// App开发商名称
	ManufactureName       string              `json:"manufactureName"`
	/// 备注，需要注重说明的内容，在Description外单独列出
	Note                  string              `json:"note"`
}


func (sdk *NSPlusSdk)LoadAppRegFiles() error  {
	///adding  add dockerapp directly; updating, delete the old and then add the new

	/// ==== 应从Config Svr上查询App对应的profile文件， HTTP接口====
	path := "./res/"

	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Printf("LoadAppRegFiles: couldn't create absolute path for: %s; %v", path, err)
		return err
	}

	///LoggingClient.Error(fmt.Sprintf("profiles: created absolute path for loading pre-defined Device Profiles: %s", absPath))
	fileInfo, err := ioutil.ReadDir(absPath)
	if err != nil {
		log.Printf("LoadAppRegFiles: couldn't read directory: %s; %v", absPath, err)
		return err
	}

	var Regfile EdgeBaseAppRegInfo

	for _, file := range fileInfo {
		fName := file.Name()
		lfName := strings.ToLower(fName)
		fullPath := absPath + "/" + fName
		log.Printf("LoadAppRegFiles: filename [%v]", fullPath)

		////if strings.HasSuffix(lfName, yamlExt) || strings.HasSuffix(lfName, ymlExt) {
		if strings.HasSuffix(lfName, RegfileExt) {
			//var profile NsPlusProfile
			err := nrprivate.LoadJsonFromFile(absPath, fName, &Regfile) ///StationReg
			if err != nil {
				errMsg := fmt.Sprintf("LoadAppRegFiles: 加载nsplus App [%v] reginfo fail， error[%v]", fName, err.Error())
				log.Printf("LoadAppRegFiles: %v ", errMsg)
				return err
			}
		}
	}

	sdk.Regfile = Regfile
	log.Printf("LoadAppRegFiles:  [%v]", sdk.Regfile)

	return nil
}


