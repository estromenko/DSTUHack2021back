package server

import (
	"dstuhack/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
)

func (s *Server) GetUserInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value("user").(*models.User)
		if !ok {
			w.WriteHeader(403)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "Not authorized",
			})
			return
		}

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(user)
	}
}

func (s *Server) GetAllUserOperations() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value("user").(*models.User)
		if !ok {
			w.WriteHeader(403)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "Not authorized",
			})
			return
		}

		operations, err := s.db.Operation().GetAllByUserId(user.ID)
		if err != nil {
			w.WriteHeader(403)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(operations)
	}
}

func (s *Server) GetUserPortfolio() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value("user").(*models.User)
		if !ok {
			w.WriteHeader(403)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "Not authorized",
			})
			return
		}

		portfolio, err := s.User().GetPortfolio(user.ID)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(portfolio)
	}
}

func (s *Server) UpdateUserBalance() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value("user").(*models.User)
		if !ok {
			w.WriteHeader(403)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "Not authorized",
			})
			return
		}

		strBalance := struct {
			Balance string `json:"balance"`
		}{}

		json.NewDecoder(r.Body).Decode(&strBalance)

		balance, err := strconv.Atoi(strBalance.Balance)

		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		user.Balance = float32(balance)

		s.db.User().Update(user)

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"balance": user.Balance,
		})
	}
}
