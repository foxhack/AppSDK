package define


type EdgeAuthInfo3900 struct {
	DevSN       string                            `json:"devSn"`
	EdgeId      int64                             `json:"edgeId"`

	EdgeReg     EdgeRegInfo3900                   `json:"EdgeReg"`
	Topics      map[string]TopicListjson          `json:"topics"`

	Devices     map[string]DeviceAuthInfo3900     `json:"devices"`

	AppList     map[string]string                 `json:"appList"`
}

type EdgeRegInfo3900 struct {
	DevId         string             `json:"devId"`      /// 当前注册设备在IOT系统内部的ID
	Tenancy       int64              `json:"tenancy"`    ///当前分配资源的租期，在租期范围内，IOT Agent每次启动时不用重新认证
	ChannelList   []ChannelListjson  `json:"channelList"`
	TopicList     []TopicListjson    `json:"topicList"`
	Profile       Profilejson        `json:"profile"`
}

type ChannelListjson struct {
	ChannelID     string             `json:"channelID"` ///通道ID
	Protocol      string             `json:"protocol"`  ///通道采用的协议，取值范围为：http、https、mqtt、mqtts、mqtt-ws、mqtt-wss
	Addr          string             `json:"addr"`      ///通道地址，如果是IP和端口则需组装成“ip:port”的形式
}
type TopicListjson struct {
	TopicId       string             `json:"topicId"`   ///topicID
	Topic         string             `json:"topic"`     ///topic
	Direction     string             `json:"direction"` ///up、down
	Qos           int                `json:"qos"`
}
type Profilejson struct {
	Name          string             `json:"name"`     ///
	Addr          string             `json:"addr"`     ///文件在文件共享系统中的路径
	Size          string             `json:"size"`     ///文件的大小，单位为字节
	Md5           string             `json:"md5"`      ///文件的MD5值
}

type DeviceAuthInfo3900 struct {
	DevId         string              `json:"devId"`

	OnOff         int                 `json:"onOff"`
	ForceOff      bool                `json:"forceOff"`
	ServiceInfo   DevServiceInfo3900  `json:"serviceInfo"`
	//DevSN       string                       `json:"devSn"`
	//Manufacture string                       `json:"manufacture"`
	//Model       string                       `json:"model"`
	//Name        string                       `json:"name"`

	DevInfo       DeviceRegBody       `json:"devInfo"`
}



type DevServiceInfo3900 struct{
	Name          string              `json:"name"`          // Unique name for identifying a device ， use as deviceSN
	///LastConnected  int64           `json:"lastConnected"` // Time (milliseconds) that the device last provided any feedback or responded to any request
	///LastReported   int64           `json:"lastReported"`  // Time (milliseconds) that the device reported data to the core microservice
	//Service        DeviceService    `json:"service"`       // Associated Device Service - One per device
	Address       string               `json:"address"`      // Address of the addressable  "address": "192.168.43.150",
	Port          int                  `json:"port,Number"`  // Port for the address        "port": 49982,
	//Profile      DeviceProfile       `json:"profile"`      // Associated Device Profile - Describes the device
	ProfileName   string               `json:"profilename"`  // Non-database identifier (must be unique)

	///ResourceNameType   []ResourceNameTypeInfo   `json:"resourcenametype"`
}





