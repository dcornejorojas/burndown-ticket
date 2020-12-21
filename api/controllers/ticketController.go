package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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
- {folio}: id of the ticket, it will be able with EAN13-CODE128-DUN14 format
*/
func (server *Server) ScanTicket(w http.ResponseWriter, req *http.Request) {

	store := utils.GetStore()
	service := os.Getenv("GET_TICKET")
	vars := mux.Vars(req)
	folio := vars[`folio`]
	if len(folio) > 12 {
		err := errors.New("Folio demasiado largo")
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	url := fmt.Sprintf("%s%s%s%s%s", service, "/", store, "/", folio)
	fmt.Println(url)

	var ticketInfo models.Ticket
	client := &http.Client{}
	request, _ := http.NewRequest("GET", "http://golang.org", nil)
	request.Header.Set("User-Agent", "Mozilla/5.0")
	request.Header.Set("Content-Type", "application/json")
	res, err := client.Do(request)
	if err != nil {
		utils.ERROR(w, http.StatusNotFound, err)
		return
	}
	fmt.Println(res.StatusCode)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		utils.ERROR(w, http.StatusNotFound, err)
		return
	}
	fsb := string(body)
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
		utils.ResponseJSON(w, http.StatusOK, "Ticket encontrado", fsb, errObj)
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
