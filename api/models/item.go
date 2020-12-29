package models

import (
	_ "database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"fmt"

	"github.com/jinzhu/gorm"
)



//Item that would be scanned
type Item struct {
	gorm.Model
	//BarCode     string                 `json:"barcode"`
	//Name        string                 `json:"name"`
	//Quantity    string                 `json:"quantity"`
	//Value       string                 `json:"value"`
	//CheckedInfo map[string]interface{} `json:"checkedInfo"`
	
	IdItem				string	`gorm:"size:100;not null" json:"id,omitempty"`
	ItemName			string	`gorm:"size:100;not null" json:"name,omitempty"`
	ItemBarCode			string	`gorm:"size:100;not null" json:"barcode,omitempty"`
	ItemQuantity		string	`gorm:"size:100;not null" json:"quantity,omitempty"`
	ItemWeight			string	`gorm:"size:100;not null" json:"weight,omitempty"`
	ItemTotalAmount		string	`gorm:"size:100;not null" json:"value,omitempty"`
	ItemRecall			string	`gorm:"size:100;not null" json:"recall,omitempty"`
	ItemCategory		string	`gorm:"size:100;not null" json:"category,omitempty"`	
	ItemRange			bool	`gorm:"size:100;not null" json:"range,omitempty"`
	ItemChecked			bool	`gorm:"size:100;not null" json:"checked,omitempty"`
	CheckedEditInfo 	map[string]interface{} `gorm:"size:100;" json:"checkedEditInfo"`
	Rules       		[]Rule                 `gorm:"size:100;" json:"rule"`
	TicketID	uint
	Ticket		Ticket `gorm:"foreignkey:TicketID"`
}

//AllProducts list of products
type AllProducts []Item

func (n *Item) UnmarshalJSON(buf []byte) error {
	tmp := []interface{}{
		&n.IdItem,
		&n.ItemName,
		&n.ItemBarCode,
		&n.ItemQuantity,
		&n.ItemWeight,
		&n.ItemTotalAmount,
		&n.ItemRecall,
		&n.ItemCategory,
		&n.ItemRange,
		&n.ItemChecked,
		&n.CheckedEditInfo,
		&n.Rules,
	}
	wantLen := len(tmp)
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in Notification: %d != %d", g, e)
	}
	return nil
}
