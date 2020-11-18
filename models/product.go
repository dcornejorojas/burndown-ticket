package models

import (
	"time"
)

//Product that would be scanned
type Product struct {
	BarCode  string `json:"barcode"`
	Name     string `json:"name"`
	Quantity string `json:"quantity"`
	Price    string `json:"price"`
	Rules    []Rule
	Time     time.Time
}
