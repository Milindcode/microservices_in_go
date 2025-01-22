package main

import (
	"encoding/json"
	"net/http"
)

func Routes(mux *http.ServeMux) {

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": false, 
			"message": "hello from broker",
		})
	})
}
