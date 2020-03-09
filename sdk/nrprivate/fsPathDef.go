package nrprivate

import (
	"log"
	"os"
)

/// 文件系统目录定义和初始化操作

const(
	/// ==== files
	IOTdevTransFileExt        =".xml"

	///  ==== dirs

	RootPath_FS                = "/home/NEdge/fsroot/"


	MODULENAME_IOT                = "iot"
	RootPath_Edge                 = "/home/NEdge/fsroot/edge/"

	RootPath_Edge_global          = "/home/NEdge/fsroot/edge/global/"

	RootPath_Engine               = "/home/NEdge/fsroot/engine/"

	RootPath_Engine_IOTConfig     = "/home/NEdge/fsroot/engine/config/"

	RootPath_NS3900IOT            = "/home/NEdge/fsroot/ns3900iot/"
	RootPath_NS3900IOT_Profiles   = "/home/NEdge/fsroot/ns3900iot/profiles/"
	RootPath_NS3900IOT_XMLfiles   = "/home/NEdge/fsroot/ns3900iot/xmls/"
	RootPath_NS3900IOT_ResDesc    = "/home/NEdge/fsroot/ns3900iot/ResDesc/"
	RootPath_NS3800IOT            = "/home/NEdge/fsroot/ns3800iot/"

	MODULENAME_TOPO               = "topo"
	RootPath_Topo                 = "/home/NEdge/fsroot/topo/"
	RootPath_Topo_Topofile        = "/home/NEdge/fsroot/topo/topofile/"


	RootPath_Version              = "/home/NEdge/fsroot/version/"


	MODULENAME_NSPLUSApp          = "nsplus"
	RootPath_Nsplus               = "/home/NEdge/fsroot/nsplus/"
	RootPath_Nsplus_Pofile        = "/home/NEdge/fsroot/nsplus/profile/"
	RootPath_Nsplus_PofileTmp     = "/home/NEdge/fsroot/nsplus/profiletmp/"
	RootPath_Nsplus_DataSetConf   = "/home/NEdge/fsroot/nsplus/dataset/"


)


func InitfsPaths(module string){

	MakeDirAllforPath(RootPath_Version)

	switch module {
	case MODULENAME_IOT:
		MakeDirAllforPath(RootPath_Edge)
		MakeDirAllforPath(RootPath_Edge_global)

		MakeDirAllforPath(RootPath_Engine)
		MakeDirAllforPath(RootPath_Engine_IOTConfig)

		MakeDirAllforPath(RootPath_NS3900IOT)
		MakeDirAllforPath(RootPath_NS3900IOT_Profiles)
		MakeDirAllforPath(RootPath_NS3900IOT_XMLfiles)
		MakeDirAllforPath(RootPath_NS3900IOT_ResDesc)

		MakeDirAllforPath(RootPath_NS3800IOT)

	case MODULENAME_TOPO:
		MakeDirAllforPath(RootPath_Topo)
		MakeDirAllforPath(RootPath_Topo_Topofile)
		MakeDirAllforPath(RootPath_NS3900IOT_XMLfiles)
		MakeDirAllforPath(RootPath_NS3900IOT_ResDesc)

	case MODULENAME_NSPLUSApp:
		MakeDirAllforPath(RootPath_Nsplus)
		MakeDirAllforPath(RootPath_Nsplus_Pofile)
		MakeDirAllforPath(RootPath_Nsplus_PofileTmp)

		MakeDirAllforPath(RootPath_Nsplus_DataSetConf)
	default:
		//return fmt.Errorf("Invalid update operation")
	}

}


func MakeDirAllforPath(path string){
	path = RootPath_Edge
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Printf("InitfsPaths: os.MkdirAll Path[%s] error[%v]", path, err)
		//return false
	}
}
