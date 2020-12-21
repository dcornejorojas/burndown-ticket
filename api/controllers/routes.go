package controllers

import (
	"ticket/api/middlewares"
)

func (s *Server) InitRoutes() {
	//fs := http.FileServer(http.Dir("./assets"))
	//Home Route
	s.Router.HandleFunc("/", s.HomeEndpoint).Methods("GET")

	//Profile Routes
	s.Router.HandleFunc("/profile", middlewares.SetAuthMiddleware(s.CreateProfile)).Methods("POST")
	s.Router.HandleFunc("/profile/list", middlewares.SetAuthMiddleware(s.ListProfiles)).Methods("POST")
	s.Router.HandleFunc("/profile/image", middlewares.SetAuthMiddleware(s.GetAvatars)).Methods("GET")
	s.Router.HandleFunc("/profile/{idProfile}", middlewares.SetAuthMiddleware(s.UpdateProfile)).Methods("PUT")
	s.Router.HandleFunc("/profile/{idProfile}", middlewares.SetAuthMiddleware(s.GetProfile)).Methods("GET")
	s.Router.HandleFunc("/profile/{idProfile}", middlewares.SetAuthMiddleware(s.DeleteProfile)).Methods("DELETE")

	//Ticket Routes
	s.Router.HandleFunc("/ticket/{folio}", s.ScanTicket).Methods("GET")
	s.Router.HandleFunc("/ticket", s.BurnTicket).Methods("POST")

	//User Routes
	s.Router.HandleFunc("/user/login", s.Login).Methods("POST")
	s.Router.HandleFunc("/user/logout", s.Logout).Methods("POST")

}
