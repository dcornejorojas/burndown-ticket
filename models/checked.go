package models

import (
	"time"
)

//Ticket that would be scanned
type CheckedTicket struct {
	Ticket       string `json:"ticket"`
	Status		int64	`json:"status"`
	POS           int64 `json:"pos"`
	Store         int64 `json:"store"`
	Trx           int64 `json:"trx"`
	Folio         string `json:"folio"`
	RutFormato    string `json:"rutFormato"`
	IdProfile		string `json:"idProfile"`
	TotalAmount   int64 `json:"totalAmount"`
	DateTime          time.Time `json:"dateTime"`
	InitTime          time.Time `json:"initTime"`
	EndTime          time.Time `json:"endTime"`
	Products 	[]Product `json:"products"`
	ProductsNotFound	[]Product `json:"productsNotFound"`
	CurrentTime   time.Time
}