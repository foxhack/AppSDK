package nrprivate


const (
	AlmCreate    ="Create"
	AlmRest      = "Reset"
	AlmRescover  ="Rescover"
	AlmUp    =	"Upperlimit"
	AlmDw    =  "LowerLimit"


	ResYc  =   "yc"
	ResYx  =   "yx"
	ResYk  =   "yk"
	ResPr  =   "para"
	AttrType  =  "type"         // 类型  遥测 YC，遥信 YX，遥控 YK，遥调 YT，参数 PARA
	AttrIotName  =  "iotname"      //别名 用于主站上报
	AttrPrType   =   "privatetype"   //私有类型 ，可能存在的二次数据处理
	AttrRelation =   "relation"    //关联数据，遥控专有，关联的遥信信息
	AttrYkType   =   "yktype"       // 遥控专有， 遥控类型 0：单点遥控；1：双点遥控
	AttrClM      =   "ctrlmode"    // 遥控专有，控制模式：0：常规的直接控制；1：常规的SBO控制（先选择再执行）；2：增强安全的直接控制；3：增强安全的SBO控制（先选择再执行
	AttrNodeId   =   "nodeid"      //  104 私有，暂时没有具体的使用场景
	AttrDesc   =   "desc"    //描述字符串


)
//update

const (
	//表的三种变化
	ModTable ="ChangeDevice"
	AddTable  ="DeviceAdd"
	DelTable="DeviceDel"
	//目前需要跟踪的表
	DeviceTable="Device"
	ProFileTable="ProfileName"
)


