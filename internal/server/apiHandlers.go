package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) GetAllSymbolStocks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("symbol")
		if query == "" {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "provide query param with symbol",
			})
			w.WriteHeader(400)
			return
		}

		stocks, err := s.API().GetAllSymbolStocks(query)
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(stocks)
	}
}
