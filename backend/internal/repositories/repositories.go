// Package repositories
package repositories

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	QueryTimeoutDuration = time.Second * 5
	ErrNotFound          = errors.New("record not found")
)

type Repositories struct {
	User     UserRepository
	Category CategoryRepository
	Supplier SupplierRepository
}

func NewRepo(db *gorm.DB) *Repositories {
	return &Repositories{
		User:     &userRepository{db: db},
		Category: &categoryRepository{db: db},
		Supplier: &supplierRepository{db: db},
	}
}
