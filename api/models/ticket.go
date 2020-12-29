package models

import (
	"time"
	"os"
	"fmt"
	"encoding/json"
	"errors"
	_ "strings"

	"database/sql/driver"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	_ "github.com/lib/pq"
)

var table = `tickets`


type JSONB []map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	valueString, err := json.Marshal(j)
    return string(valueString), err
}

func (j *JSONB) Scan(value interface{}) error {
    b, ok := value.([]byte)
    if !ok {
        return errors.New("type assertion to []byte failed")
	}
	var result Item
	fmt.Printf("%T - %T", b, result)

	fmt.Println(string(b))
	if err := json.Unmarshal(value.([]byte), &j); err != nil {
		return err
	}
	return nil
}


//Ticket2 that would be scanned
type Ticket2 struct {
	gorm.Model
	IDTicket           string    `gorm:"size:100;not null;primary_key;" json:"IdTicket"`
	Ticket           string    `gorm:"size:100;not null" json:"ticket"`
	Status           int64     `gorm:"size:100;not null" json:"status"`
	POS              string     `gorm:"size:100;not null" json:"pos"`
	Store            string     `gorm:"size:100;not null" json:"store"`
	Trx              string     `gorm:"size:100;not null" json:"trx"`
	Folio            string    `gorm:"size:100;not null" json:"folio"`
	Checked			bool    `gorm:"size:100;not null" json:"checked"`
	CheckerDni       string    `gorm:"size:100;not null" json:"checkerDni"`
	CheckerName       string    `gorm:"size:100;not null" json:"checkerName"`
	TotalAmount      int64     `gorm:"size:100;not null" json:"totalAmount"`
	Date         string `gorm:"size:100;not null" json:"date"`
	Time         string `gorm:"size:100;not null" json:"time"`
	Datetime	time.Time `gorm:"size:100;not null" json:"dateTime"`
	InitTime         time.Time `gorm:"size:100;not null" json:"initTime"`
	EndTime          time.Time `gorm:"size:100;" json:"endTime"`
	HasDifference            bool    `gorm:"size:100;not null" json:"hasDifference"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
	TypeAlert		TypeAlert	`gorm:"size:100;not null" json:"typeAlert"`
}

//Ticket that would be scanned
type Ticket struct {
	ID           string    `sql:"type:uuid_primary_key;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"CreatedAt"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"UpdatedAt"`
	DeleteddAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"DeletedAt"`
	Ticket           string    `gorm:"size:100;not null" json:"ticket"`
	Status           int64     `gorm:"size:100;not null" json:"status"`
	POS              string     `gorm:"size:100;not null" json:"pos"`
	Store            string     `gorm:"size:100;not null" json:"store"`
	Trx              string     `gorm:"size:100;not null" json:"trx"`
	Folio            string    `gorm:"size:100;not null" json:"folio"`
	Checked			bool    `gorm:"size:100;not null" json:"checked"`
	CheckerDni       string    `gorm:"size:100;not null" json:"checkerDni"`
	CheckerName       string    `gorm:"size:100;not null" json:"checkerName"`
	TotalAmount      int64     `gorm:"size:100;not null" json:"totalAmount"`
	Date         string `gorm:"size:100;not null" json:"date"`
	Time         string `gorm:"size:100;not null" json:"time"`
	Datetime	time.Time `gorm:"size:100;not null" json:"dateTime"`
	InitTime         time.Time `gorm:"size:100;not null" json:"initTime"`
	EndTime          time.Time `gorm:"size:100;" json:"endTime"`
	HasDifference            bool    `gorm:"size:100;not null" json:"hasDifference"`
	Products         JSONB `sql:"type:jsonb" json:"products"`
	ProductsNotFound JSONB `gorm:"size:100;" json:"productsNotFound"`
	TypeAlert		TypeAlert	`gorm:"size:100;not null" json:"typeAlert"`
}

//TCTicket is the struct of the ticketChecking integration
type TCTicket struct {
	ResponseCode	int64	`json:"ResponseCode"`
	ResponseDescription	string	`json:"ResponseDescription"`
	GetTransaction	[]Transaction	`json:"GetTransaction"`
}

type Transaction struct {
	Id	string `json:"Id"`
	Store	string `json:"Store"`
	Pos	string `json:"Pos"`
	Transaction	string `json:"Transaction"`
	Date	string `json:"Date"`
	Time	string `json:"Time"`
	CheckerId	string `json:"CheckerId"`
	CheckerNationalId	string `json:"CheckerNationalId"`
	CheckerName	string `json:"CheckerName"`
	Items JSONB `json:"Items"`
	//Items []Item `json:"Items"`
}

func (ticket *Ticket) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.NewV4()
	return scope.SetColumn("ID", uuid.String())
}


//Prepare populate ticket with the object info
func (t *Ticket) MapTicket(obj Transaction, folio string){
	//products, err := json.Marshal(obj.Items)
	strDate := fmt.Sprintf("%s-%s-%sT%s:%s:%sZ",obj.Date[0:4],obj.Date[4:6],obj.Date[6:8],obj.Time[0:2],obj.Time[2:4],obj.Time[4:6])
	datetime, _ := time.Parse(time.RFC3339, strDate)
	fmt.Printf("\n------%v--------------%v-----------------%v----------",obj.Date,obj.Time,datetime)
	t.Ticket = obj.Store + obj.Pos + obj.Transaction
	t.Status = 1
	t.POS = obj.Pos
	t.Store = obj.Store
	t.Trx = obj.Transaction
	t.Folio = folio
	t.Checked = false
	t.HasDifference = false
	t.CheckerDni = obj.CheckerId
	t.TotalAmount = 0
	t.Date = obj.Date
	t.Time = obj.Time
	t.Datetime = datetime
	t.InitTime = time.Now()
	t.Products = obj.Items
	t.ProductsNotFound = JSONB{}
}

func (t *Ticket) Prepare(){
	t.UpdatedAt = time.Now()
	t.Products = JSONB{}
	t.ProductsNotFound = JSONB{}
}

func (t *Ticket) UnmarshalTicket(obj []byte) error {
	err := json.Unmarshal(obj, &t)
	if err != nil {
		return err
	}
	return nil
}


//FindTicketByFolio return a ticket by the given folio
func (t *Ticket) FindTicketByFolio(db *gorm.DB, folio string) (*Ticket, error) {
	var err error
	err = db.Debug().Table(os.Getenv(`DB_SCHEMA`)+`.`+table).Where(`"folio" = ?`, folio).Take(&t).Error
	if gorm.IsRecordNotFoundError(err) {
		return &Ticket{}, nil
	}
	if err != nil {
		return &Ticket{}, err
	}
	return t, err
}

//SaveTicket Save a ticket in the DB
func (t *Ticket) SaveTicket(db *gorm.DB) (*Ticket, error) {
	var err error
	err = db.Debug().Table(os.Getenv(`DB_SCHEMA`)+`.`+table).Create(&t).Error
	if err != nil {
		return &Ticket{}, err
	}
	return t, nil
}

//UpdateTicket Update a ticket in the DB
func (t *Ticket) UpdateTicket(db *gorm.DB, folio string, decoded map[string]interface{}) (*Ticket, error) {
	items, _ := json.Marshal(decoded["products"])
	itemsNotFound, _ := json.Marshal(decoded["productsNotFound"])
	fmt.Println("\n-----------------------ITEMS-----------------------------")
	fmt.Printf("%T\n%+v", items, string(items))
	fmt.Println("\n-----------------------ITEMS NOT FOUND-----------------------------")
	fmt.Printf("%T\n%+v", items, string(itemsNotFound))
	fmt.Println("\n-----------------------Decoded-----------------------------")
	fmt.Printf("%T\n%+v", decoded, decoded)
	db = db.Debug().Table(os.Getenv(`DB_SCHEMA`)+`.`+table).Where(`"folio" = ?`, folio).UpdateColumns(
		map[string]interface{}{
			"ticket": decoded["ticket"],
            "status": decoded["status"],
            "pos": decoded["pos"],
            "store": decoded["store"],
            "trx": decoded["trx"],
            "folio": decoded["folio"],
            "checked": true,
            "checker_dni": decoded["checkerDni"],
            "checker_name": decoded["checkerName"],
            "total_amount": decoded["totalAmount"],
            "date": decoded["date"],
			"time": decoded["time"],
			"datetime": decoded["dateTime"],
            "init_time": decoded["initTime"],
			"end_time": time.Now(),
			"products": items,
			"products_not_found": itemsNotFound,

		},
	)
	if db.Error != nil {
		return &Ticket{}, db.Error
	}
	err := db.Debug().Table(os.Getenv(`DB_SCHEMA`)+`.`+table).Where(`"folio" = ?`, folio).Take(&t).Error
	if err != nil {
		return &Ticket{}, err
	}
	return t,nil
}