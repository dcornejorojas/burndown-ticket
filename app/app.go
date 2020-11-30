package app

import (
	"ticket/routers"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = routers.InitRoutes()
}

func (a *App) Run(addr string) {
	fmt.Println(addr)
	err := http.ListenAndServe(addr, a.Router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
