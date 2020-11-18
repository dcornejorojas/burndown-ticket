package models

import (
	"net/http"
)

type Error struct {
	Flag    bool    `json:"flag"`
	Code    float64 `json:"code"`
	Message string  `json:"message"`
}

func (err *Error) NoError() {
	err.Code = http.StatusOK
	err.Flag = false
	err.Message = "Sin Error"
}
