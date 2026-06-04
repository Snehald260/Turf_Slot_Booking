package repository

import "github.com/turf-booking-system/internal/turf/model"

// TurfRepository defines the contract for turf data access.
type TurfRepository interface {
	Create(turf *model.Turf) error
	GetByID(id string) (*model.Turf, error)
	List(city string, sportType string) ([]model.Turf, error)
	Update(turf *model.Turf) error
	Delete(id string) error
}
