package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

/*Server struct of the Server
@DB : Gorm instance of the DB
@Router : Mux Router instance for the API
*/
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

//Initialize init the db and the routes of the API
func (server *Server) Initialize(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	if os.Getenv("DB_ENABLE") == "true" {
		fmt.Printf(`Trying to connect with the db: %s`, DBURL)
		server.DB, err = gorm.Open(DbDriver, DBURL)
		fmt.Println(err)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", DbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", DbDriver)
		}
	}

	server.Router = mux.NewRouter()
	//server.DB.AutoMigrate(&models.User{},&models.Profile{},&models.Ticket{},&models.Item{})

	server.InitRoutes()
}

//Run run the API
func (server *Server) Run(addr string) {
	fmt.Printf("\nServer listening on %s", addr)
	log.Fatal(http.ListenAndServe(":"+addr, server.Router))
}
