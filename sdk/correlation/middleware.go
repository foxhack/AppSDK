package correlation

import (
	"context"
	"github.com/edgexfoundry/nsplussdk/sdk/nsplus"
	"github.com/google/uuid"
	"net/http"
	"time"

	///"github.com/edgexfoundry/edgex-backend/internal"
	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
)

var LoggingClient logger.LoggingClient

func ManageHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hdr := r.Header.Get(clients.CorrelationHeader)
		if hdr == "" {
			hdr = uuid.New().String()
		}

		ctx := context.WithValue(r.Context(), clients.CorrelationHeader, hdr)
		w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
		//w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
		//w.Header().Set("content-type", "application/json")             //返回数据格式是json
		w.Header().Set("hgftest", "hgftest")             //返回数据格式是json
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func OnResponseComplete(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now()
		next.ServeHTTP(w, r)
		correlationId := FromContext(r.Context())
		if LoggingClient != nil {
			LoggingClient.Trace("Response complete", clients.CorrelationHeader, correlationId, nsplus.LogDurationKey, time.Since(begin).String())
		}
	})
}

func OnRequestBegin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationId := FromContext(r.Context())
		if LoggingClient != nil {
			LoggingClient.Trace("Begin request", clients.CorrelationHeader, correlationId, "path", r.URL.Path)

		}
		next.ServeHTTP(w, r)
	})
}
