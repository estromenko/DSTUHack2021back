package models

type Operation struct {
	ID       int     `json:"id"`
	UserId   int     `json:"user_id"`
	Type     string  `json:"type"`
	Name     string  `json:"name"`
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"`
}
