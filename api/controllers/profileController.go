package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"ticket/api/models"
	"ticket/api/utils"

	"github.com/gorilla/mux"
)

var allProfiles = models.AllProfile{}

var allImages = models.Images{
	{
		IDImage: 1,
		Name:    "emptyprofile",
	},
	{
		IDImage: 2,
		Name:    "emptyprofile",
	},
	{
		IDImage: 3,
		Name:    "emptyprofile",
	},
	{
		IDImage: 4,
		Name:    "emptyprofile",
	},
	{
		IDImage: 5,
		Name:    "emptyprofile",
	},
	{
		IDImage: 6,
		Name:    "emptyprofile",
	},
	{
		IDImage: 7,
		Name:    "emptyprofile",
	},
	{
		IDImage: 8,
		Name:    "defaultprofile",
	},
	{
		IDImage: 9,
		Name:    "emptyprofile",
	},
}

//GetAvatars return the list of the avatars in the DB
func (server *Server) GetAvatars(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAvatars")
	var errObj = models.Error{}
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		errObj.HasError(true, http.StatusUnprocessableEntity, "Failed to process Entity")
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	for i := range allImages {
		allImages[i].Path = fmt.Sprint(path, "/assets/", allImages[i].Name, ".svg")
		fmt.Println(allImages[i].Path)
	}

	errObj.NoError()
	utils.ResponseJSON(w, http.StatusCreated, "Perfil agregado exitosamente", allImages, errObj)
}

//CreateProfile creates a new profile in the environment
func (server *Server) CreateProfile(w http.ResponseWriter, r *http.Request) {
	var profile models.Profile
	var errObj = models.Error{}
	errObj.NoError()

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errObj.HasError(true, http.StatusUnprocessableEntity, "Failed to process Entity")
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = json.Unmarshal(reqBody, &profile)
	if err != nil {
		errObj.HasError(true, http.StatusUnprocessableEntity, `Failed to process Entity`)
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	profile.Prepare()
	err = profile.Validate("")
	if err != nil {
		errObj.HasError(true, http.StatusUnprocessableEntity, `Failed to process Entity`)
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	//INSERT PROFILE IN DB
	allProfiles = append(allProfiles, profile)

	utils.ResponseJSON(w, http.StatusCreated, "Perfil agregado exitosamente", allProfiles, errObj)

}

//ListProfiles shows a profile list
func (server *Server) ListProfiles(w http.ResponseWriter, r *http.Request) {
	profile := models.Profile{}
	length := 0
	profiles, err := profile.FindAllProfiles(server.DB)
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	if reflect.TypeOf(profiles).Kind() == reflect.Slice {
		length = reflect.ValueOf(profiles).Len()
	}

	err2 := models.Error{}
	err2.NoError()
	utils.ResponseJSON(w, http.StatusOK, fmt.Sprint(length, " perfiles encontrados"), profiles, err2)
}

//GetProfile returns profile info by the given ID
func (server *Server) GetProfile(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["idProfile"], 10, 32)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	profile := models.Profile{}
	profGotten, err := profile.FindProfileByID(server.DB, uint32(uid))
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}

	//VALID INFO OF THE PROFILE
	err2 := models.Error{}
	err2.NoError()
	utils.ResponseJSON(w, http.StatusOK, `Usuario encontrado`, profGotten, err2)
}

//UpdateProfile can update the info of a given idProfile
func (server *Server) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["idProfile"], 10, 32)
	if err != nil {
		utils.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	profile := models.Profile{}
	err = json.Unmarshal(body, &profile)
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	profile.Prepare()
	err = profile.Validate("update")
	if err != nil {
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	//VALIDATE PROF DATA
	updatedProf, err := profile.UpdateProfile(server.DB, uint32(uid))
	if err != nil {
		utils.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	err2 := models.Error{}
	err2.NoError()
	utils.ResponseJSON(w, http.StatusOK, `Perfil actualizado`, updatedProf, err2)

}

//DeleteProfile delete a profile by the given IdProfile
func (server *Server) DeleteProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var deleted = false
	var response200 models.Response
	profileID, err := strconv.Atoi(vars["idProfile"])
	err200 := models.Error{
		Message: "Sin Error",
		Code:    http.StatusOK,
		Type:    false,
	}
	if err != nil {
		fmt.Fprintf(w, "Invalid profile ID")
		return
	}

	for i, profile := range allProfiles {
		if profile.IDProfile == profileID {
			deleted = true
			allProfiles = append(allProfiles[:i], allProfiles[i+1:]...)
		}
	}
	if !deleted {
		response200 = models.Response{
			Code:    201,
			Message: fmt.Sprintf("El usuario con ID %v no fue encontrado", profileID),
			Data:    allProfiles,
			Error:   err200,
		}
	} else {
		response200 = models.Response{
			Code:    201,
			Message: fmt.Sprintf("El usuario con ID %v ha sido eliminado correctamente", profileID),
			Data:    allProfiles,
			Error:   err200,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response200)
}
