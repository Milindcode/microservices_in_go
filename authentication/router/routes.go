package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Milindcode/authentication-service/database"
)

func Routes(mux *http.ServeMux) {

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":   false,
			"message": "Hi from authentication",
		})
	})

	mux.HandleFunc("POST /authenticate", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		var payload struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		var user database.User

		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error":   true,
				"message": "Invalid Request",
			})
			return
		}

		er := database.DB_OBJ.DB.Where("email =?", payload.Email).First(&user).Error
		if er != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error":   true,
				"message": "Invalid Credentials: 1",
			})
			return
		}

		if !ComparePassword(user.Password, payload.Password) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error":   true,
				"message": "Invalid Credentials: 2",
			})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":     false,
			"message":   "Authenticated",
			"user_data": user,
		})
	})

	mux.HandleFunc("POST /adduser", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		var user database.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {

			log.Println("HELLO 1")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error":   true,
				"message": "Invalid Request",
			})
			return
		}

		var er error
		var hash string
		log.Println("HELLO 2")
		hash, er = HashPassword(user.Password)
		if er != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error":   true,
				"message": "Cannot add to database",
			})
			return
		}

		log.Println("HELLO 3")

		user.Password = hash

		er = database.DB_OBJ.DB.Create(&user).Error
		if er != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error":   true,
				"message": "Cannot add to database",
			})
			return
		}

		log.Println("HELLO 4")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":   false,
			"message": "Added to database",
			"user":    user,
		})
	})

	// mux.HandleFunc("GET /getuser", func(w http.ResponseWriter, r *http.Request){})
}
