package interfaces

// import (
// 	"context"
// 	"mafiasu_ws/internal/models"
// )

type BookingService interface {
	GetAllBooking()
	GetBookingByID()
	AddBooking()
	UpdateBooking()
	DeleteBooking()
}
