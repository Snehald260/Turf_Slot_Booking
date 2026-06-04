package model

import "time"

// Slot represents an available time window for a turf on a specific date.
type Slot struct {
	ID        string    `json:"id" db:"id"`
	TurfID    string    `json:"turf_id" db:"turf_id"`
	Date      string    `json:"date" db:"date"`           // YYYY-MM-DD
	StartTime string    `json:"start_time" db:"start_time"` // HH:MM
	EndTime   string    `json:"end_time" db:"end_time"`     // HH:MM
	Status    string    `json:"status" db:"status"`        // "available" | "reserved" | "booked"
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// CreateSlotRequest is the payload for creating slots.
type CreateSlotRequest struct {
	Date      string `json:"date" binding:"required"`       // YYYY-MM-DD
	StartTime string `json:"start_time" binding:"required"` // HH:MM
	EndTime   string `json:"end_time" binding:"required"`   // HH:MM
}
