package models

import (
	"github.com/google/uuid"
)
type Engine struct {
	EngID uuid.UUID `json:"id"`
	Displacement int64 `json:"displacement"`
	NoOfCylinders int64 `json:"no_of_cylinders"`
	CarRange int64 `json:"car_range"`
}