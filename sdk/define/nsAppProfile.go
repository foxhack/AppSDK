// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2018 Canonical Ltd
// Copyright (C) 2018 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package define

const(
	NsAppProfileConfFILE          = "nsappprofile.json"
	NsAppProfileExt               = ".profile"
)


type NsPlusProfile struct {
	Description       string            `json:"description"`        /// App功能描述 20191202
	Encoding          string            `json:"encoding"`
	Uuid              string            `json:"uuid"`               // Uuid由程序分配，
	TemplateVersion   string            `json:"templateVersion"`    /// V1.0.0
	Type              string            `json:"type"`
	Model             string            `json:"model"`              // Model of the device
	ManufactureId     string            `json:"manufactureId"`      // Manufacturer id of the device
	ManufactureName   string            `json:"manufactureName"`    // Manufacturer name of the device
	Channel           []ChannelConfig   `json:"channel"`  ////  暂定不通过用户接口配置channel， 程序内部使用sdk的接口实现数据交互
	Inputs            []InputsInfo      `json:"inputs"`
	Outputs           []OutputsInfo     `json:"outputs"`
	Parameter         []ParameterInfo   `json:"parameter"`          // List of params for app associated with this profile

	EnableFlag            string              `json:"enable"`            // 使能开关
}

// MqttServer represents a server for app
type ChannelConfig struct {
	Name              string            `json:"name"`
	Protocol          string            `json:"protocol"`
	Host              string            `json:"host"`
	Port              int               `json:"port"`
	User              string            `json:"user"`
	Password          string            `json:"password"`
	Qos               int               `json:"qos"`
	KeepAlive         int               `json:"keepAlive"`
	TopicPath         string            `json:"topicPath"`
}
type InputsInfo struct {
	Name              string            `json:"name"`
	Description       string            `json:"description"`
	Type              string            `json:"type"`
	MaxResCount       int               `json:"maxResCount"`
	Coefficient       float32           `json:"coefficient"`
	Offset            float32           `json:"offset"`
	Reference         string            `json:"reference"`
}
type OutputsInfo struct {
	Name              string            `json:"name"`
	Description       string            `json:"description"`
	Type              string            `json:"type"`
	Unit              string            `json:"unit"`
	AttrType          string            `json:"attrType"`
}
type ParameterInfo struct {
	Name              string            `json:"name"`
	Description       string            `json:"description"`
	Type              string            `json:"type"`
	Unit              string            `json:"unit"`
	ReadWrite         string            `json:"readWrite"`
	Minimum           interface{}       `json:"minimum"`
	Maximum           interface{}       `json:"maximum"`
	Step              interface{}       `json:"step"`
	DefaultValue      interface{}       `json:"defaultValue"`
}


type ValueSOEDef struct{
	State             string            `json:"state"`
	Reference         string            `json:"reference"`
}
type ValueFileDef struct{
	Name              string            `json:"name"`
	Size              int64             `json:"size"`
}
type ValueAlarmDef struct{
	Code              float64           `json:"code"`
	Level             float64           `json:"level"`
}
type ValueRawDef struct{
	Text              string            `json:"text"`
}
type EventAlarmDefine struct{
	Type          string         `json:"type"`
	Info          string         `json:"info"`
	Time          string         `json:"time"`
	ValueSOE      ValueSOEDef    `json:"valueSOE"`
	ValueFILE     ValueFileDef   `json:"valueFILE"`
	ValueALARM    ValueAlarmDef  `json:"valueALARM"`
	ValueRAW      ValueAlarmDef  `json:"valueRAW"`
}

