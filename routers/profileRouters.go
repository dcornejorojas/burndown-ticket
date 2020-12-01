package routers

import (
	"ticket/controllers"

	"github.com/gorilla/mux"
)

//SetProfileRoutes is a func to set routes for Profile Logic
func SetProfileRoutes(router *mux.Router) *mux.Router {

	router.HandleFunc("/profile", controllers.CreateProfile).Methods("POST")
	router.HandleFunc("/profile/list", controllers.ListProfiles).Methods("POST")
	router.HandleFunc("/profile/image", controllers.GetAvatars).Methods("GET")
	router.HandleFunc("/profile/{idProfile}", controllers.UpdateProfile).Methods("PUT")
	router.HandleFunc("/profile/{idProfile}", controllers.GetProfile).Methods("GET")
	router.HandleFunc("/profile/{idProfile}", controllers.DeleteProfile).Methods("DELETE")

	return router
}
