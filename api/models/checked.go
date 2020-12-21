package models

import (
	"time"
)

//CheckedTicket that have been scanned
type CheckedTicket struct {
	Ticket           string    `gorm:"size:100;not null" json:"ticket"`
	Status           int64     `gorm:"size:100;not null" json:"status"`
	POS              int64     `gorm:"size:100;not null" json:"pos"`
	Store            int64     `gorm:"size:100;not null" json:"store"`
	Trx              int64     `gorm:"size:100;not null" json:"trx"`
	Folio            string    `gorm:"size:100;not null" json:"folio"`
	RutFormato       string    `gorm:"size:100;not null" json:"rutFormato"`
	IDProfile        string    `gorm:"size:100;not null" json:"idProfile"`
	TotalAmount      int64     `gorm:"size:100;not null" json:"totalAmount"`
	DateTime         time.Time `gorm:"size:100;not null" json:"dateTime"`
	InitTime         time.Time `gorm:"size:100;not null" json:"initTime"`
	EndTime          time.Time `gorm:"size:100;not null" json:"endTime"`
	Products         []Product `gorm:"size:100;not null" json:"products"`
	ProductsNotFound []Product `gorm:"size:100;not null" json:"productsNotFound"`
	CurrentTime      time.Time
}
