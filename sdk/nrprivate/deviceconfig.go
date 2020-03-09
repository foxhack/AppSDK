package nrprivate

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type  DeviceConfig struct {
	Id                string  	 			 	`json:"_id,omitempty"`
	DeviceName		  string         	 		`json:"devicename,omitempty"`//	[1]	设备名称
	ServiceName       string                  	`json:"servicename,omitempty"`
	StationAlias      map[string]string          `json:"staionalias,omitempty"`
	LimitSetting      map[string]ResConfig       `json:"limitsetting,omitempty"`
}

type  ResConfig struct {
	LowLimit      float64       `json:"lowlimit"`
	HighLimit     float64       `json:"highlimit"`
	RecordPeriod  int32         `json:"recordperod"`
	MaxRecordTime int64         `json:"maxrecordtime"`   //最大存储时间 0 默认全周期存储， -1 不存储

}


// MarshalJSON implements the Marshaler interface in order to make empty strings null
func (deviceonfig  DeviceConfig) MarshalJSON() ([]byte, error) {
	test := struct {
		Id                *string  	 			 	`json:"_id,omitempty"`
		DeviceName		  *string         	 		`json:"devicename,omitempty"`//	[1]	设备名称
		ServiceName       *string                  	`json:"servicename,omitempty"`
		StationAlias      map[string]string          `json:"staionalias,omitempty"`
		LimitSetting      map[string]ResConfig       `json:"limitsetting,omitempty"`
	}{
	}
	// Empty strings are null
	if deviceonfig.Id != "" {
		test.Id=&deviceonfig.Id
	}
	if deviceonfig.DeviceName !="" {
		test.DeviceName=&deviceonfig.DeviceName
	}
	if deviceonfig.ServiceName !="" {
		test.ServiceName=&deviceonfig.ServiceName
	}

	if deviceonfig.StationAlias !=nil {
		test.StationAlias=deviceonfig.StationAlias
	}
	if deviceonfig.LimitSetting !=nil {
		test.LimitSetting=deviceonfig.LimitSetting

	}



	return json.Marshal(test)
}

/*
 * To String function for DeviceResource
 */
func (basic DeviceConfig) String() string {
	out, err := json.Marshal(basic)
	if err != nil {
		return err.Error()
	}
	return string(out)
}



// MarshalJSON implements the Marshaler interface in order to make empty strings null
func (rescfg  ResConfig) MarshalJSON() ([]byte, error) {
test:= struct {
	LowLimit      float64       `json:"lowlimit"`
	HighLimit     float64       `json:"highlimit"`
	RecordPeriod  int32         `json:"recordperod"`
	MaxRecordTime int64         `json:"maxrecordtime"`   //最
}{
	rescfg.LowLimit,
	rescfg.HighLimit,
	rescfg.RecordPeriod,
	rescfg.MaxRecordTime,
}	// Empty strings are null


  fmt.Sprint("$%s  v",reflect.TypeOf(rescfg),test)
	return json.Marshal(test)
}


/*
 * To String function for Resconfig
 */
func (rescfg  ResConfig) String() string {
	out, err := json.Marshal(rescfg)
	if err != nil {
		return err.Error()
	}
	return string(out)
}
