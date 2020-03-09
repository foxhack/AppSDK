//
// Copyright (c) 2017 Cavium
//
// SPDX-License-Identifier: Apache-2.0
//

package internal

import (
	//"bytes"
	"encoding/json"
	"fmt"
	"github.com/edgexfoundry/nsplussdk/sdk/correlation"
	"github.com/edgexfoundry/nsplussdk/sdk/telemetry"
	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/gorilla/mux"

	"io"
	"io/ioutil"
	"net/http"
)



func StartHTTPServer(errChan chan error) {
	go func() {
		p := fmt.Sprintf(":%d", 27010)
		NsplusSdk.LoggingClient.Info(fmt.Sprintf("Starting %s %s", AppNameKey, p))

		err := http.ListenAndServe(p, httpServer())
		if err != nil {
			NsplusSdk.LoggingClient.Error(fmt.Sprintf("ListenAndServe with error: %s", err.Error()))
			errChan <- err

			ExitService(err)
			return
		}
	}()
}

// HTTPServer function
func httpServer() http.Handler {
	r := mux.NewRouter()

	// Ping Resource
	r.HandleFunc(clients.ApiPingRoute, pingHandler).Methods(http.MethodGet)

	// Configuration
	///r.HandleFunc(clients.ApiConfigRoute, configHandler).Methods(http.MethodGet)

	// Metrics
	r.HandleFunc(clients.ApiMetricsRoute, metricsHandler).Methods(http.MethodGet)

	r.HandleFunc(clients.ApiNotifyRegistrationRoute, replyNotifyRegistrations).Methods(http.MethodPut)

	///b := r.PathPrefix(clients.ApiBase).Subrouter()
	///AppRestRoutes(b)

	r.HandleFunc("/api/v1/debug/calldata", HandlerGetEdgeConfig).Methods(http.MethodGet)

	r.HandleFunc("/api/v1/debug/appdata", HandlerGetOutput).Methods(http.MethodGet)

	r.HandleFunc("/api/v1/debug/datafilter", HandlerGetSdk).Methods(http.MethodGet)

	r.Use(correlation.ManageHeader)
	r.Use(correlation.OnResponseComplete)
	r.Use(correlation.OnRequestBegin)

	return r
}




func HandlerGetEdgeConfig(w http.ResponseWriter, r *http.Request) {
	/// 解析请求消息中的参数
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",	"Content-Type")   //有使用自定义头 需要这个,Action, Module是例子
	}

	calldata := NsplusSdk.CallAllDataForIndexTable() ///[]DataDistroWithIndex
	NsplusSdk.LoggingClient.Error(fmt.Sprintf("call all data: %v", calldata))


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&calldata)
}

func HandlerGetOutput(w http.ResponseWriter, r *http.Request) {
	/// 解析请求消息中的参数
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",	"Content-Type")   //有使用自定义头 需要这个,Action, Module是例子
	}

	//calldata := NsplusSdk.CallAllDataForIndexTable() ///[]DataDistroWithIndex
	//NsplusSdk.LoggingClient.Error(fmt.Sprintf("call all data: %v", calldata))


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&NsplusSdk.AppParamsCtrl)
}

func HandlerGetSdk(w http.ResponseWriter, r *http.Request) {
	/// 解析请求消息中的参数
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",	"Content-Type")   //有使用自定义头 需要这个,Action, Module是例子
	}

	//calldata := NsplusSdk.CallAllDataForIndexTable() ///[]DataDistroWithIndex
	//NsplusSdk.LoggingClient.Error(fmt.Sprintf("call all data: %v", calldata))


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&NsplusSdk.DataIndexMap)
}



// Test if the service is working
func pingHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("pong"))
}

//func configHandler(w http.ResponseWriter, _ *http.Request) {
//	encode(Configuration, w)
//}

func replyNotifyRegistrations(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		NsplusSdk.LoggingClient.Error(fmt.Sprintf("Failed read body. Error: %s", err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}

	update := models.NotifyUpdate{}
	if err := json.Unmarshal(data, &update); err != nil {
		NsplusSdk.LoggingClient.Error(fmt.Sprintf("Failed to parse %X", data))
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, err.Error())
		return
	}
	if update.Name == "" || update.Operation == "" {
		NsplusSdk.LoggingClient.Error(fmt.Sprintf("Missing json field: %s", update.Name))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if update.Operation != models.NotifyUpdateAdd &&
		update.Operation != models.NotifyUpdateUpdate &&
		update.Operation != models.NotifyUpdateDelete {
		NsplusSdk.LoggingClient.Error(fmt.Sprintf("Invalid value for operation %s", update.Operation))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	//RefreshRegistrations(update)
}

func metricsHandler(w http.ResponseWriter, _ *http.Request) {
	s := telemetry.NewSystemUsage()

	encode(s, w)

	return
}

// Helper function for encoding things for returning from REST calls
func encode(i interface{}, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	err := enc.Encode(i)
	// Problems encoding
	if err != nil {
		NsplusSdk.LoggingClient.Error("Error encoding the data: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}



