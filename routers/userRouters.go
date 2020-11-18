package routers

import (
	"burndown-ticket/controllers"

	"github.com/gorilla/mux"
)

//SetUserRoutes is for set all User Routes Logic
func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/user/login", controllers.HelloWorld).Methods("POST")
	router.HandleFunc("/user/logout", controllers.HelloWorld).Methods("POST")
	return router
}
