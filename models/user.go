package models

import (
	"time"
)

//User that would login the app
type User struct {
	UserDni  int64  `json:"idUser"`
	Password string `json:"password"`
	Name     string `json:"name"`
	User     string `json:"user"`
	LastName string `json:"lastName"`
	Avatar   string `json:"avatar"`
	Rol      string `json:"rol"`
	Token    string `json:"token"`
	Time     time.Time
}

//AllUsers is a list of users
type AllUsers []User
