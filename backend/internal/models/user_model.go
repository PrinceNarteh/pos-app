package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Base
	FirstName    string         `gorm:"size:255;not null" json:"firstName"`
	LastName     string         `gorm:"size:255;not null" json:"lastName"`
	Username     string         `gorm:"size:100;not null" json:"username"`
	Email        string         `gorm:"size:255;not null;unique" json:"email"`
	Password     string         `gorm:"size:255;not null" json:"-"`
	Role         string         `gorm:"size:6;not null" json:"role"`
	IsActive     bool           `gorm:"default:true" json:"isActive"`
	RefreshToken string         `gorm:"type:text" json:"-"`
	Carts        []Cart         `gorm:"foreignKey:UserID" json:"carts"`
	Orders       []Order        `gorm:"foreignKey:UserID" json:"orders"`
	Purchases    []Purchase     `gorm:"foreignKey:UserID" json:"purchases"`
	OrderReturns []OrderReturn  `gorm:"foreignKey:UserID" json:"orderReturns"`
	LastLoginAt  *time.Time     `json:"lastLoginAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type UserProfileResponse struct {
	Base
	FirstName string         `json:"firstName"`
	LastName  string         `json:"lastName"`
	Username  string         `json:"username"`
	Email     string         `json:"email"`
	Password  string         `json:"-"`
	Role      string         `json:"role"`
	IsActive  bool           `json:"isActive"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

type UserDetailResponse struct {
	UserProfileResponse
	Carts        []Cart        `json:"carts"`
	Orders       []Order       `json:"orders"`
	Purchases    []Purchase    `json:"purchases"`
	OrderReturns []OrderReturn `json:"orderReturns"`
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

func (u *User) ToUserProfileResponse() UserProfileResponse {
	return UserProfileResponse{
		Base: Base{
			ID:        u.Base.ID,
			CreatedAt: u.Base.CreatedAt,
			UpdatedAt: u.Base.UpdatedAt,
		},
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Username:  u.Username,
		Email:     u.Email,
		Role:      u.Role,
		IsActive:  u.IsActive,
		DeletedAt: u.DeletedAt,
	}
}

func (u *User) UserDetailResponse() UserDetailResponse {
	return UserDetailResponse{
		UserProfileResponse: UserProfileResponse{
			Base: Base{
				ID:        u.Base.ID,
				CreatedAt: u.Base.CreatedAt,
				UpdatedAt: u.Base.UpdatedAt,
			},
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Username:  u.Username,
			Email:     u.Email,
			Role:      u.Role,
			IsActive:  u.IsActive,
			DeletedAt: u.DeletedAt,
		},
		Carts:        u.Carts,
		OrderReturns: u.OrderReturns,
		Orders:       u.Orders,
		Purchases:    u.Purchases,
	}
}
