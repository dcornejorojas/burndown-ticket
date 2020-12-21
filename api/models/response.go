package models

//Response struct of the response object
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

//ResponseProfileList struct of the response profile list
type ResponseProfileList struct {
	Code    float64   `json:"code"`
	Message string    `json:"message"`
	Data    []Profile `json:"data"`
	Error   Error     `json:"error"`
}
