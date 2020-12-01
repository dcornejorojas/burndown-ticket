package models

import (
	"time"
)

//Profile that would use the app
type Profile struct {
	IDProfile int    `json:"idProfile"`
	Name      string `json:"name"`
	LastName  string `json:"lastName"`
	Avatar    string `json:"avatar"`
	Type      string `json:"type"`
	Time      time.Time
}

//AllProfile list of profiles
type AllProfile []Profile
