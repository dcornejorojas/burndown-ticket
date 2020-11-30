package routers

import (
	"ticket/controllers"
	"github.com/gorilla/mux"
)

/*SetTicketRoutes is for set all Ticket Routes Logic
- GET /ticket/{id}: Use to scan a ticket, return the details of ticket
*/
func SetTicketRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/ticket/{id}", controllers.ScanTicket).Methods("GET")
	router.HandleFunc("/ticket", controllers.BurnTicket).Methods("POST")
	return router
}
