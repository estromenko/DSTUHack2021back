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

func (s *Server) BuyOrSellStoke() http.HandlerFunc {
	type request struct {
		Type     string      `json:"type"`
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

		if req.Type != "buy" && req.Type != "sell" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "invalid operation",
			})
			return
		}

		price, _ := req.Price.Float64()
		quantity, _ := req.Quantity.Int64()

		if req.Type == "buy" {
			if user.Balance < float32(price)*float32(int(quantity)) {
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"error": "not enough money",
				})
				return
			}

		} else {
			price *= -1
		}

		user.Balance -= float32(price) * float32(int(quantity))

		s.User().Repo().Update(user)

		if err := s.db.Operation().Create(&models.Operation{
			UserId: user.ID,
			Type:   req.Type,
			Name:   req.Ticker,
			Price:  float32(price),
			Quantity: int(quantity),
		}); err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": "operation done successfully",
			"balance": user.Balance,
		})
		return
	}
}

func (s *Server) GetAllTickers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tickers, err := s.API().GetAllTickers()
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(tickers)
	}
}
