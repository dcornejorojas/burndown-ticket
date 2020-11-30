package models

import (
	"net/http"
)

type Error struct {
	Type    bool    `json:"type"`
	Code    float64 `json:"code"`
	Message string  `json:"message"`
}

func (err *Error) NoError() {
	err.Code = http.StatusOK
	err.Type = false
	err.Message = "Without Errors"
}

func (err *Error) HasError(flag bool, code float64, message string){
	err.Code = code
	err.Type = flag
	err.Message = message
}
