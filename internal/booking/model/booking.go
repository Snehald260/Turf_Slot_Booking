package model

import "time"

// Booking links a user to a slot.
type Booking struct {
	ID            string     `json:"id" db:"id"`
	SlotID        string     `json:"slot_id" db:"slot_id"`
	UserID        string     `json:"user_id" db:"user_id"`
	Status        string     `json:"status" db:"status"`                 // "pending" | "confirmed" | "cancelled"
	PaymentStatus string     `json:"payment_status" db:"payment_status"` // "unpaid" | "paid" | "refunded"
	Amount        float64    `json:"amount" db:"amount"`
	BookedAt      time.Time  `json:"booked_at" db:"booked_at"`
	CancelledAt   *time.Time `json:"cancelled_at,omitempty" db:"cancelled_at"`
}

// CreateBookingRequest is the payload for creating a booking.
type CreateBookingRequest struct {
	SlotID string `json:"slot_id" binding:"required"`
}
