package models

type Operation struct {
	ID       int     `json:"id"`
	UserId   int     `json:"user_id"`
	Type     string  `json:"type"`
	Symbol   string  `json:"symbol"`
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"`
}
