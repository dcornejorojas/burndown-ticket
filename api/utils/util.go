package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ticket/api/models"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// func Respond(w http.ResponseWriter, data map[string]interface{}) {
// 	w.Header().Add("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(data)
// }

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func ResponseJSON(w http.ResponseWriter, code int, message string, payload interface{}, obj interface{}) {
	// response, _ := json.Marshal(payload)
	response := models.Response{
		Code:    code,
		Message: message,
		Data:    payload,
		Error:   obj,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, obj error) {
	var data interface{}
	err := models.Error{}
	if obj != nil {
		err.HasError(true, statusCode, obj.Error())
		ResponseJSON(w, statusCode, err.Message, data, err)
		return
	}
	err.NoError()
	ResponseJSON(w, statusCode, err.Message, data, err)
}
