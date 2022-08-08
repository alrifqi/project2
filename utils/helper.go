package utils

import (
	"encoding/json"
	"net/http"
)

type Helper struct{}

type HelperIface interface {
	HttpResp(w http.ResponseWriter, r *http.Request, status int, resp interface{})
}

func (h *Helper) HttpResp(w http.ResponseWriter, r *http.Request, status int, resp interface{}) {
	dataBytes, _ := json.Marshal(resp)
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataBytes)
}

func InitHelper() HelperIface {
	return &Helper{}
}
