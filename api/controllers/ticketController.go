package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"ticket/api/models"
	"ticket/api/utils"

	"github.com/gorilla/mux"
)

var tickets = []byte(`[
	{
	"Ticket": "12475869",
	"Status": 1,
	"Store": 626,
	"POS": 21,
	"Trx": 52,
	"Folio": "1583036823",
	"Date": "2018-04-09T23:00:00Z",
	"Time": "2018-04-09T23:00:00Z",
	"TotalAmount": 11200,
	"HasDifference": true,
	"Products": [{
		"Barcode": "00001234",
		"Name": "EJEMPLO 1",
		"Quantity": "2",
		"Value": "$ 1.990"
	},{
		"Name": "CONFORT",
		"Barcode": "7806500507239",
		"Quantity": "12",
		"Value": "$ 6.590",
		"Rule": [{
			"productID": "7806500507239",
			"store": 626,
			"type": 2,
			"Description": "Producto de Alta Merma, favor revisar."
		}]
	}
	]
},
{
	"Ticket": "89764251",
	"Status": 1,
	"Store": 626,
	"POS": 21,
	"Trx": 52,
	"Folio": "1583036823",
	"Date": "2018-04-09T23:00:00Z",
	"Time": "2018-04-09T23:00:00Z",
	"TotalAmount": 11200,
	"HasDifference": true,
	"Products": [{
		"Barcode": "00001234",
		"Name": "EJEMPLO 1",
		"Quantity": "2",
		"Value": "$ 1.990"
	},{
		"Name": "CONFORT",
		"Barcode": "7806500507239",
		"Quantity": "12",
		"Value": "$ 6.590",
		"Rule": [{
			"productID": "7806500507239",
			"store": 626,
			"format": 1,
			"type": 2,
			"Description": "Producto de Alta Merma, favor revisar."
		}]
	}
	]
}
]`)

/*ScanTicket handle the scan of a ticket.
- {id}: id of the ticket, it will be able with EAN13-CODE128-DUN14 format
*/
func (server *Server) ScanTicket(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	var ticketInfo models.Ticket

	/* Logic to be deprecated */
	var alltickets models.AllTickets
	err2 := json.Unmarshal(tickets, &alltickets)
	if err2 != nil {
		fmt.Println(err2)
	}

	/* Ticket must be found firstly in CheckedTickets (Postgresql) */

	/* Ticket must be found in ticket`s repository (Postgresql) */
	for _, fticket := range alltickets {
		if fticket.Ticket == vars["id"] {
			ticketInfo = fticket
		}
	}
	fmt.Println(ticketInfo)
	errObj := models.Error{}
	if !utils.IsValidString(ticketInfo.Ticket) {
		errObj.HasError(true, http.StatusNotFound, `Ticket no encontrado`)
		utils.ResponseJSON(w, http.StatusNotFound, "No se encontró ticket", ticketInfo, errObj)
	} else {
		errObj.NoError()
		utils.ResponseJSON(w, http.StatusOK, "Ticket encontrado", ticketInfo, errObj)
	}
}

func (server *Server) BurnTicket(w http.ResponseWriter, req *http.Request) {
	errObj := models.Error{}
	errObj.NoError()
	var burnedTicket models.CheckedTicket
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		utils.ResponseJSON(w, http.StatusOK, "Error", "", errObj)
		return
	}
	err = json.Unmarshal(reqBody, &burnedTicket)
	if err != nil {
		fmt.Println(err)
		utils.ResponseJSON(w, http.StatusOK, "Error", "", errObj)
		return
	}
	utils.ResponseJSON(w, http.StatusOK, fmt.Sprint("Ticket Quemado: ", burnedTicket.Ticket), burnedTicket, errObj)
}
