package models

import (
	"time"
)

//Product that would be scanned
type Product struct {
	BarCode     string                 `json:"code"`
	Name        string                 `json:"name"`
	Quantity    string                 `json:"quantity"`
	Value       string                 `json:"value"`
	Rules       []Rule                 `json:"rule"`
	CheckedInfo map[string]interface{} `json:"checkedInfo"`
	Time        time.Time
}

//AllProducts list of products
type AllProducts []Product
