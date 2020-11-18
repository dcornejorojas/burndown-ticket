package routers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//RootEndpoint is the handle of root route of the ms
func RootEndpoint(response http.ResponseWriter, request *http.Request) {
	port := os.Getenv("PORT")
	msName := os.Getenv("MS_NAME")

	if port == "" {
		port = ":8000" //localhost
	}
	response.WriteHeader(200)
	response.Write([]byte(fmt.Sprintf("%s Running on PORT %s", msName, port)))
}

//SetUtilsRouter is for set all utils Routes Logic
func SetUtilsRouter(router *mux.Router) *mux.Router {
	router.HandleFunc("/", RootEndpoint).Methods("GET")
	return router
}
