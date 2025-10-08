package model

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Ignored   string  `json:"-"` // This is how you ignore a field when this is serialized to JSON
}
