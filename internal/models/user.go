package models

type User struct {
	ID        int     `json:"id"`
	Email     string  `json:"email"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Password  string  `json:"password"`
	Balance   float32 `json:"balance"`
}
