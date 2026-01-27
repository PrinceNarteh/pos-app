package models

import "encoding/json"

type NullableSlice[T any] struct {
	Slice []T
}

func (ns NullableSlice[T]) MarshalJSON() ([]byte, error) {
	if ns.Slice == nil {
		return json.Marshal([]T{})
	}
	return json.Marshal(ns.Slice)
}
