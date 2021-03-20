package server

import (
	"dstuhack/internal/models"
	"dstuhack/internal/db"
	"net/http"
	"encoding/json"
)

/*
func (s *Server) RegisterUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
*/

func (s *Server) GetAllOperationsByUserId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user")

		json.NewEncoder(w).Encode(map[string]interface{}{
			
		})
	}
}

func (s* Server) GetAllOperationsByUserIdAndType() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user")

		json.NewEncoder(w).Encode(map[string]interface{}{
			"asdasd" : "sad"
		})
	}
}

func (s *Server) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user")

		json.NewEncoder(w).Encode(map[string]interface{}{
			
		})
	}
}

func (s *Server) ChangeOperation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user")

		json.NewEncoder(w).Encode(map[string]interface{}{
			
		})
	}
}