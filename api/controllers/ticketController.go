package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	_ "strconv"
	"ticket/api/models"
	"ticket/api/utils"

	"github.com/gorilla/mux"
	log "github.com/jeanphorn/log4go"
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
func (server *Server) ScanTicket(w http.ResponseWriter, req *http.Request){
	log.LOGGER("Routes").Info("GET - /ticket/ - Trying to retrieve the ticket")
	errObj := models.Error{}
	errObj.NoError()
	store := utils.GetStore()
	vars := mux.Vars(req)
	folio := vars[`folio`]
	fmt.Println(store)
	ticket := models.Ticket{}
	ticketGotten, err := ticket.FindTicketByFolio(server.DB, folio)
	if err != nil {
		log.LOGGER(`Error`).Info(err.Error())
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	if utils.IsValidString(ticket.Ticket) {
		var items []models.Item
		b, _ := json.Marshal(ticket.Products)
		fmt.Println("\n---------------------TYPES--------------------------")
		fmt.Printf("%T --- %T\n", b, items)
		fmt.Printf("\n%v",string(b))
		fmt.Println("\n----------------------ITEMS-------------------------")
		err = json.Unmarshal([]byte(b),&items)
		fmt.Printf("\n%v",items)
		utils.ResponseJSON(w, http.StatusOK, `Ticket de folio `+ folio +` encontrado`, ticketGotten, errObj)
		return
	} else {
		//client := &http.Client{}
		// request.Header.Set("User-Agent", "Mozilla/5.0")
		// request.Header.Set("Content-Type", "application/json")
		//res, err := client.Do(request)
		//url := `http://api.plos.org/search?q=title:DNA`
		url := os.Getenv(`TC_TRANSACTION`) + folio + `/`+ os.Getenv(`USERID`)
		resp, err := http.Get(url)
		
		if err != nil {
			log.LOGGER("Error").Info(err.Error())
			utils.ERROR(w, http.StatusNotFound, err)
			return
		}
		defer resp.Body.Close()
		
		// var example models.Example
		
		// body, err := ioutil.ReadAll(resp.Body)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// err = json.Unmarshal(body,&example)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		
		// fmt.Printf("%+v\n",string(body))
		// fmt.Println(`-----------------------------------------------`)
		// fmt.Printf("%+v\n",example.Response.Docs[0])

		var tc models.TCTicket
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.LOGGER("Error").Info(err.Error())
			utils.ERROR(w, http.StatusNotFound, err)
			return
		}
		//err = json.Unmarshal(body,&tc)
		
		err = json.Unmarshal(body,&tc)
		if err != nil {
			log.LOGGER("Error").Info(err.Error())
			utils.ERROR(w, http.StatusNotFound, err)
			return
		}
		if !utils.IsValidTicket(tc.GetTransaction) {
			ticket.Prepare()
			ticket.TypeAlert.Prepare(tc)
			utils.ResponseJSON(w, http.StatusOK, `Ticket con folio `+ folio +` no fue encontrado`, ticket, errObj)
			return
			
		} else {
			//alert := models.TypeAlert{}
			log.LOGGER("Ticket").Info(`-----------------------------------------------`)
			ticket.MapTicket(tc.GetTransaction[0], folio)
			log.LOGGER("Ticket").Info(string(body))
			log.LOGGER("Ticket").Info(ticket)
			log.LOGGER("Ticket").Info(`-----------------------------------------------`)
			items := []models.Item{}
			b, _ := json.Marshal(ticket.Products)
			fmt.Printf("%T --- %T\n", b, items)
			log.LOGGER("Ticket").Info(b)
			log.LOGGER("Ticket").Info(`-----------------------------------------------`)

			// err = json.Unmarshal(b,&items)
			// if err != nil {
			// 	log.LOGGER("Error").Info(err.Error())
			// 	errObj.HasError(true, http.StatusUnprocessableEntity, "Failed to Unmarshal the items")
			// 	utils.ERROR(w, http.StatusUnprocessableEntity, err)
			// 	return
			// }
			log.LOGGER("Ticket").Info(`-----------------------------------------------`)
			log.LOGGER("Ticket").Info(items)
			newTicket, err := ticket.SaveTicket(server.DB)
			if err != nil {
				log.LOGGER("Error").Info(err.Error())
				errObj.HasError(true, http.StatusUnprocessableEntity, "Failed to create the ticket")
				utils.ERROR(w, http.StatusUnprocessableEntity, err)
				return
			}
			utils.ResponseJSON(w, http.StatusOK, `Ticket con folio `+ folio +` fue encontrado`, newTicket, errObj)
			return
			
		}
		
	}
}


func (server *Server) BurnTicket(w http.ResponseWriter, req *http.Request) {
	errObj := models.Error{}
	errObj.NoError()
	vars := mux.Vars(req)
	folio := vars[`folio`]
	if folio == `` {
		err := errors.New(`Folio no encontrado`)
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	var ticket models.Ticket
	//err = ticket.UnmarshalTicket(body)
	var decoded map[string]interface{}
	fmt.Printf("%T - %T", body, decoded)
	err = json.Unmarshal(body,&decoded)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// tokenID, err := auth.ExtractTokenID(r)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
	// 	return
	// }
	// if tokenID != uint32(uid) {
	// 	responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
	// 	return
	// }

	
	// err = ticket.Validate("update")
	// if err != nil {
	// 	utils.ERROR(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }
	ticket.Prepare()
	
	updatedUser, err := ticket.UpdateTicket(server.DB, folio, decoded)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	utils.ResponseJSON(w, http.StatusOK, fmt.Sprintf("Ticket Quemado: %v", updatedUser.Folio), updatedUser, errObj)


	// errObj := models.Error{}
	// errObj.NoError()
	// var burnedTicket models.CheckedTicket
	// reqBody, err := ioutil.ReadAll(req.Body)
	// if err != nil {
	// 	utils.ResponseJSON(w, http.StatusOK, "Error", "", errObj)
	// 	return
	// }
	// err = json.Unmarshal(reqBody, &burnedTicket)
	// if err != nil {
	// 	fmt.Println(err)
	// 	utils.ResponseJSON(w, http.StatusOK, "Error", "", errObj)
	// 	return
	// }
	// utils.ResponseJSON(w, http.StatusOK, fmt.Sprint("Ticket Quemado: ", burnedTicket.Ticket), burnedTicket, errObj)
}
