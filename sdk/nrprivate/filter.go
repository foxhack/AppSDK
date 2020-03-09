package nrprivate
type NrFilter struct {
	Id             string          `json:"id,omitempty"`//
	ClientName     string    	   `json:"clientname,omitempty"`//导出设备名称，也可以是聚合
	FilterSubType  string           `json :"filtersubtype,omitempty"`  // 默认该字段不需要，如果是针对供服的特制接口则需要校验类型
	Filters 	   []NrFilterApp   `json:"filters"`
}

//新增数据聚合
type   NrAggregation struct {
	Id              string  	 			 	 `json:"id"`
	AggergName 	    string   				     `json:"aggergname"` // 聚合名称
	AggerDesc       string                       `json:"aggerdesc"`  //聚合描述
	State           bool                         `json:"state"`
	Filters 	    []NrFilterApp                `json:"filters"`
}
type NrAggrCache struct {
	AggergName 	    string   				     `json:"aggergname"` // 聚合名称
	AggerDesc       string                       `json:"aggerdesc"`  //聚合描述
	State           bool                         `json:"state"`
	//Filters 	    []NrFilterApp                `json:"filters"`
	DevRes      map[string]map[string]NrFilterType
}
func (nr * NrAggregation)Cache() *NrAggrCache {
  nrcache :=NrAggrCache{}
  nrcache.AggergName=nr.AggergName
  nrcache.AggerDesc=nr.AggerDesc
  nrcache.State=nr.State
  nrcache.DevRes=make(map[string]map[string]NrFilterType)
  for _,filter :=range nr.Filters {
  		if (nrcache.DevRes[filter.DeviceName]==nil){
  			nrcache.DevRes[filter.DeviceName]=make(map[string]NrFilterType)
		}
	  nrcache.DevRes[filter.DeviceName][filter.ResourceName]=filter.Filter
  }
  return  &nrcache

}


type NrFilterApp struct {
	 DeviceName    string   	 	`json:"devicename"`
	 ResourceName  string    		`json:"resourcename"`
	 Filter   NrFilterType 	       `json:"filter"`
}

type NrFilterType struct {
	Type        FilterType          `bson:"type,omitempty"`
	Threshold   float64			  `bson:"threshold,omitempty"`
}


type FilterType int32

const (
	Filter_Update   FilterType  =0   //更新，或者无计算
	Filter_Change   FilterType  =1   // 变化
	Filter_Timer    FilterType  =2   // 定时
	Filter_gt       FilterType  =3   //大于等于
	Filter_lt       FilterType  =4   //小于
	Filter_gte      FilterType  =5   //大于等于
	Filter_lte      FilterType  =6   //小于等于
	Filter_eq       FilterType  =7   //等于
)



type NrFilterValue struct {
	Value      string    		  `json:"value,omitempty"`  //历史值
	Created    int64   			  `json:"created,omitempty"` //历史时间
	Filter     []NrFilterType     `json:"filter,omitempty"`
}

