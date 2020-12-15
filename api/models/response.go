package models

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

type ResponseProfileList struct {
	Code    float64   `json:"code"`
	Message string    `json:"message"`
	Data    []Profile `json:"data"`
	Error   Error     `json:"error"`
}

type ResponseProfile struct {
	Code    float64 `json:"code"`
	Message string  `json:"message"`
	Data    Profile `json:"data"`
	Error   Error   `json:"error"`
}

type ResponseTicket struct {
	Code     float64   `json:"code"`
	Message  string    `json:"message"`
	Products []Product `json:"products"`
	Error    Error     `json:"error"`
}

type ResponseImage struct {
	Code    float64 `json:"code"`
	Message string  `json:"message"`
	Data    []Image `json:"data"`
	Error   Error   `json:"error"`
}
