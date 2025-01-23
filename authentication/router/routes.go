package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Milindcode/authentication-service/database"
)

func Routes(mux *http.ServeMux) {

	mux.HandleFunc("POST /authenticate", func(w http.ResponseWriter, r *http.Request) {
		var payload struct {
			Email string `json:"email"`
			Password string `json:"password"`
		}

		var user database.User

		err := json.NewDecoder(r.Body).Decode(&payload).Error(); if err!= "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": true, 
				"message": "Invalid Request: " + err,
			})
			return 
		}

		er := database.DB_OBJ.DB.Where("email =?", payload.Email).First(&user).Error; if er != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": true, 
				"message": "Invalid Credentials: 1",
			})
			return 
		}

		if !ComparePassword(user.Password, payload.Password) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": true, 
				"message": "Invalid Credentials: 2",
			})
			return 
		}

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": false, 
			"message": "Authenticated",
			"user_data": user,
		}) 
	})


	mux.HandleFunc("POST /adduser", func(w http.ResponseWriter, r *http.Request) {
		var user database.User

		err := json.NewDecoder(r.Body).Decode(&user).Error(); if err!= "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": true, 
				"message": "Invalid Request: " + err,
			})
			return 
		}

		var er error 
		var hash string 
		hash, er = HashPassword(user.Password); if er != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": true, 
				"message": "Cannot add to database: " + err,
			})
			return 
		}

		user.Password = hash

		er = database.DB_OBJ.DB.Create(&user).Error; if er != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": true, 
				"message": "Cannot add to database: " + err,
			})
			return 
		}


		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": false, 
			"message": "Added to database",
			"user": user,
		})
	})

	// mux.HandleFunc("GET /getuser", func(w http.ResponseWriter, r *http.Request){})
}