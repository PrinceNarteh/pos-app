package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
	RoleGuest Role = "guest"
)

type User struct {
	Base
	FirstName    string         `gorm:"size:255;not null" json:"firstName"`
	LastName     string         `gorm:"size:255;not null" json:"lastName"`
	Username     string         `gorm:"size:100;not null" json:"username"`
	Email        string         `gorm:"size:255;not null;unique" json:"email"`
	Password     string         `gorm:"size:255;not null" json:"-"`
	Role         Role           `gorm:"size:6;not null" json:"role"`
	IsActive     bool           `gorm:"default:true" json:"isActive"`
	RefreshToken string         `gorm:"type:text" json:"-"`
	Carts        []Cart         `gorm:"foreignKey:UserID" json:"carts"`
	Orders       []Order        `gorm:"foreignKey:UserID" json:"orders"`
	Purchases    []Purchase     `gorm:"foreignKey:UserID" json:"purchases"`
	OrderReturns []OrderReturn  `gorm:"foreignKey:UserID" json:"orderReturns"`
	LastLoginAt  *time.Time     `json:"lastLoginAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deletedAt"`
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

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

type UserWithToken struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}
