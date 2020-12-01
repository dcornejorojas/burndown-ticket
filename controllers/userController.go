package controllers

import (
	"ticket/models"
	u "ticket/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var allUsers = models.AllUsers{
	{
		UserDni:  int64(2),
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

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var user models.User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task Data")
	}

	json.Unmarshal(reqBody, &user)
	user.UserDni = int64(len(allUsers) + 1)
	allUsers = append(allUsers, user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(allUsers)

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
