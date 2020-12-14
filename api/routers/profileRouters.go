package routers

import (
	"ticket/api/controllers"
	"ticket/api/middlewares"

	"github.com/gorilla/mux"
)

//SetProfileRoutes is a func to set routes for Profile Logic
func SetProfileRoutes(router *mux.Router) *mux.Router {

	router.HandleFunc("/profile", middlewares.SetAuthMiddleware(controllers.CreateProfile)).Methods("POST")
	router.HandleFunc("/profile/list", middlewares.SetAuthMiddleware(controllers.ListProfiles)).Methods("POST")
	router.HandleFunc("/profile/image", middlewares.SetAuthMiddleware(controllers.GetAvatars)).Methods("GET")
	router.HandleFunc("/profile/{idProfile}", middlewares.SetAuthMiddleware(controllers.UpdateProfile)).Methods("PUT")
	router.HandleFunc("/profile/{idProfile}", middlewares.SetAuthMiddleware(controllers.GetProfile)).Methods("GET")
	router.HandleFunc("/profile/{idProfile}", middlewares.SetAuthMiddleware(controllers.DeleteProfile)).Methods("DELETE")

	return router
}
