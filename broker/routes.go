package main

import (
	"encoding/json"
	"net/http"
)

type Auth_payload struct {
	Email string `json:"email"1`
	Password string `json:"password"`
}

func Routes(mux *http.ServeMux) {

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": false, 
			"message": "hello from broker",
		})
	})

	// mux.HandleFunc("/authenticate", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 	w.Header().Set("Content-Type", "application/json")

	// 	url := "http://authentication-service/authenticate"

	// 	var payload Auth_payload
	// 	err := json.NewDecoder(r.Body).Decode(&payload)
	// 	if err!=nil {
	// 		w.WriteHeader(http.StatusBadRequest)
	// 		json.NewEncoder(w).Encode(map[string]interface{}{
	// 			"error": true, 
	// 			"message": "Bad Request",
	// 		})
	// 	}



	// 	w.WriteHeader(http.StatusOK)
	// 	json.NewEncoder(w).Encode(map[string]interface{}{
	// 		"error": false, 
	// 		"message": "hello from broker",
	// 	})
	})
}
