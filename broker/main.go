package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from api"))
	})

	fmt.Println("Server starting at localhost:8080")
	http.ListenAndServe(":8080", mux)
}
