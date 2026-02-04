package models

type User struct {
	ID           uint                       `json:"id"`
	Name         string                     `json:"name"`
	Username     string                     `json:"username"`
	Email        string                     `json:"email"`
	Password     string                     `json:"-"`
	Role         string                     `json:"role"`
	Carts        NullableSlice[Cart]        `json:"carts"`
	Orders       NullableSlice[Order]       `json:"orders"`
	Purchases    NullableSlice[Purchase]    `json:"purchases"`
	OrderReturns NullableSlice[OrderReturn] `json:"orderReturns"`
}

type UserWithToken struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}
