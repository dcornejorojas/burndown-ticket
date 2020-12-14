package controllers

import (
	"fmt"
	"net/http"
	u "ticket/api/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HelloWorld")
	resp := map[string]interface{}{
		"Name": "Eve",
		"Age":  "6.0",
	}
	u.Respond(w, resp)
}
