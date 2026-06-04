package repository

import "github.com/turf-booking-system/internal/auth/model"

// AuthRepository defines the contract for user data access.
type AuthRepository interface {
	CreateUser(user *model.User) error
	GetUserByEmail(email string) (*model.User, error)
	GetUserByID(id string) (*model.User, error)
}
