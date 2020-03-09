// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018 Canonical Ltd
// Copyright (C) 2018 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package define

const(
	DatasetConfFILE          = "datasetconf.json"
)

type DataSetConfig struct{
	Name             string                   `json:"name"`            // 数据集的名称，唯一标识，英文+数字串，例如 ycset
	Description      string                   `json:"description"`    // 数据集的相关描述，如含有什么设备和资源，针对的需求等等
	DevRes           []DataSetByDevRes        `json:"devRes"`         // 设备+资源的列表， 数组形式, 支持多个设备，每个设备中多个资源

	Version          string                   `json:"version"`        // 数据集版本（对数据集成员修改后自动取时间戳做版本号），由程序自动生成，不需要通过接口配置

	//SelOption        DevSelectOption          `json:"selOption"`
}

type DataSetByDevRes struct{
	DevName          string                   `json:"devname"`        // Unique name for identifying a device ， use as deviceSN
	Parent           string                   `json:"parent"`         // 父节点网关SN，网关SN默认全网唯一，与devSN一起唯一标识一个dev
	Res              []string                 `json:"resource"`       /// 数据集重选中的资源列表

	//SelOption        ResSelectOption          `json:"selOption"`
}

type ResSelectOption struct{
	/// 设备资源选择条件，暂定可设置为：
	// 0-UserDefine,不使用该选项，由用户自己选择，默认值
	// 1-All, 选择全部资源，
	// 2-All-YC-Res，当前设备的全部遥测资源
	// 3-All-YX-Res, 当前设备的全部遥信数资源
	// =====
	// 10-All-Int-Res，当前设备的全部整数资源
	// 11-All-Float-Res, 当前设备的全部浮点数资源
	// 12-All-String-Res, 当前设备的全部字符串资源
	// =====
	Option           string                   `json:"option"`
}
type DevSelectOption struct{
	/// 设备的选择条件，暂定可设置为：
	// 0-UserDefine,不使用该选项，由用户自己选择，默认值, 当设备的选择条件Option值为0-None时，DataSetByDevRes中的设置有效
	// 1-All, 选择全部设备全部资源，
	// 2-All-YC-Res，所有设备的所有遥测资源
	// 3-All-YX-Res, 所有设备的所有遥信资源
	// ==== 按设备的资源类型选择 ===
	// 10-All-Int-Res，所有设备的所有整数资源
	// 11-All-Float-Res, 所有设备的所有浮点数资源
	// 12-All-String-Res, 所有设备的所有字符串资源
	// ====
	Option           string                   `json:"option"`
}

