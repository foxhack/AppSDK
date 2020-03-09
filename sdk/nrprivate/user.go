package nrprivate

type  NrUserAdmin struct {
	  Id       string       `json:"-"`
	  User     string       `json:"user"`
	  Key 	   string       `json:"key"`
	  Created  int64        `json:"created"`
	  Updated  int64        `json:"updated"`
	  Period   int64        `json:"period,omitempty"` //存储周期
}