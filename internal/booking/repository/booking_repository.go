package repository

import "github.com/turf-booking-system/internal/booking/model"

// BookingRepository defines the contract for booking data access.
type BookingRepository interface {
	CreateBooking(booking *model.Booking) error
	GetBookingByID(id string) (*model.Booking, error)
	GetBookingsByUserID(userID string) ([]model.Booking, error)
	CancelBooking(id string) error
}

// SlotRepository defines the contract for slot data access.
type SlotRepository interface {
	CreateSlot(slot *model.Slot) error
	GetSlotByID(id string) (*model.Slot, error)
	GetAvailableSlots(turfID string, date string) ([]model.Slot, error)
	UpdateSlotStatus(id string, status string) error
}
