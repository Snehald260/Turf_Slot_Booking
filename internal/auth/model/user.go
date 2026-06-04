package model

import "time"

// User represents a player, turf owner, or admin.
type User struct {
	ID           string    `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	Email        string    `json:"email" db:"email"`
	Phone        string    `json:"phone,omitempty" db:"phone"`
	PasswordHash string    `json:"-" db:"password_hash"` // never expose in JSON
	Role         string    `json:"role" db:"role"`       // "player" | "owner" | "admin"
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

// SignupRequest is the payload for user registration.
type SignupRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone,omitempty"`
	Password string `json:"password" binding:"required,min=8"`
	Role     string `json:"role" binding:"required,oneof=player owner"`
}

// LoginRequest is the payload for user login.
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// LoginResponse is returned on successful login.
type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
