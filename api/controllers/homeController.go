package controllers

import (
	"fmt"
	"net/http"
	"os"
)

/*HomeEndpoint is the root route of the API*/
func (server *Server) HomeEndpoint(response http.ResponseWriter, request *http.Request) {
	port := os.Getenv("PORT")
	msName := os.Getenv("MS_NAME")

	if port == "" {
		port = ":8000" //localhost
	}
	response.WriteHeader(200)
	response.Write([]byte(fmt.Sprintf("%s Running on PORT %s", msName, port)))
}
