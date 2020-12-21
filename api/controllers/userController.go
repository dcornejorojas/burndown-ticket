package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"ticket/api/models"
	"ticket/api/utils"
	u "ticket/api/utils"
	log "github.com/jeanphorn/log4go"
)

var allUsers = models.AllUsers{
	{
		Dni:      "123",
		Password: "PASS123",
		Name:     "Nombre",
		User:     "usuario1",
		LastName: "apellido1",
		Avatar:   "avatar/seleccionado",
		Rol:      "newRol",
	},
}

var HelloWorld = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HelloWorld")
	resp := map[string]interface{}{
		"Name": "Eve",
		"Age":  "6.0",
	}
	u.Respond(w, resp)
}

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {
	log.LOGGER("Routes").Info("POST - /user/ - Trying to create new user")
	var user models.User
	var errObj = models.Error{}
	errObj.NoError()

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errObj.HasError(true, http.StatusUnprocessableEntity, "Failed to process the body")
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		fmt.Fprintf(w, "Insert a Valid Task Data")
		return
	}

	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		errObj.HasError(true, http.StatusUnprocessableEntity, "Failed to get the new User")
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		fmt.Fprintf(w, "Failed to UnMarshal")
		return
	}

	user.Prepare()
	err = user.Validate("")
	if err != nil {
		errObj.HasError(true, http.StatusUnprocessableEntity, "Failed to process Entity")
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		fmt.Fprintf(w, "Insert a Valid Task Data")
		return
	}

	fmt.Println(user)
	newUser, err := user.SaveUser(server.DB)
	if err != nil {
		errObj.HasError(true, http.StatusUnprocessableEntity, "Failed to create the user")
		utils.ERROR(w, http.StatusUnprocessableEntity, err)
		fmt.Fprintf(w, "Insert a Valid Task Data")
		return
	}
	//allUsers = append(allUsers, user)

	utils.ResponseJSON(w, http.StatusCreated, "Usuario agregado exitosamente", newUser, errObj)

}

// func getTasks(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(tasks)
// }

// func getOneTask(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	taskID, err := strconv.Atoi(vars["id"])
// 	if err != nil {
// 		return
// 	}

// 	for _, task := range tasks {
// 		if task.ID == taskID {
// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(task)
// 		}
// 	}
// }

// func updateTask(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	taskID, err := strconv.Atoi(vars["id"])
// 	var updatedTask task

// 	if err != nil {
// 		fmt.Fprintf(w, "Invalid ID")
// 	}

// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "Please Enter Valid Data")
// 	}
// 	json.Unmarshal(reqBody, &updatedTask)

// 	for i, t := range tasks {
// 		if t.ID == taskID {
// 			tasks = append(tasks[:i], tasks[i+1:]...)

// 			updatedTask.ID = t.ID
// 			tasks = append(tasks, updatedTask)

// 			// w.Header().Set("Content-Type", "application/json")
// 			// json.NewEncoder(w).Encode(updatedTask)
// 			fmt.Fprintf(w, "The task with ID %v has been updated successfully", taskID)
// 		}
// 	}

// }

// func deleteTask(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	taskID, err := strconv.Atoi(vars["id"])

// 	if err != nil {
// 		fmt.Fprintf(w, "Invalid User ID")
// 		return
// 	}

// 	for i, t := range tasks {
// 		if t.ID == taskID {
// 			tasks = append(tasks[:i], tasks[i+1:]...)
// 			fmt.Fprintf(w, "The task with ID %v has been remove successfully", taskID)
// 		}
// 	}
// }

// var CreateContact = func(w http.ResponseWriter, r *http.Request) {

// 	user := r.Context().Value("user").(uint) //Grab the id of the user that send the request
// 	contact := &models.Contact{}

// 	err := json.NewDecoder(r.Body).Decode(contact)
// 	if err != nil {
// 		u.Respond(w, u.Message(false, "Error while decoding request body"))
// 		return
// 	}

// 	contact.UserId = user
// 	resp := contact.Create()
// 	u.Respond(w, resp)
// }
