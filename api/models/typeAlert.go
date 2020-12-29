package models

type TypeAlert struct{
	AlertCode           int64    `gorm:"size:100;not null" json:"alertCode"`
	Title           string     `gorm:"size:100;not null" json:"title"`
	Message           string     `gorm:"size:100;not null" json:"message"`
}

func (t *TypeAlert) Prepare(tc TCTicket){
	switch tc.ResponseCode {
	case 108: 
		t.AlertCode = 0
		t.Title = "Fecha Incorrecta"
	case 107: 
		t.AlertCode = 1
		t.Title = "Tienda Incorrecta"
	case 110: 
		t.AlertCode = 2
		t.Title = "Boleta no encontrada"
	case 109: 
		t.AlertCode = 3
		t.Title = "Boleta ya revisada"
	default: 
		t.AlertCode = 0
		t.Title = "Boleta encontrada"
	}
	t.Message = tc.ResponseDescription
}