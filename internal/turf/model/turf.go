package model

import "time"

// Turf represents a box cricket ground.
type Turf struct {
	ID           string    `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	Location     string    `json:"location" db:"location"`
	City         string    `json:"city" db:"city"`
	OwnerID      string    `json:"owner_id" db:"owner_id"`
	SportType    string    `json:"sport_type" db:"sport_type"`
	PricePerHour float64   `json:"price_per_hour" db:"price_per_hour"`
	Amenities    []string  `json:"amenities" db:"amenities"`
	IsActive     bool      `json:"is_active" db:"is_active"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// CreateTurfRequest is the payload for creating a new turf.
type CreateTurfRequest struct {
	Name         string   `json:"name" binding:"required"`
	Location     string   `json:"location" binding:"required"`
	City         string   `json:"city" binding:"required"`
	SportType    string   `json:"sport_type" binding:"required,oneof=cricket football badminton"`
	PricePerHour float64  `json:"price_per_hour" binding:"required,gt=0"`
	Amenities    []string `json:"amenities,omitempty"`
}
