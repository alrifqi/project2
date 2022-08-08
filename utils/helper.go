package utils

import (
	"encoding/json"
	"net/http"
)

type Helper struct{}
type ErrorMessage struct {
	Msg string `json:"msg"`
}

type HelperIface interface {
	HttpResp(w http.ResponseWriter, r *http.Request, status int, resp interface{})
	HttpRespError(w http.ResponseWriter, r *http.Request, err error)
}

func (h *Helper) HttpResp(w http.ResponseWriter, r *http.Request, status int, resp interface{}) {
	dataBytes, _ := json.Marshal(resp)
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataBytes)
}

func (h *Helper) HttpRespError(w http.ResponseWriter, r *http.Request, err error) {
	var status int
	if e, ok := err.(*CustomerError); ok {
		status = e.ErrorStatusCode()
	} else {
		status = http.StatusInternalServerError
	}
	message := ErrorMessage{
		Msg: err.Error(),
	}

	dataBytes, err := json.Marshal(message)
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataBytes)
}

func InitHelper() HelperIface {
	return &Helper{}
}
