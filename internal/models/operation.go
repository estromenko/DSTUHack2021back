package models

type Operation struct {
	ID int `json:"id"`
	UserId int 	`json:"user_id"`
	Type string `json:"type"`
	Name string `json:"name"`
	PurchasePrice float32 `json:"purchase_price"`
	Amount int `json:"amount"`
}