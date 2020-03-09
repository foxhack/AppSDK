package nrprivate

import (
    "encoding/json"
)

//历史库。 遥信遥测 必须有默认值，默认值是正常。不正常就复位了
// ，存储目前存在的告警，同时要存存储数据库。主要是为了防止历史库 是否会回复。
// 记录告警发生时间，恢复时间 。这两个应该都在同一个 存储在一个具体的 记录里。产生时间。恢复时间。
// 1) 正常恢复
// 2) 进程重启恢复，去除
// 3)手动恢复，去除
// 4)门限修改检查
// 4) 遥测类型。每个告警都记录，恢复只
// 实在不行，只记录先前的告警。
type NrAlarm struct {
    Id           string   	     `json:"id,omitempty"`
    AlarmId      string           `json:"alarmid,omitempty"` //
    Device       string       	 `json:"device,omitempty"`
    Resource     string       	 `json:"resource,omitempty"`
    ResType      string 			 `json:"restype,omitempty"`//YX YC
    AlarmType    string           `json:"alarmtype,omitempty"`//1 Cteate Rescover, Reset告警，2 恢复 ,3 遥信复位
    Oldvalue     string           `json:"oldvalue,omitempty"`// 旧值
    NewValue     string           `json:"newvalue,omitempty"`// 新值
    CreateTime   int64            `json:"createtime,omitempty"`// 事件产生时间
    UpdateTime   int64            `json:"updatetime,omitempty"`// 遥信存储变位的周期时间
    OldCreatetime int64   	     `json:"oldcreatetime,omitempty"`//遥测告警产生时间 ，告警为0，复归，复位为告警时间
    Lowlimit      float64        `json:"lowlimit"`//下线
    HighLimit     float64         `json:"highlimit"`//上线
    YcType        string            `json:"YcType,omitempty"`//是否越上限，越下限
}
/*
// MarshalJSON implements the Marshaler interface in order to make empty strings null
func (nralarm  NrAlarm) MarshalJSON() ([]byte, error) {


    return json.Marshal(&nralarm)
}
*/
/*
 * To String function for DeviceResource
 */
func (nralarm  NrAlarm) String() string {
    out, err := json.Marshal(nralarm)
    if err != nil {
        return err.Error()
    }
    return string(out)
}
