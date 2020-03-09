package define

////======== 数据发布引擎配置文件 ========

import (
	"encoding/json"
	"github.com/edgexfoundry/nsplussdk/sdk/nrprivate"
	"io/ioutil"
	"log"
	"path/filepath"
)

const (
	// nari apps
	UserConfigPath     = "/home/NEdge/nredgex/serviceconfig/"

	APISystemUserConfig     = "/api/v1/edgeconfig"

	EdgeConfigFILE          = "edgeconfig.json"

)

type EdgeConfigStruct struct {
	IOTSetting      EdgeConfigInfo      `json:"IOTSetting"`
	MQTTServer      MQTTServerInfo      `json:"MQTTServer"`

	TopoSetting     TopoSettingApi      `json:"TopoSetting"`
}

type EdgeConfigInfo struct{
	//Host         string      `json:"Host"`
	//Port         string      `json:"Port"`
	EdgeSN       string      `json:"EdgeSN"`
	Name         string      `json:"Name"`
	DeviceKey    string      `json:"DeviceKey"`
}

type MQTTServerInfo struct{
	Schema       string      `json:"Schema"`
	Host         string      `json:"Host"`
	Port         string      `json:"Port"`
	//ClientId     string      `json:"ClientId"`
	User         string      `json:"User"`
	Password     string      `json:"Password"`
	//Topic        string      `json:"Topic"`
}



type TopoSettingApi struct{
	IsRoot        int                 `json:"IsRoot"`   /// 0-拓扑根节点，默认；1-拓扑子节点
	//EdgeSN        string              `json:"EdgeSN"`
	CheckID       string              `json:"CheckID"` /// 接入算法/密钥, 备用字段
	ParentHost    string              `json:"ParentHost"`
	ParentPort    int                 `json:"ParentPort"`
}




func LoadEdgeConfig() EdgeConfigStruct {
	///adding  add dockerapp directly; updating, delete the old and then add the new
	var reqIn EdgeConfigStruct
	var path string
	path = nrprivate.RootPath_Edge/// nari.ConfigDirectory
	if path == "" {
		return reqIn
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Printf("LoadIOTConfig: couldn't create absolute path for: %s; %v", path, err.Error())
		return reqIn
	}
	///LoggingClient.Error(fmt.Sprintf("profiles: created absolute path for loading pre-defined Device Profiles: %s", absPath))
	//var profile nari.AppProfile
	fName := EdgeConfigFILE
	//lfName := strings.ToLower(fName)
	fullPath := absPath + "/" + fName
	log.Printf("LoadIOTConfig: filename [%v]", fullPath)

	////if strings.HasSuffix(lfName, yamlExt) || strings.HasSuffix(lfName, ymlExt) {
	//if strings.HasSuffix(lfName, IOTServerConfigFileExt) {

	yamlFile, err := ioutil.ReadFile(fullPath)
	if err != nil {
		log.Printf("LoadIOTConfig: couldn't read file: %s; %v", fullPath, err.Error())
		return reqIn
	}
	//err = yaml.Unmarshal(yamlFile, &profile)
	err = json.Unmarshal([]byte(yamlFile), &reqIn)
	if err != nil {
		log.Printf("LoadIOTConfig: invalid UserConfig: %s; %v", fullPath, err.Error())
		return reqIn
	}
	log.Printf("LoadIOTConfig : yaml.Unmarshal UserConfig[%v]", reqIn)
	//}
	///重复判断, 并缓存 profileMap
	//profileSliceToMap(profileMapTemp)
	return reqIn
}


