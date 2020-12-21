package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"ticket/api/auth"
	"ticket/api/models"
	"ticket/api/utils"

	"golang.org/x/crypto/bcrypt"
)

//Login used to login the user in the API, creates the token and retrieve the info if the login is correct.
func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	//var data []string
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	fmt.Println(user)
	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.User, user.Password)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	fmt.Println(token)
	user.Token = token
	err2 := models.Error{}
	err2.NoError()
	utils.ResponseJSON(w, http.StatusOK, "Login correcto", user, err2)
}

//Logout used to logout the user by the given id. Delete the token.
func (server *Server) Logout(w http.ResponseWriter, r *http.Request) {
	//var data []string
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	fmt.Println(user)
	err2 := models.Error{}
	err2.NoError()
	utils.ResponseJSON(w, http.StatusOK, "Login correcto", user, err2)
}

/*SignIn used to search the user by his user and password in the DB*/
func (server *Server) SignIn(username, password string) (string, error) {
	var err error

	user := models.User{}
	if os.Getenv("DB_ENABLE") == "true" {
		err = server.DB.Debug().Model(models.User{}).Where("user = ?", username).Take(&user).Error
		if err != nil {
			return "", err
		}
		err = models.VerifyPassword(user.Password, password)
		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
			return "", err
		}
	}
	return auth.CreateToken(user.IDUser)
}
