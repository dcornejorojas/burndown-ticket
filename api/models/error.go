package models

import (
	"net/http"
)

//Error is the struct of the error given by the API
type Error struct {
	Type    bool   `json:"type"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

//NoError create a error object that doesnt have any error
func (err *Error) NoError() {
	err.Code = http.StatusOK
	err.Type = false
	err.Message = "Without Errors"
}

//HasError create a error object that have a error
func (err *Error) HasError(flag bool, code int, message string) {
	err.Code = code
	err.Type = flag
	err.Message = message
}
