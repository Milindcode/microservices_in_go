package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Auth_payload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Routes(mux *http.ServeMux) {

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":   false,
			"message": "hello from broker",
		})
	})

	mux.HandleFunc("/authenticate", func(w http.ResponseWriter, r *http.Request) {
		// Add CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		url := "http://authentication-service:8001/authenticate"

		var payload Auth_payload
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error":   true,
				"message": "Bad Request",
			})
			return
		}

		log.Println(payload)

		jsonData, _ := json.Marshal(payload)
		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))

		req.Header.Set("Content-Type", "application/json")

		// Create an HTTP client and send the request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("HTTP Client Error:", err) // Log the err
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error":   true,
				"message": "Internal Server Problem",
			})
			return
		}
		defer resp.Body.Close()

		var authResponse map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&authResponse)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error":   true,
				"message": "Internal Server Error",
			})
			return
		}

		log.Println(authResponse)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":   authResponse["error"],
			"message": authResponse["message"],
			"data":    authResponse["user_data"],
		})
	})
}
