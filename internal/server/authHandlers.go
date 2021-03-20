package server

import (
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
