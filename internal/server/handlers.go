package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"asd": "asd",
		})
	}
}
