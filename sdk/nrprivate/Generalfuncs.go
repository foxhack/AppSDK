package nrprivate

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	// nari apps

)

/*** =====================   通用函数列表   ======================
WriteJsonFile(path string, fileName string, jsonInfo interface{}) error

///   =====  configuration 必须传入指针（地址）  =====
LoadJsonFromFile(path string, fileName string, configuration interface{}) error

XmlFileLoader(path string, fileName string, config interface{}) error

RunShellCmd(cmdarg string) ([]byte, error)

IsKeyItemInMapStringString(keyItem string, MapRef map[string]string) bool

NewHttpClient() *http.Client

 ====================================================================***/
////
func WriteJsonFile(path string, fileName string, jsonInfo interface{}) error{
	//var filename = "C:/Users/18717/go/src/awesomeProject/res/App.profile"
	if path == "" {
		return errors.New("save to file fail, error[the Path is null]")
	}

	//path = path + "/res"
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Printf("WriteJsonFile: couldn't create absolute path for: %s; %v", path, err.Error())
		return err
	}
	///LoggingClient.Error(fmt.Sprintf("profiles: created absolute path for loading pre-defined Device Profiles: %s", absPath))
	//var profile nari.AppProfile
	fName := fileName
	//lfName := strings.ToLower(fName)
	fullPath := absPath + "/" + fName

	err = os.MkdirAll(absPath, os.ModePerm)
	if err != nil {
		log.Printf("InitfsPaths: os.MkdirAll Path[%s] error[%v]", path, err)
		//return false
	}


	/**_, err = nrprivate.RunShellCmd("rm -f " + fullPath)
	//fmt.Println(fmt.Sprintf("HandlerAppLoad: nrprivate.RunShellCmd cmdout[%v][%s], AppName=[%s]\n", cmdout,splitimagename(string(cmdout)), ImageFile))
	if err != nil {
		//fmt.Fprintf(os.Stdout, " io.Copy   %v.\n", err)
		fmt.Println(fmt.Sprintf("delete temp old file error[%v]", err.Error()))
		///ack.Result = "App镜像文件安装失败！error: " + err.Error()
	}
	lfName := strings.ToLower(fName)
	fullPathTmp := absPath + "/" + lfName
	_, err = nrprivate.RunShellCmd("rm -f " + fullPathTmp)
	//fmt.Println(fmt.Sprintf("HandlerAppLoad: nrprivate.RunShellCmd cmdout[%v][%s], AppName=[%s]\n", cmdout,splitimagename(string(cmdout)), ImageFile))
	if err != nil {
		//fmt.Fprintf(os.Stdout, " io.Copy   %v.\n", err)
		log.Printf("delete temp or old file(ToLower) error[%v]", err.Error())
		///ack.Result = "App镜像文件安装失败！error: " + err.Error()
	}**/

	b, err := json.MarshalIndent(jsonInfo, "", "\t") //json.Marshal(prof)
	if err != nil {
		log.Printf("json.MarshalIndent for jsonInfo[%v], error[%v]", jsonInfo, err.Error())
		return err
	}
	////========== bufio.NewWriter 写文件 =============
	var f *os.File
	f, err = os.Create(fullPath) //创建文件
	if(err != nil) {
		log.Printf("os.Create err %v", err.Error())
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f) //创建新的 Writer 对象
	defer w.Flush()
	_, err = w.WriteString(string(b))
	if(err != nil) {
		log.Printf("WriteString [%v] error[%v]", string(b), err.Error())
		return err
	}

	return nil
}


///   =====  configuration 必须传入指针（地址）  =====
func LoadJsonFromFile(path string, fileName string, configuration interface{}) error {
	log.Printf("LoadJsonFromFile path=%s, file=%s\n", path, fileName)
	//path := determinePath()
	absPath, err := filepath.Abs(path)
	if err != nil {
		errMsg := fmt.Sprintf("LoadJsonFromFile: couldn't create absolute path for: %s; %v", path, err.Error())
		log.Printf(errMsg)
		return err
	}

	//fileName := path + "/" + internal.ConfigFileName //default profile
	if len(absPath) > 0 {
		fileName = absPath + "/" + fileName
	}

	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		msg := fmt.Sprintf("could not load configuration file (%s): %s", fileName, err.Error())
		//LoggingClient.Error(msg)
		return errors.New(msg)
	}
	// Decode the configuration from TOML
	err = json.Unmarshal([]byte(contents), configuration)
	if err != nil {
		msg := fmt.Sprintf("unable to parse configuration file (%s): %s", fileName, err.Error())
		log.Printf("LoadFromFile 5: [%v][%v]\n", configuration, err.Error())
		//LoggingClient.Error(msg)
		return errors.New(msg)
	}
	log.Printf("LoadJsonFromFile success, configuration[%v]\n", configuration)
	return nil
}


///   =====  configuration 必须传入指针（地址）  =====
func XmlFileLoader(path string, fileName string, config interface{}) error {
	log.Printf("XmlFileLoader path=%s, file=%s\n", path, fileName)
	//path := determinePath()
	absPath, err := filepath.Abs(path)
	if err != nil {
		errMsg := fmt.Sprintf("XmlFileLoader: couldn't create absolute path for: %s; %v", path, err.Error())
		log.Printf(errMsg)
		return err
	}

	//fileName := path + "/" + internal.ConfigFileName //default profile
	if len(absPath) > 0 {
		fileName = absPath + "/" + fileName
	}

	contents, err := ioutil.ReadFile(fileName)
	if err != nil {
		msg := fmt.Sprintf("could not load xml file (%s): %s", fileName, err.Error())
		//LoggingClient.Error(msg)
		return errors.New(msg)
	}

	err = xml.Unmarshal(contents, config)
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}

	return nil
}



func RunShellCmd(cmdarg string) ([]byte, error) {
	///arg := "docker stop " + ContainerBuf[containerindex].ContainerID
	cmd := exec.Command("/bin/sh", "-c", cmdarg)
	cmdout, err := cmd.Output()
	if err != nil {
		log.Printf("run shell cmd[%s] fail, error[%v] \n", cmd, string(cmdout), err.Error())
	}
	return cmdout, err
}

func IsKeyItemInMapStringString(keyItem string, MapRef map[string]string) bool {
	_, ok := MapRef[keyItem]
	if ok {
		return true
	}
	return false
}



func NewHttpClient() *http.Client{
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: false,
	}
	HttpClient := &http.Client{Transport: tr}
	return HttpClient
}


func GetHttpPath(Address string, Port int, Path string)string{
	return fmt.Sprintf("http://%v:%v%v", Address, Port, Path)
}