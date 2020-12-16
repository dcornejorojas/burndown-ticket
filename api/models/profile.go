package models

import (
	"errors"
	"html"
	"os"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

//Profile that would use the app
type Profile struct {
	IDProfile int    `gorm:"size:100;not null" json:"idProfile"`
	Name      string `gorm:"size:100;not null" json:"name"`
	LastName  string `gorm:"size:100;not null" json:"lastName"`
	Avatar    string `gorm:"size:100;not null" json:"avatar"`
	Type      string `gorm:"size:100;not null" json:"type"`
	Time      time.Time
}

//AllProfile list of profiles
type AllProfile []Profile

func (p *Profile) Prepare() {
	p.IDProfile = 0
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.LastName = html.EscapeString(strings.TrimSpace(p.LastName))
	p.Avatar = html.EscapeString(strings.TrimSpace(p.Avatar))
	p.Type = html.EscapeString(strings.TrimSpace(p.Type))
}

func (p *Profile) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if p.Name == "" {
			return errors.New(`campo 'Name' requerido`)
		}
		if p.LastName == "" {
			return errors.New(`campo 'LastName' requerido`)
		}
		if p.Type == "" {
			return errors.New(`campo 'Type' requerido`)
		}

		return nil
	default:
		if p.Name == "" {
			return errors.New(`campo 'Name' requerido`)
		}
		if p.LastName == "" {
			return errors.New(`campo 'LastName' requerido`)
		}
		if p.Type == "" {
			return errors.New(`campo 'Type' requerido`)
		}
		return nil
	}
}

//FindAllProfiles retrieves all profiles in the BD
func (p *Profile) FindAllProfiles(db *gorm.DB) (*[]Profile, error) {
	var err error
	profiles := []Profile{}
	if os.Getenv("DB_ENABLE") == "true" {
		err = db.Debug().Model(&Profile{}).Limit(100).Find(&profiles).Error
		if err != nil {
			return &[]Profile{}, err
		}
	}
	return &profiles, err
}

//FindProfileByID return specific Profile by his ID
func (p *Profile) FindProfileByID(db *gorm.DB, uid uint32) (*Profile, error) {
	var err error
	if os.Getenv("DB_ENABLE") == "true" {
		err = db.Debug().Model(Profile{}).Where("id = ?", uid).Take(&p).Error
		if err != nil {
			return &Profile{}, err
		}
		if gorm.IsRecordNotFoundError(err) {
			return &Profile{}, errors.New("Profile Not Found")
		}
	}
	return p, err
}
