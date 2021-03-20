package services

import (
	"encoding/json"
	"net/http"
)

type Stock struct {
	Open   float32 `json:"open"`
	High   float32 `json:"high"`
	Low    float32 `json:"low"`
	Close  float32 `json:"close"`
	Volume float32 `json:"volume"`
	Date   string  `json:"date"`
	Symbol string  `json:"symbol"`
}

type APIService struct {
	location  string
	accessKey string
	client    *http.Client
}

func (s *APIService) GetAllSymbolStocks() ([]*Stock, error) {
	req, err := http.NewRequest("GET", s.location+"?access_key="+s.accessKey, nil)
	if err != nil {
		return nil, err
	}

	var stocks []*Stock

	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&stocks); err != nil {
		return nil, err
	}
	return stocks, nil
}
