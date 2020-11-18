package models

import (
	"time"
)

//Ticket that would be scanned
type Ticket struct {
	BarCode       string `json:"barcode"`
	POS           string `json:"pos"`
	Store         string `json:"store"`
	Trx           string `json:"trx"`
	Folio         string `json:"folio"`
	RutFormato    string `json:"rutFormato"`
	TotalAmount   string `json:"totalAmount"`
	CheckerNumber string `json:"checkerNumber"`
	CheckerName   string `json:"checkerName"`
	Date          string `json:"date"`
	Time          string `json:"time"`
	Products      []Product
	CurrentTime   time.Time
}
