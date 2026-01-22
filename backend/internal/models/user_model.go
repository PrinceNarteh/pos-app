package models

type User struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	Username     string        `json:"username"`
	Email        string        `json:"email"`
	Password     string        `json:"password"`
	Role         string        `json:"role"`
	Carts        []Cart        `json:"carts"`
	Orders       []Order       `json:"orders"`
	Purchases    []Purchase    `json:"purchases"`
	OrderReturns []OrderReturn `json:"orderReturns"`
}
