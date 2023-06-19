package httputils

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"io"
	"net/http"
	"runtime"
)

func WrapRpc[RequestType, ResponseType any](mtr *CustomMetrics, rpcHandler func(req RequestType) (*ResponseType, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		if r.Method != http.MethodPost {
			http.Error(w, "unexpect method", http.StatusMethodNotAllowed)
			return
		}
		data, err := io.ReadAll(r.Body)
		if err != nil {
			WriteErrorResponse(w, fmt.Errorf("%w: %v", ReadBodyError, err))
			return
		}

		defer r.Body.Close()

		var rpcRequest RequestType
		if err := json.Unmarshal(data, &rpcRequest); err != nil {
			WriteErrorResponse(w, fmt.Errorf("%w: %v", UnmarshalError, err))
			return
		}

		rpcResponse, err := rpcHandler(rpcRequest)
		if err != nil {
			WriteErrorResponse(w, err)
			return
		}

		responseByte, err := json.Marshal(rpcResponse)
		if err != nil {
			fmt.Println()
			return
		}

		mtr.requestSize.With(prometheus.Labels{"request_url": r.URL.Path}).Add(float64(len(data)))
		mtr.responseSize.With(prometheus.Labels{"response_url": r.URL.Path}).Add(float64(len(responseByte)))
		goroutines := runtime.NumGoroutine()
		mtr.goroutines.Add(float64(goroutines))
		mtr.goroutinesCount.Inc()

		WriteSuccessResponse(w, http.StatusOK, rpcResponse)
	}
}
