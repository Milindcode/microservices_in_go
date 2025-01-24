package main

import (
	"log"
	"net/http"

	"github.com/Milindcode/authentication-service/database"
	routes "github.com/Milindcode/authentication-service/router"
)

func main() {

	err := database.InitDB()
	if err != nil {
		log.Println(err)
	}

	mux := http.NewServeMux()
	routes.Routes(mux)

	log.Println("Server listening at port: 8001")
	http.ListenAndServe(":8001", mux)
}
