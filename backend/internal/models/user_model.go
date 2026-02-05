package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName    string                     `gorm:"size:255;not null" json:"firstName"`
	LastName     string                     `gorm:"size:255;not null" json:"lastName"`
	Username     string                     `gorm:"size:100;not null" json:"username"`
	Email        string                     `gorm:"size:255;not null;unique" json:"email"`
	Password     string                     `gorm:"size:255;not null" json:"-"`
	Role         string                     `gorm:"size:6;not null" json:"role"`
	Carts        NullableSlice[Cart]        `gorm:"foreignKey:UserID" json:"carts"`
	Orders       NullableSlice[Order]       `gorm:"foreignKey:UserID" json:"orders"`
	Purchases    NullableSlice[Purchase]    `gorm:"foreignKey:UserID" json:"purchases"`
	OrderReturns NullableSlice[OrderReturn] `gorm:"foreignKey:UserID" json:"orderReturns"`
}

type UserWithToken struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}
