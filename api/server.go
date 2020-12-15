package api

import (
	"fmt"
	"log"
	"os"
	"ticket/api/controllers"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var server = controllers.Server{}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("Cannot load .env file")
	}
}

// func (s *Server) Initialize(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
// 	var err error
// 	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
// 	if os.Getenv("DB_ENABLE") == "true" {
// 		s.DB, err = gorm.Open(DbDriver, DBURL)
// 		if err != nil {
// 			fmt.Printf("Cannot connect to %s database", DbDriver)
// 			log.Fatal("This is the error:", err)
// 		} else {
// 			fmt.Printf("We are connected to the %s database", DbDriver)
// 		}
// 	}

// 	s.Router = mux.NewRouter()
// 	s.InitRoutes(s)
// }

func Run() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server.Initialize(
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"))

	server.Run(port)
}
