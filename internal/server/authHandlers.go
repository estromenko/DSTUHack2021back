package server

import (
	"database/sql"
	"dstuhack/internal/models"
	"encoding/json"
	"net/http"
)

func (s *Server) RegisterUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		user.Balance = 0

		token, err := s.User().Create(&user)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		w.WriteHeader(201)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"token": token,
		})
	}
}

func (s *Server) LoginUser() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		var data request
		json.NewDecoder(r.Body).Decode(&data)

		user, err := s.User().Repo().FindByEmail(data.Email)
		if err != nil {
			w.WriteHeader(400)
			if err == sql.ErrNoRows {
				json.NewEncoder(w).Encode(map[string]interface{}{
					"error": "No user with such email: " + data.Email,
				})
				return
			}
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		if !s.User().ComparePasswords(user, data.Password) {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "Wrong email or password",
			})
			return
		}

		token, err := s.User().GenerateToken(user)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"token": token,
		})
	}
}
