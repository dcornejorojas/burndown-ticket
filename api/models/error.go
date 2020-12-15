package models

import (
	"net/http"
)

type Error struct {
	Type    bool   `json:"type"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (err *Error) NoError() {
	err.Code = http.StatusOK
	err.Type = false
	err.Message = "Without Errors"
}

func (err *Error) HasError(flag bool, code int, message string) {
	err.Code = code
	err.Type = flag
	err.Message = message
}
