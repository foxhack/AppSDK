// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018 Canonical Ltd
// Copyright (C) 2018 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0
package nrprivate

///==== 网关上线接口定义 ====
type EdgeRegReqDefine struct{
	EdgeSn            string        `json:"edgeSn"`
	EdgeAddr          string        `json:"edgeAddr"`
	EdgePort          int           `json:"edgePort"`
}

type EdgeRegAckDefine struct{
	Code              int              `json:"code"`
	ErrMsg            string           `json:"errMsg"`
	EdgeId            int              `json:"edgeId"`
}

///==== 网关配置查询接口定义 ====
/// 设备Id， 为0是表示召唤网关配置信息，-1标识召唤网关和所有设备配置，其有效值则查询对应设备的配置
type EdgeConfigReqDefine struct{
	DevId             int              `json:"devId"`
}

// type is -1， 返回所有配置
type AllConfigAckDefine struct{
	EdgeConfig        EdgeConfigAckDefine    `json:"edgeConfig"`
	DevConfig         []DevConfigAckDefine   `json:"devConfig"`  ///
}

/// type = 0
type EdgeConfigAckDefine struct{
	StationName       string           `json:"stationName"`
	VoltageLevel      string           `json:"voltageLevel"`  /// 220,表示 220kV
	PosLongitude      float64          `json:"posLongitude"`
	PosLatitude       float64          `json:"posLatitude"`
	EdgeSn            string           `json:"edgeSn"`
	EdgeModel         string           `json:"edgeModel"`
	EdgeManufacture   string           `json:"edgeManufacture"`
	EdgeLocation      Location         `json:"edgeLocation"`

	EdgeAddr          string           `json:"edgeAddr"`
	EdgePort          int              `json:"edgePort"`
	EdgeWebPort       string            `json:"edgeWebPort"`
}
// type is devId
type DevConfigAckDefine struct{
	DevSn             string           `json:"devSn"`
	DevType           string           `json:"devType"`  ///
	DevModel          string           `json:"devModel"`
	DevManufacture    string           `json:"devManufacture"`

	Capacity          []string         `json:"capacity"`    /// 能力列表，待定义
	AppBindDevId      int              `json:"appBindDevId"`   /// App管理绑定的设备Id
	DevLocation       Location       `json:"devLocation"`

	DevName           string           `json:"devName"`  ///
	//EdgeAddr          string           `json:"edgeAddr"`
	//EdgePort          int              `json:"edgePort"`
}

type Location struct{
	Room              string           `json:"room"`     /// 安装小室
	Cabinet           string           `json:"cabinet"`  /// 屏柜
	Rock              string           `json:"rock"`     /// 槽位
}

///==== 二次设备和App上线接口定义 ====
type DevsRegReqDefine struct{
	Devices           []DeviceRegDefine  `json:"devices"`
	EdgeId            int              `json:"edgeId"` ///上报devlist
	/// ----
	Cer               string           `json:"cer"`  /// 证书请求文件（后期迭代版本考虑增加证书加密接入）
}
type DeviceRegDefine struct{
	DevSn             string           `json:"devSn"`
	DevType           string           `json:"devType"`  /// 220,表示 220kV
	DevModel          string           `json:"devModel"`
	DevManufacture    string           `json:"devManufacture"`

	DevName           string           `json:"devName"`    /// 能力列表，待定义
}
/// 数组形式返回
type DevsRegAckDefine struct{
	Sn                string           `json:"sn"`
	DevId             int              `json:"devId"` //
	Code              int              `json:"code"`  /// 返回码:0代表成功,1代表失败
	Msg               string           `json:"msg"`  ///  如果失败，返回结果描述，最长不超过256字符

}


///==== 网关、二次设备和App下线命令接口定义 ====
/// 设备Id， 为0是表示网关下线（就是所以设备下线重启上线流程），其他有效devId值则下线对应设备
type EdgeDevRebootCmdDefine struct{
	DevId             int              `json:"devId"`
}
/// 网关返回 StatusOK

///==== 二次设备和App在线状态上报接口定义 ====
type OnlineStausReportDefine struct{
	DevId             int              `json:"devId"`
	IsOnline          int              `json:"isOnline"`
}
/// 平台返回 StatusOK=================IOT 通用

///==== 网关在线状态查询接口定义 ====
type EdgeOnlineStausQueryDefine struct{
	HeartBeat         string           `json:"heartBeat"`
}
/// 平台返回 StatusOK
type EdgeOnlineStausAckDefine struct{
	EdgeId            string           `json:"edgeId"`
	Time              string           `json:"time"`
}



