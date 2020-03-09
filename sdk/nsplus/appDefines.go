//
// Copyright (c) 2018 Tencent
// Copyright (c) 2019 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//
package nsplus

import (
	//"github.com/edgexfoundry/edgex-backend/internal/nari"
	"golang.org/x/text/encoding/simplifiedchinese"
	"math/rand"
	"strconv"
	"time"
)

const(


	//profileExt = ".profile"
	//ProfilePath        = "/home/NEdge/res"
	//ProfileBakPath        = "/home/NEdge/res/timeout/"
	OutputWriteOutputCmd     = "writeOutputs"
	OutputReadInputCmd     = "writeOutputs"

	RequestReadOutputCmd     = "readOutputs"
	RequestReadParamCmd      = "readParameter"
	RequestWriteParamCmd     = "writeParameter"
	ReloadProfileCmd         = "reloadProfile"

	AppTimeOutTick         int    = 10

	INPUTTopic string = "Input"
	OUTPUTTopic string = "Output"
	REQUESTTopic string = "Request"
	RESPONSETopic string = "Response"

	//ApiAppOnOff                 = "/api/v1/apponoff"
	//ApiAppSyn                   = "/api/v1/appsyn"
	ApiRequestCommands          = "/api/v1/request"
	ApiInputCommands            = "/api/v1/input"

	AppInfo                     = "app"
	Uuid                        = "uuid"

	AppRunStatusReady           = "App启动中"
	AppRunStatusRunning         = "运行中"
	AppRunStatusLostHeart       = "心跳丢失"
	AppRunStatusStopped         = "停止"
	AppRunStatusNoProfile       = "未配置Profile"

	AppRunStatusNotCreated      = "容器未创建"
	AppRunStatusCreated         = "容器已创建"
	AppRunStatusUnknown         = "未知异常(GetStatusFail)"


	CODEUTF8    = "UTF-8"
	CODGBK      = "GBK"

)

////var gInputsMessageId int
type inputsInterface struct {
	Sqid       int           `json:"sqid"`
	Timedate   string        `json:"time"`
	Inputs     []NameValues  `json:"inputs"`
}
type NameValues struct {
	Name       string        `json:"name"`
	Value      interface{}   `json:"value"`
}

// for http commands
type requestWriteHttp struct {
	AppId      string        `json:"appid"`
	Command    string        `json:"command"`
	Arguments  []NameValues      `json:"arguments"`
}

// for commands send by mqtt
type requestWriteInterface struct {
	Sqid       int           `json:"sqid"`
	Timedate   string        `json:"time"`
	Command    string        `json:"command"`
	Arguments  []NameValues      `json:"arguments"`
}

//// for read commands
type requestInterface struct {
	Sqid       int           `json:"sqid"`
	Timedate   string        `json:"time"`
	Command    string        `json:"command"`
	Arguments  []argInterface      `json:"arguments"`
}
type argInterface struct {
	Name       string        `json:"name"`
	//Value      string        `json:"value"`
}

//// for reloadprofile commands
type requestReloadProfile struct {
	Sqid       int           `json:"sqid"`
	Timedate   string        `json:"time"`
	Command    string        `json:"command"`
	//Arguments  []argInterface      `json:"arguments"`
}

//// beatheart
type timeOutInterface struct {
	TimeoutTick   int
	CurSqid       int       //      `json:"sqid"`
	LastSqid      int       //   `json:"time"`
	LastStatus    string    //      `json:"status"`
	Timedate      string        `json:"time"`
	///Results    interface{}     `json:"results"`
}

////var  int
type ResponseInterface struct {
	Sqid       int             `json:"sqid"`
	Timedate   string          `json:"time"`
	Status     string          `json:"status"`
	///Command    string        `json:"command"`
	Results    []NameValues    `json:"results"`
}

////var  int
type OutputInterface struct {
	Sqid       int           `json:"sqid"`
	Timedate   string        `json:"time"`
	Command    string        `json:"command"`
	Outputs    []NameValues  `json:"arguments"`   ////  `json:"outputs"`  ///[]nari.NameValues
}












/// profile list： file name， file size and modify time
type profileList struct {
	AppId          string         `json:"appid"`
	FileName       string         `json:"name"`
	ModTime        time.Time      `json:"time"`
	FileSize       int64          `json:"size"`
	BValid         bool           `json:"valid"`
}

type AppListtoHttp struct {
	Uuid            string `json:"uuid"`            // Non-database identifier (must be unique)
	TemplateVersion string `json:"templateVersion"` /// V1.0.0
	AppName         string `json:"appname"`   //// 暂定对应容器名称
	Type            string `json:"type"`
	Model           string `json:"model"`           // Model of the device
	ManufactureId   string `json:"manufactureId"`   // Manufacturer of the device
	ManufactureName string `json:"manufactureName"` // Manufacturer of the device

	//////////////////////ContainerInfo   oneKeyqueryAck `json:"containerinfo"`
}

type AppRunDatatoHttp struct {
	Uuid       string   `json:"uuid"` /// Appid
	Inputs     []Params `json:"inputs"` /// 当前inputs
	Outputs    []Params `json:"outputs"`/// 最新outputs
	Parameters []Params `json:"parameters"`/// Parameters
}

type Params struct {
	Name       string               `json:"name"`
	Value      interface{}          `json:"value"`
	Define     interface{}          `json:"define"`
	//Unit       string          `json:"unit"`
}



type DockerAppInfo struct {
	ContainerName       string          `json:"containername"`
	Name                string          `json:"name"`
	AppId               string          `json:"appid"`
	RunStatus           string          `json:"status"`
	bSyn                bool            `json:"bSyn"`
}


////get or post bind , gin
type AppAck struct {
	AppName  string    `json:"appname"`
	Result   string    `json:"result"`
}


//// GBK 转 utf-8
type Charset string
const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)
func ConvertByte2String(byte []byte, charset Charset) []byte {
	//var str string
	switch charset {
	case GB18030:
		var decodeBytes,_=simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		return decodeBytes
	case UTF8:
		fallthrough
	default:
		return  byte
	}
	return  byte
}

//// -------------  for test --------------
//随机整数
func RandInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
func  GetRandomFloatString(l int) float64 {
	str := "0123456789."
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	v1, err := strconv.ParseFloat(string(result), 32)

	if err != nil {
		return rand.Float64() * 10000.01
	}
	return v1
}


