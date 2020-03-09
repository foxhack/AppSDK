//
// Copyright (c) 2017
// Cavium
// Mainflux
// Copyright (c) 2019 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package main

import (
	"flag"
	"fmt"
	"github.com/edgexfoundry/nsplussdk/apps/nsdemo/internal"
	///"github.com/edgexfoundry/nsplussdk/sdk/usage"

	"os"
	"os/signal"
	//"strconv"
	"time"

	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/models"

	"github.com/edgexfoundry/nsplussdk/sdk/correlation"
	"github.com/edgexfoundry/nsplussdk/sdk/nsplus"
	"github.com/edgexfoundry/nsplussdk/sdk/usage"
)

var AppNameKey string

const(
	/// 注：NsPlusAppNameKey 应与app名称一致， 不应为空，尤其在app以单独应用程序运行的情况
	// 建议 app可执行文件名称和profile文件名和reginfo文件名称 保持一致，作为容器app的基础镜像名称
	NsPlusAppNameKey = "nsdemo"
)

func main() {
	var useRegistry bool
	var useProfile string
	start := time.Now()

	AppNameKey = NsPlusAppNameKey
	flag.BoolVar(&useRegistry, "registry", false, "Indicates the service should use Registry.")
	flag.BoolVar(&useRegistry, "r", false, "Indicates the service should use Registry.")
	flag.StringVar(&useProfile, "profile", "", "Specify a profile other than default.")
	flag.StringVar(&useProfile, "p", "", "Specify a profile other than default.")
	flag.Usage = usage.HelpCallback
	flag.Parse()

	//params := startup.BootParams{UseRegistry: useRegistry, UseProfile: useProfile, BootTimeout: internal.BootTimeoutDefault}
	//startup.Bootstrap(params, nsplusdemo.Retry, logBeforeInit)



	//useRegistry = false
	//if ok := nsplusdemo.Init(useRegistry); !ok {
	//	logBeforeInit(fmt.Errorf("%s: Service bootstrap failed", AppNameKey))
	//	os.Exit(1)
	//}

	//nsplusdemo.NsplusSdk.LoggingClient.Info("Service dependencies resolved...")
	//nsplusdemo.NsplusSdk.LoggingClient.Info(fmt.Sprintf("Starting %s %s ", AppNameKey, edgex.Version))

	//http.TimeoutHandler(nil, time.Millisecond*time.Duration(nsplusdemo.Configuration.Service.Timeout), "Request timed out")
	//nsplusdemo.NsplusSdk.LoggingClient.Info(nsplusdemo.Configuration.Service.StartupMsg)

	/// 1、 new 一个 app sdk
	var err error
	internal.NsplusSdk, err = nsplus.NewNsplusApp(AppNameKey)
	if err != nil {
		logBeforeInit(err)
		os.Exit(1)
	}

	///2、 回调函数
	internal.NsplusSdk.SetInitClientsCBFunc(internal.InitClientsCBFuns)

	internal.NsplusSdk.SetInputRecvCBFunc(internal.InputCBFuns)

	internal.NsplusSdk.SetParameterRecvCBFunc(internal.ParameterCBFuns)


	errs := make(chan error, 2)

	internal.StartHTTPServer(errs)

	listenForInterrupt(errs)


	/// 3、回调函数赋值后，启动app sdk
	err = internal.NsplusSdk.SdkRun()
	if err != nil {
		internal.NsplusSdk.LoggingClient.Info(fmt.Sprintf("SdkRun fail, error =[%s] ", err.Error()))
		os.Exit(1)
	}

	///internal.Loop(errs)

	// Time it took to start service
	internal.NsplusSdk.LoggingClient.Info("Service started in: " + time.Since(start).String())
	//nsdemo.NsplusSdk.LoggingClient.Info("Listening on port: " + strconv.Itoa(nsdemo.Configuration.Service.Port))
	c := <-errs
	internal.Destruct()
	internal.NsplusSdk.LoggingClient.Warn(fmt.Sprintf("terminating: %v", c))

	close(errs)
	os.Exit(0)
}

func logBeforeInit(err error) {
	l := logger.NewClient(AppNameKey, false, "", models.InfoLog)
	l.Error(err.Error())
}

func listenForInterrupt(errChan chan error) {
	go func() {
		correlation.LoggingClient = internal.NsplusSdk.LoggingClient
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt)
		errChan <- fmt.Errorf("%s", <-c)
	}()
}
