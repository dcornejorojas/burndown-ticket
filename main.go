package main

import (
	"ticket/api"
)

func main() {
	api.Run()
	// e := godotenv.Load()
	// if e != nil {
	// 	fmt.Print(e)
	// }
	// port := os.Getenv("PORT")
	// // if port == "" {
	// // 	port = ":8000" //localhost
	// // }
	// a := app.App{}
	// a.Initialize(
	// 	os.Getenv("APP_DB_USERNAME"),
	// 	os.Getenv("APP_DB_PASSWORD"),
	// 	os.Getenv("APP_DB_NAME"))

}
