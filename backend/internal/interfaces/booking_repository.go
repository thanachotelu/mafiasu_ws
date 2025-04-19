package interfaces

// import (
// 	"context"
// 	"mafiasu_ws/internal/models"
// )

type BookingRepository interface {
	GetAllBooking()
	GetBookingByID()
	AddBooking()
	UpdateBooking()
	DeleteBooking()
}
