package utils

import (
	"burndown-ticket/models"
	"encoding/json"
	"fmt"
	"net/http"
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

func ResponseJSON(w http.ResponseWriter, code int, message string, payload interface{}, err models.Error) {
	// response, _ := json.Marshal(payload)
	fmt.Println("RESPONSEJSON")
	response := models.Response{
		Code:    code,
		Message: message,
		Data:    payload,
		Error:   err,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}
