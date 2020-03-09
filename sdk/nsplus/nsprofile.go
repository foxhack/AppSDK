/*******************************************************************************
 * Copyright 2019 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package nsplus

import (
	"encoding/json"
	//"github.com/edgexfoundry/edgex-backend/internal/nari"
	"log"
	"math/rand"
	//"encoding/json"
	"fmt"
	"strconv"
	"time"

)


type NsPlusProfile struct {
	///DescribedObject `yaml:",inline"`
	Encoding          string            `json:"encoding"`
	Uuid              string            `json:"uuid"`               // Non-database identifier (must be unique)
	TemplateVersion   string            `json:"templateVersion"`    /// V1.0.0
	Type              string            `json:"type"`
	Model             string            `json:"model"`              // Model of the device
	ManufactureId     string            `json:"manufactureId"`      // Manufacturer id of the device
	ManufactureName   string            `json:"manufactureName"`    // Manufacturer name of the device
	Channel           []ChannelConfig   `json:"channel"`       ////  暂定不通过用户接口配置channel， 程序内部使用sdk的接口实现数据交互
	Inputs            []InputsInfo      `json:"inputs"`
	Outputs           []OutputsInfo     `json:"outputs"`
	Parameter         []ParameterInfo   `json:"parameter"`          // List of params for app associated with this profile

	Description       string            `json:"description"`        /// App功能描述 20191202

	Config            ServerConfig      `json:"config"`       ////  设置业务需要的port和端口


	EnableFlag        string            `json:"enable"`            // 使能开关
}

type ServerConfig struct {
	Port              int               `json:"port"`
	Mount             []string          `json:"mount"`
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



type FileInterface struct {
	Name              string            `json:"name"`
	Size              int               `json:"size"`
	Time              string            `json:"time"`
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
	//DevType       string           `json:"devType"`
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



func  GetRandomString(leng int, stringRange string) string {
	//str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(stringRange)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < leng; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	//LoggingClient.Info(fmt.Sprintf("GetRandomString:[%v]leng[%v] => [%v] ",leng,stringRange, string(result)))
	return string(result)
}


func TimeStringLen16ForNow(timeIn time.Time) string {
	//now := time.Now()
	dotS := timeIn.Nanosecond() / 1000000
	return timeIn.Format("200601021504") + fmt.Sprintf("%04v", dotS)
}


func FormatStringForNow(timeIn time.Time) string {
	//now := time.Now()
	dotS := timeIn.Nanosecond() / 1000000
	return timeIn.Format(Timer_Format_UCT) + "." + fmt.Sprintf("%03v", dotS)
}


func (sdk *NSPlusSdk)InitNameValue(Type string, reading string) interface{} {
	var Value interface{}
	if(reading == "") {
		switch Type {
		case AppType_Int:
			Value = 0
		case AppType_Double:
			Value = 0.001
		case AppType_String:
			Value = ""
		case AppType_Bool:
			Value = false
		case AppType_Enum:
			Value = ""
		case AppType_Time:
			///now := time.Now()
			Value = FormatStringForNow(time.Now())
			log.Printf("initNameValue: init get Time Value %v \n", Value)
		case AppType_Array:
			Value = nil
		case AppType_File:
			var filePara FileInterface
			filePara.Name = ""
			filePara.Size = 0
			filePara.Time = ""
			Value = filePara
		case AppType_Event:
			var event EventAlarmDefine
			event.Type = "ALARM"
			//event.Level = 0
			event.Info = ""
			event.Time = FormatStringForNow(time.Now())
			Value = event
		case AppType_Stream:
			Value = ""
		default:
			log.Printf("Type not supported: %s", Type)
			Value = ""
		}
	}else{
		switch Type {
		case AppType_Int:
			valueint, _ := strconv.Atoi(reading)
			Value = valueint
			//inputsData.Inputs = append(inputsData.Inputs, names) ///reading.Value
		case AppType_Double:
			valueint, _ := strconv.ParseFloat(reading, 32)
			Value = valueint
			//inputsData.Inputs[index].Value = RandInt(1, 80)/10
			//log.Printf("processEvent: get double %v \n", Value)
		case AppType_String:
			Value = reading
			//names.Value = reading.Value
		case AppType_Bool:
			boolValue, _ := strconv.ParseBool(reading)
			Value = boolValue
		case AppType_Enum:
			Value = reading
		case AppType_Time:
			timestring := fmt.Sprintf("%v", time.Now())
			stamp, _ := time.ParseInLocation(Time_Format_dot001, timestring, time.Local) //使用parseInLocation将字符串格式化返回本地时区时间
			log.Printf("Get read for time: read[%v] timestring[%v], stamp[%v]", reading, timestring, stamp)
			Value = FormatStringForNow(stamp)
		case AppType_Array: ///// ======================= 具体格式待确定 需要实例调试 =========================
			Value = reading
		case AppType_Event:
			var event EventAlarmDefine
			var data []byte = []byte(reading)
			err := json.Unmarshal([]byte(data), &event)
			if err != nil {
				sdk.LoggingClient.Info(fmt.Sprintf("receive file resource but tranform to json err!! [%v] \n ", err.Error()))
				//continue
			}
			//LoggingClient.Error(fmt.Sprintf("files: json:[%v] \n ", filePara))
			Value = event
			log.Printf("get files: reading [%v]: to [%v] \n ",reading, Value)
		case AppType_Stream: ///// ======================= 具体格式待确定 需要实例调试 =========================
			Value = reading
		default:

			Value = reading
			sdk.LoggingClient.Info(fmt.Sprintf("Type transparent for[%s], reading[%v]", Type, reading))
			//return false
		}
	}
	return Value
}

