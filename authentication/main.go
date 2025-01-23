package main

import (
	"log"
	"net/http"

	"github.com/Milindcode/authentication-service/database"
	routes "github.com/Milindcode/authentication-service/router"
)

func main(){

	database.InitDB()

	mux := http.NewServeMux()
	routes.Routes(mux)

	log.Println("Server listening at port: 8080")
	http.ListenAndServe(":8080", mux)
}