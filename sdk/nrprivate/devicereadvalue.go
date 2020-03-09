package nrprivate

import (
	"encoding/json"
	"sync"
)

type ReadValue struct {  //把读写锁和资源map封装在一起
	sync.RWMutex
	value map[string]map[string]DeviceValue
}

type AlmMapValue struct {  //把读写锁和资源map封装在一起
	sync.RWMutex
	Alarm map[string]map[string]NrAlarm
}



type ResValue struct {
	sync.RWMutex
	Device      string
	DeviceName  string
	Updattime   int64  //最新一次更新時間
	Value       map[string]*DeviceValue
}

type PreDeviceMap struct {  //把读写锁和资源map封装在一起
	sync.RWMutex
	Device map[string]DeviceAutoRegInfo
}
type DeviceValue struct {
	//Id          string    `json:"id",bson:"id" `
	Created     int64     `json:"created" bson:"created"` // When the reading was created
	Device      string    `json:"device" bson:"device""`
	Devicename  string    `json:"devicename" bson:"devicename"`
	Name        string    `json:"name" bson:"name" `
	Value       string    `json:"value"  bson:"value"`            // Device sensor data value
	Description string    `json:"description"  bson:"description"`            // Device sensor data value
	Unit 		string    `json:"unit"  bson:"unit" `
	Measuretype string    `json:"measuretype"  bson:"measuretype" `
	DataType    string     `json:"-"  bson:"datatype"`
	RecordTime   int64     `json:"-"  bson:"recordtime"`
}

type NrDevPloyLoad struct {
	Id 			string  `json:"id" codec:"id,omitempty"`
	Created     int64     `json:"created" codec:"created,omitempty"` // When the reading was created
	Device      string    `json:"device" codec:"device,omitempty"`
	Devicename  string    `json:"devicename" codec:"devicename,omitempty"`
	NrPloyLoad   []NrValue `json:"nrployload" codec:"nrployload,omitempty"`
}

type  NrValue struct {
	Name        string    `json:"name" codec:"name,omitempty"`
	Description string    `json:"description"  codec:"description,omitempty"`            // Device sensor data value
	Value       string    `json:"value"  codec:"value,omitempty"`           			 // Device sensor data value
	Unit 		string    `json:"unit"  codec:"unit,omitempty"`
	Measuretype string    `json:"measuretype"  codec:"unit,omitempty"`
	DataType    string    `json:"datatype" codec:"datatype,omitmpty"`
}

func NewNrDevPloyLoad() *NrDevPloyLoad {
	ployload:=NrDevPloyLoad{}
	ployload.NrPloyLoad=make([]NrValue,0)
	return &ployload
}
func (devployload *NrDevPloyLoad)Transfor(value *DeviceValue,Id string ){

	if devployload.Id!="" {
		devployload.Id=Id
		devployload.Devicename=value.Devicename
		devployload.Device=value.Device
	}

	ployload:=NrValue{}
	ployload.DataType=value.DataType
	ployload.Measuretype=value.Measuretype
	ployload.Unit=value.Unit
	ployload.Value=value.Value
	ployload.Name=value.Name
	ployload.Description=value.Description

	devployload.NrPloyLoad=append(devployload.NrPloyLoad,ployload)



}

// String provides a JSON representation of the Event as a string
func (devployload NrDevPloyLoad) String() string {
	out, err := json.Marshal(devployload)
	if err != nil {
		return err.Error()
	}

	return string(out)
}