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

type resStocks struct {
	Data []*Stock `json:"data"`
}

func NewAPIService(accessKey string) *APIService {
	return &APIService{
		location:  "https://api.marketstack.com/v1",
		accessKey: accessKey,
		client:    &http.Client{},
	}
}

func (s *APIService) GetAllSymbolStocks(symbol string) ([]*Stock, error) {
	path := s.location + "/eod?access_key=" + s.accessKey + "&symbols=" + symbol
	req, err := http.NewRequest("GET", path, nil)

	if err != nil {
		return nil, err
	}

	var stocks resStocks

	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&stocks); err != nil {
		return nil, err
	}

	return stocks.Data, nil
}
