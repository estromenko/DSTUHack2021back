package server

import (
	"dstuhack/internal/models"
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

func (s *Server) BuyStoke() http.HandlerFunc {
	type request struct {
		Ticker   string      `json:"ticker"`
		Price    json.Number `json:"price,omitempty"`
		Quantity json.Number `json:"quantity"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value("user").(*models.User)
		if !ok {
			w.WriteHeader(403)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "Not authorized",
			})
			return
		}

		var req request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		price, _ := req.Price.Float64()
		quantity, _ := req.Quantity.Int64()

		if user.Balance < float32(price)*float32(int(quantity)) {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "not enough money",
			})
			return
		}

		user.Balance -= float32(price) * float32(int(quantity))

		s.User().Repo().Update(user)

		if err := s.db.Operation().Create(&models.Operation{
			UserId:        user.ID,
			Type:          "stock",
			Name:          req.Ticker,
			PurchasePrice: float32(price),
			Amount:        int(quantity),
		}); err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": "stock successfully bought",
		})
		return
	}
}
