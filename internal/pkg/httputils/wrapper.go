package httputils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func WrapRpc[RequestType, ResponseType any](rpcHandler func(req RequestType) (*ResponseType, error)) http.HandlerFunc {
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

		WriteSuccessResponse(w, http.StatusOK, rpcResponse)
	}
}
