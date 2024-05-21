package main

import (
	"example/databaseConnection"
	"example/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main(){

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	fmt.Println(port)

	databaseConnection.DBInstance()

	router := routes.CreateRouter()
	routes.CreateRoutes(router)
	http.ListenAndServe(port,router.Router)
	
}
