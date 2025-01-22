package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	Routes(mux)	

	fmt.Println("Server starting at localhost:8080")
	http.ListenAndServe(":8080", mux)
}
