package models

import (
	"time"
)

type Operation struct {
	ID       int       `json:"id"`
	UserId   int       `json:"user_id"`
	Type     string    `json:"type"`
	Symbol   string    `json:"symbol"`
	Name     string    `json:"name"`
	Date     time.Time `json:"date"`
	Price    float32   `json:"price"`
	Quantity int       `json:"quantity"`
}
