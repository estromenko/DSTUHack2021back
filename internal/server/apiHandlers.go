package server

import (
	"dstuhack/internal/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (s *Server) GetAllSymbolStocks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		symbol := r.URL.Query().Get("symbol")
		if symbol == "" {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "provide query param with symbol",
			})
			w.WriteHeader(400)
			return
		}

		stocks, err := s.API().GetAllSymbolStocks(symbol)
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

		if price == 0 || quantity == 0 || req.Ticker == "" {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "price, quantity and ticker must be provided",
			})
			return
		}

		if req.Type == "buy" {
			if user.Balance < float32(price)*float32(int(quantity)) {
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"error": "not enough money",
				})
				return
			}

		} else {
			portfolio, err := s.User().GetPortfolio(user.ID)
			if err != nil {
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"error":   err.Error(),
					"message": "error getting portfolio",
				})
				return
			}

			val, ok := portfolio[req.Ticker]
			if !ok {
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"error": "user does not have and stocks of this type",
				})
				return
			}

			if quantity > int64(val) {
				w.WriteHeader(400)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"error": "not enough stocks to sell",
				})
				return
			}

			price *= -1
		}

		user.Balance -= float32(price) * float32(int(quantity))

		s.User().Repo().Update(user)

		if err := s.db.Operation().Create(&models.Operation{
			UserId:   user.ID,
			Type:     req.Type,
			Symbol:   req.Ticker,
			Price:    float32(price),
			Quantity: int(quantity),
		}); err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error":   err.Error(),
				"message": "error creating new operation",
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

func (s *Server) GetAllStokes() http.HandlerFunc {
	type stockEntity struct {
		Name      string  `json:"name"`
		Symbol    string  `json:"symbol"`
		Close     float32 `json:"close"`
		Diference float32 `json:"diference"`
	}
	type responseEntity struct {
		Balance float32 `json:"balance"`
		Stocks  []stockEntity
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

		pageStr := r.URL.Query().Get("page")
		page, _ := strconv.Atoi(pageStr)

		var resp responseEntity
		resp.Balance = user.Balance

		tickers, _ := s.API().GetAllTickers()

		for _, v := range tickers[page : page+10] { // Bad pagination
			stokes, _ := s.API().GetAllSymbolStocks(v.Symbol)
			for _, vv := range stokes {
				stock := stockEntity{
					Name:      v.Name,
					Symbol:    v.Symbol,
					Close:     vv.Close,
					Diference: vv.Close - vv.Open,
				}

				fmt.Println(stock)

				resp.Stocks = append(resp.Stocks, stock)
			}
		}

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(resp)
	}
}
