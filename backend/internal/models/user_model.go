package models

import "gorm.io/gorm"

type User struct {
	Base
	FirstName    string         `gorm:"size:255;not null" json:"firstName"`
	LastName     string         `gorm:"size:255;not null" json:"lastName"`
	Username     string         `gorm:"size:100;not null" json:"username"`
	Email        string         `gorm:"size:255;not null;unique" json:"email"`
	Password     string         `gorm:"size:255;not null" json:"-"`
	Role         string         `gorm:"size:6;not null" json:"role"`
	Carts        []Cart         `gorm:"foreignKey:UserID" json:"carts"`
	Orders       []Order        `gorm:"foreignKey:UserID" json:"orders"`
	Purchases    []Purchase     `gorm:"foreignKey:UserID" json:"purchases"`
	OrderReturns []OrderReturn  `gorm:"foreignKey:UserID" json:"orderReturns"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	if u.Carts == nil {
		u.Carts = []Cart{}
	}
	if u.Orders == nil {
		u.Orders = []Order{}
	}
	if u.Purchases == nil {
		u.Purchases = []Purchase{}
	}
	if u.OrderReturns == nil {
		u.OrderReturns = []OrderReturn{}
	}
	return
}

type UserWithToken struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}
