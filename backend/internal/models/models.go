package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type NullableSlice[T any] struct {
	Slice []T
}

// MarshalJSON customizes the JSON marshaling for NullableSlice.
func (ns NullableSlice[T]) MarshalJSON() ([]byte, error) {
	if ns.Slice == nil {
		return json.Marshal([]T{})
	}
	return json.Marshal(ns.Slice)
}
