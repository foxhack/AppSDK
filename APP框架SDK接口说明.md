# 支持数据集的APP框架SDK接口设计说明

## 功能简述


## 初始化接口  
```golang
func NewNSPlusSdk(appNameKey string) (*NSPlusSdk, error) {

}

//调用示例：
var err error
nsplusdemo.NsplusSdk, err = nsplus.NewNSPlusSdk(“demo”)
if err != nil {
   logBeforeInit(err)
}

```
返回说明：如果sdk初始化失败，则返回error，否则，返回正常的sdk指针；

## 回调函数
### 1.接收input数据  
``` golang
type InputRecvCallback func(recData *DataDistroWithIndex) error
```
``` golang
//注册接收Input数据的回调函数接口
func (sdk *NSPlusSdk) SetInputRecvCBFunc(f InputRecvCallback)
```
``` golang
//最后再实现回调函数
func f() error {}
```
使用参考
``` golang
//example
//注册回调函数
internal.NsplusSdk.SetInputRecvCBFunc(internal.InputCBFuns)
//实现回调函数
func InputCBFuns() (recvParam *nsplus.DataDistroWithIndex) error {
   NsplusSdk.LoggingClient.Error(fmt.Sprintf("InputCBFuns recv data[%v]", recvParam))
   return nil
}
```
### 2.接收parameter更新  
```golang
type ParameterRecvCallback func(recvParam *AppParams) error
```
```golang
func (sdk *NSPlusSdk) SetParameterRecvCBFunc(f ParameterRecvCallback)  
```
使用参考
``` golang
//注册回调函数
internal.NsplusSdk.SetParameterRecvCBFunc(ParameterCBFuns)
//实现回调函数
func ParameterCBFuns(recvParam *nsplus.AppParams) error {
    NsplusSdk.LoggingClient.Error(fmt.Sprintf("ParameterCBFuns recv data[%v]", recvParam))
    return nil
}
```
### 3.用户自定义功能  
```golang
//回调函数声明
type InitClientsCallback func() error
```
```golang
//回调函数定义
func (sdk *NSPlusSdk) SetInputRecvCBFunc(f InputRecvCallback)  
```
使用参考
``` golang
//注册回调函数
internal.NsplusSdk.SetInputRecvCBFunc(InitClientsCBFuns)
//实现回调函数
func InitClientsCBFuns() error {
    NsplusSdk.LoggingClient.Error(fmt.Sprintf("kkkkkkkkkk InitClientsCBFuns now time=[%v]", time.Now()))
   //close(errChan)
   return nil

}
```
## 输出功能接口
接口定义:
```golang
func (sdk *NSPlusSdk)SendAppOutput(data []NameValues) error 
```  
> 功能说明:将NameValues的切片,通过json格式输出到指定TOPIC(目前只有一个OUTPUT)的容器中处理  
> 调用方法: err := sdk.SendAppOutput(Outputs)  
> 参数说明: Outputs为输出的Name-Value数组

数据结构说明:
``` golang
type NameValues struct{
    Name       string        `json:"name"`
    Value      interface{}   `json:"value"`
}
```
调用示例:
```golang
Outputs := make([]NameValues, 0)
//在Outputs的map中添加元素
for _, param := range sdk.AppParamsCtrl.Output {
   var onepara NameValues
   onepara.Name = param.Name
   onepara.Value = fmt.Sprintf("%v_%v", "outValue", sdk.AppParamsCtrl.OutputSqid)

   log.Printf("debugOutputSendLoop for Output [%v] ", onepara)
   Outputs = append(Outputs, onepara)
}
//Outputs 是一个map
err := sdk.SendAppOutput(Outputs)
if err != nil {
   log.Printf("SendAppOutput for Output [%v] fail, error=[%v] ", Outputs, err.Error())
}
```
## 提供查询功能的接口  
MQTT中需要使用的基本数据结构:
```golang
type ResValues struct{
	ResValues        map[string]ValueDefine   `json:"resValues"`     // key resname, value res type and values
}

type ValueDefine struct{
	Value            string              `json:"value"`
	Type             string              `json:"type"`
}
```
### 1.Input查询接口  
``` golang
//define
func (sdk *NSPlusSdk)GetInput(pName string) []AppInputs
```
> 调用示例：**sdk.GetInput("inputname")**  
> 参数说明：pName为空的时候查询所有Input，否则按名称查询;  
> 返回说明：返回切片形式[]AppInputs，如果数组长度为0，表示无结果或查询异常

数据结构说明:
``` golang
type AppInputs struct{
	Param        InputsInfo           `json:"param"`         
	DevResValues map[string]ResValues `json:"devResValues"`  // key devname, value res name and values
}
//MQTT server要求的inputs格式
type InputsInfo struct {
	Name              string            `json:"name"`
	Description       string            `json:"description"`
	Type              string            `json:"type"`
	MaxResCount       int               `json:"maxResCount"`
	Coefficient       float32           `json:"coefficient"`
	Offset            float32           `json:"offset"`
	Reference         string            `json:"reference"`
}

```
### 2.Output查询接口  
``` golang
//define
func (sdk *NSPlusSdk)GetOutput(pName string) []AppOutputs
```
> 调用示例：**sdk.GetOutput("outPutName")**  
> 参数说明：pName为空的时候查询所有Output，否则按名称查询;  
> 返回说明：返回切片形式[]AppOutputs，如果数组长度为0，表示无结果或查询异常

数据结构说明:  
``` golang
type AppOutputs struct{
	Param            OutputsInfo              `json:"param"`   
	Values           ValueDefine              `json:"values"`  
}

type OutputsInfo struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Type        string            `json:"type"`
	Unit        string            `json:"unit"`
	AttrType    string            `json:"attrType"`
}
```

### 3.Params查询接口  
``` golang
//define
func (sdk *NSPlusSdk)GetParams(pName string) []AppParams
```
> 调用示例：**sdk.GetParams("paraName")**  
> 参数说明：pName为空的时候查询所有Param，否则按名称查询;  
> 返回说明：返回切片形式[]AppParams，如果数组长度为0，表示无结果或查询异常

数据结构说明:
```golang
type AppParams struct{
	Param            ParameterInfo            `json:"param"`   
	Values           ValueDefine              `json:"values"`  
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
```
## 其他功能API
### 业务端口查询
>接口定义：(sdk *NSPlusSdk)GetPortConfig() string  
>调用示例：Port := sdk.GetPortConfig ()  
>参数说明：无入参  
>返回说明：返回profile中设定的端口号  

### 获取APP业务路径
>接口定义：(sdk *NSPlusSdk)GetDirConfig() []string  
>调用示例：Paths := sdk.GetDirConfig()  
>参数说明：无入参  
>返回说明：返回profile中设定的业务路径（应为绝对路径），可能是多个  


