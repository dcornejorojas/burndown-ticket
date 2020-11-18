package main

import (
	"burndown-ticket/app"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// router := mux.NewRouter()
	// router.HandleFunc("/", controllers.HelloWorld).Methods("GET")
	// router.HandleFunc("/user/login", controllers.HelloWorld).Methods("POST")
	// router.HandleFunc("/user/logout", controllers.HelloWorld).Methods("POST")
	// router.HandleFunc("/profile", controllers.CreateProfile).Methods("POST")
	// router.HandleFunc("/profile/list", controllers.ListProfiles).Methods("POST")
	// router.HandleFunc("/profile/{idProfile}", controllers.UpdateProfile).Methods("PUT")
	// router.HandleFunc("/profile/{idProfile}", controllers.GetProfile).Methods("GET")
	// fs := http.FileServer(http.Dir("./assets"))
	// router.Handle("/", http.StripPrefix("/", fs)).Methods("GET")

	// e := godotenv.Load()
	// if e != nil {
	// 	fmt.Print(e)
	// }
	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "8000" //localhost
	// }
	// fmt.Println(port)

	// err := http.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8000/api
	// if err != nil {
	// 	fmt.Print(err)
	// }
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8000" //localhost
	}
	a := app.App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run(port)

}
