package controllers

import (
	"burndown-ticket/models"
	"burndown-ticket/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/*ScanTicket handle the scan of a ticket.
- {id}: id of the ticket, it will be able with EAN13-CODE128-DUN14 format
*/
func ScanTicket(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	profileID, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vars)
	fmt.Println(profileID)
	errObj := models.Error{}
	errObj.NoError()
	utils.ResponseJSON(w, http.StatusOK, fmt.Sprint(len(allProfiles), " perfiles encontrados"), allProfiles, errObj)
}
