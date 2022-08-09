package utils

import (
	"encoding/json"
	"net/http"
)

type Helper struct{}
type Response struct {
	Code int         `json:"code,omitempty"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type HelperIface interface {
	HttpResp(w http.ResponseWriter, r *http.Request, status int, resp interface{})
	HttpRespError(w http.ResponseWriter, r *http.Request, err error)
}

// func for process and returning non-error Http Response
func (h *Helper) HttpResp(w http.ResponseWriter, r *http.Request, status int, resp interface{}) {
	message := Response{
		Code: status,
		Data: resp,
	}
	dataBytes, _ := json.Marshal(message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(dataBytes)
}

// func for process and returning error Http Response
func (h *Helper) HttpRespError(w http.ResponseWriter, r *http.Request, err error) {
	var status int
	if e, ok := err.(*CustomError); ok {
		status = e.ErrorStatusCode()
	} else {
		status = http.StatusInternalServerError
	}
	message := Response{
		Msg:  err.Error(),
		Code: status,
	}

	dataBytes, err := json.Marshal(message)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(dataBytes)
}

// func for initiate Helper
func InitHelper() HelperIface {
	return &Helper{}
}
