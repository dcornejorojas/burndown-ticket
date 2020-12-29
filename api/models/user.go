package models

import (
	"errors"
	"html"
	"strings"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//User that would login the app
type User struct {
	gorm.Model
	IDUser    uint32    `gorm:"primary_key;auto_increment" json:"iduser"`
	Dni       string    `gorm:"size:100;not null" json:"dni"`
	Password  string    `gorm:"size:100;not null" json:"password"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	User      string    `gorm:"size:100;not null" json:"user"`
	LastName  string    `gorm:"size:100;not null" json:"lastname"`
	Avatar    string    `gorm:"size:200;not null" json:"avatar"`
	Rol       string    `gorm:"size:100;not null" json:"rol"`
	Token     string    `gorm:"size:200;not null" json:"token"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

//AllUsers is a list of users
type AllUsers []User

//Hash is for hashed password
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//VerifyPassword check if the password is the same as the hashed one
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

//BeforeSave is used to hash the password of the user
func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

//Prepare the data before save User
func (u *User) Prepare() {
	u.IDUser = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.LastName = html.EscapeString(strings.TrimSpace(u.LastName))
	u.Rol = html.EscapeString(strings.TrimSpace(u.Rol))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

//Validate is used to check if the user info insn´t empty
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Name == "" {
			return errors.New(`campo 'Name' requerido`)
		}
		if u.Password == "" {
			return errors.New(`campo 'Password' requerido`)
		}
		if u.LastName == "" {
			return errors.New(`campo 'LastName' requerido`)
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New(`campo 'Password' requerido`)
		}
		if u.User == "" {
			return errors.New(`campo 'User' requerido`)
		}
		return nil

	default:
		if u.Name == "" {
			return errors.New(`campo 'Name' requerido`)
		}
		if u.Password == "" {
			return errors.New(`campo 'Password' requerido`)
		}
		if u.LastName == "" {
			return errors.New(`campo 'LastName' requerido`)
		}
		return nil
	}
}

//SaveUser Save a user in the DB
func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	fmt.Println(`Save user`)
	fmt.Println(u)
	var err error
	err = db.Debug().Table(`omnicontrol.users`).Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

//FindAllUsers Return all users from the DB
func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

//FindUserByID return a user by the given userID
func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}
