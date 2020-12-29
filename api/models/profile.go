package models

import (
	"errors"
	"html"
	"os"
	"strings"
	"time"
	"fmt"

	"github.com/jinzhu/gorm"
)

//Profile that would use the app
type Profile struct {
	gorm.Model
	ID			uint `gorm:"primary_key"`
	IDprofile int       `gorm:"size:100;not null" json:"idProfile"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Lastname  string    `gorm:"size:100;not null" json:"lastName"`
	Avatar    string    `gorm:"size:100;not null" json:"avatar"`
	Type      string    `gorm:"size:100;not null" json:"type"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

//AllProfile list of profiles
type AllProfile []Profile

//Prepare used to init a profile object
func (p *Profile) Prepare() {
	p.IDprofile = 0
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.Lastname = html.EscapeString(strings.TrimSpace(p.Lastname))
	p.Avatar = html.EscapeString(strings.TrimSpace(p.Avatar))
	p.Type = html.EscapeString(strings.TrimSpace(p.Type))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Profile) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if p.Name == "" {
			return errors.New(`campo 'Name' requerido`)
		}
		if p.Lastname == "" {
			return errors.New(`campo 'Lastname' requerido`)
		}
		if p.Type == "" {
			return errors.New(`campo 'Type' requerido`)
		}

		return nil
	default:
		if p.Name == "" {
			return errors.New(`campo 'Name' requerido`)
		}
		if p.Lastname == "" {
			return errors.New(`campo 'Lastname' requerido`)
		}
		if p.Type == "" {
			return errors.New(`campo 'Type' requerido`)
		}
		return nil
	}
}

//SaveProfile Save a profile in the DB
func (p *Profile) SaveProfile(db *gorm.DB) (*Profile, error) {
	fmt.Println(`Save profile`)
	fmt.Println(p)
	var err error
	err = db.Debug().Table(`omnicontrol.profiles`).Create(&p).Error
	if err != nil {
		return &Profile{}, err
	}
	return p, nil
}

//FindAllProfiles retrieves all profiles in the BD
func (p *Profile) FindAllProfiles(db *gorm.DB) (*[]Profile, error) {
	var err error
	profiles := []Profile{}
	if os.Getenv("DB_ENABLE") == "true" {
		err = db.Debug().Table(`omnicontrol.profiles`).Limit(100).Find(&profiles).Error
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

func (p *Profile) UpdateProfile(db *gorm.DB, uid uint32) (*Profile, error) {

	if os.Getenv("DB_ENABLE") == "true" {
		db = db.Debug().Model(&Profile{}).Where("id = ?", uid).Take(&Profile{}).UpdateColumns(
			map[string]interface{}{
				"name":       p.Name,
				"lastname":   p.Lastname,
				"avatar":     p.Avatar,
				"type":       p.Type,
				"updated_at": time.Now(),
			},
		)
		if db.Error != nil {
			return &Profile{}, db.Error
		}
		// This is the display the updated Profile
		err := db.Debug().Model(&Profile{}).Where("id = ?", uid).Take(&p).Error
		if err != nil {
			return &Profile{}, err
		}
	}
	return p, nil
}
