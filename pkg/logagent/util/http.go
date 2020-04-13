package util

import (
	"encoding/json"
	"net/http"
)

type ErrorCode string

const (
	ErrorNone                  ErrorCode = ""
	ErrorInternalError         ErrorCode = "InternalError"
	ErrorInvalidParameter      ErrorCode = "InvalidParameter"
)

type ResponseError struct {
	Code ErrorCode `json:"Code,omitempty"`
	Message string `json:"Message,omitempty"`
}


func WriteResponseError(rw http.ResponseWriter, code ErrorCode, message string) {
	e := ResponseError{
		Code: code,
		Message: message,
	}
	res, err := json.Marshal(e)
	if err != nil {
		rw.Write([]byte("internal error"))
		return
	}
	_, _ = rw.Write(res)
}
