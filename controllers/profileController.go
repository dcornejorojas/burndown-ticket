package controllers

import (
	"ticket/models"
	"ticket/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

var allProfiles = models.AllProfile{}

var allImages = models.Images{
	{
		IDImage: 1,
		Name:    "emptyUser",
	},
	{
		IDImage: 2,
		Name:    "emptyUser",
	},
	{
		IDImage: 3,
		Name:    "emptyUser",
	},
	{
		IDImage: 4,
		Name:    "emptyUser",
	},
	{
		IDImage: 5,
		Name:    "emptyUser",
	},
	{
		IDImage: 6,
		Name:    "emptyUser",
	},
	{
		IDImage: 7,
		Name:    "emptyUser",
	},
	{
		IDImage: 8,
		Name:    "defaultUser",
	},
	{
		IDImage: 9,
		Name:    "emptyUser",
	},
}

func GetAvatars(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAvatars")
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	err200 := models.Error{
		Message: "Sin Error",
		Code:    201,
		Type:    false,
	}
	for i := range allImages {
		allImages[i].Path = fmt.Sprint(path, "/assets/", allImages[i].Name, ".svg")
		fmt.Println(allImages[i].Path)
	}
	response200 := models.ResponseImage{
		Code:    201,
		Message: "Lista de Imagenes",
		Data:    allImages,
		Error:   err200,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response200)
}

func CreateProfile(w http.ResponseWriter, r *http.Request) {
	var newProfile models.Profile
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task Data")
	}
	err200 := models.Error{
		Message: "Sin Error",
		Code:    201,
		Type:    false,
	}

	json.Unmarshal(reqBody, &newProfile)
	newProfile.IDProfile = int(len(allProfiles) + 1)
	allProfiles = append(allProfiles, newProfile)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response200 := models.Response{
		Code:    201,
		Message: "Perfil creado exitosamente",
		Data:    allProfiles,
		Error:   err200,
	}
	json.NewEncoder(w).Encode(response200)

}

func ListProfiles(w http.ResponseWriter, r *http.Request) {
	err := models.Error{}
	err.NoError()
	utils.ResponseJSON(w, http.StatusOK, fmt.Sprint(len(allProfiles), " perfiles encontrados"), allProfiles, err)
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	profileID, err := strconv.Atoi(vars["idProfile"])
	if err != nil {
		err200 := models.Error{
			Message: "ID Invalido",
			Code:    400,
			Type:    true,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err200)
	}
	var profileInfo = models.Profile{}
	w.Header().Set("Content-Type", "application/json")
	// reqBody, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	err200 := models.Error{
	// 		Message: "Ingresar data valida para actualizar",
	// 		Code:    400,
	// 		Type:    true,
	// 	}
	// 	json.NewEncoder(w).Encode(err200)
	// }
	for _, profile := range allProfiles {
		if profile.IDProfile == profileID {
			profileInfo = profile
		}
	}

	if profileInfo != (models.Profile{}) {
		var newProfiles = models.AllProfile{}
		newProfiles = append(newProfiles, profileInfo)

		response := models.Response{
			Code:    http.StatusOK,
			Message: fmt.Sprintf("Información del usuario con ID: %v", profileID),
			Data:    newProfiles,
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		errObj := models.Error{
			Message: "Id no encontrado",
			Code:    400,
			Type:    true,
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(errObj)
	}
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	profileID, err := strconv.Atoi(vars["idProfile"])
	var updatedProfile models.Profile
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		err200 := models.Error{
			Message: "ID Invalido",
			Code:    400,
			Type:    true,
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err200)
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err200 := models.Error{
			Message: "Ingresar data valida para actualizar",
			Code:    400,
			Type:    true,
		}
		json.NewEncoder(w).Encode(err200)
	}
	json.Unmarshal(reqBody, &updatedProfile)

	for i, t := range allProfiles {
		if t.IDProfile == profileID {
			allProfiles = append(allProfiles[:i], allProfiles[i+1:]...)

			updatedProfile.IDProfile = t.IDProfile
			allProfiles = append(allProfiles, updatedProfile)

			// w.Header().Set("Content-Type", "application/json")
			// json.NewEncoder(w).Encode(updatedTask)
			response200 := models.Response{
				Code:    201,
				Message: fmt.Sprintf("El usuario con ID %v ha sido actualizado correctamente", profileID),
				Data:    allProfiles,
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(response200)
		}
	}

}

func DeleteProfile(w http.ResponseWriter, r *http.Request) {
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
		fmt.Fprintf(w, "Invalid User ID")
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
