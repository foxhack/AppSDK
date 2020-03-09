package nsplus

import "log"

/// init时赋值
//var messageClient messaging.MessageClient


//SDK里定义一个函数指针类型
type InputRecvCallback func(recData *DataDistroWithIndex) error

type ParameterRecvCallback func(recvParam *AppParams) error

type InitClientsCallback func() error


//提供给业务app使用的接口，业务app在初始化时调用
func (sdk *NSPlusSdk) SetInitClientsCBFunc(f InitClientsCallback) {

	log.Printf("SetInitClientsCBFunc f = %v", f)
	sdk.InitClientsCBFunc = f
	//sdk.inputReceivedCB = f
}


func (sdk *NSPlusSdk) SetInputRecvCBFunc(f InputRecvCallback) {
	log.Printf("SetInputRecvCBFunc f = %v", f)
	sdk.InputRecvCBFunc = f
	//sdk.inputReceivedCB = f
}

func (sdk *NSPlusSdk) SetParameterRecvCBFunc(f ParameterRecvCallback) {
	log.Printf("SetParameterRecvCBFunc f = %v", f)
	sdk.ParameterRecvCBfunc = f
	//sdk.inputReceivedCB = f
}



