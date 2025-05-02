package interfaces

import (
	"context"
	"mafiasu_ws/internal/models"
)

type BookingService interface {
	GetBookingByID(ctx context.Context, id string) (models.Booking, error)
	GetAllBooking(ctx context.Context) ([]models.Booking, error)
	AddBooking(ctx context.Context, booking models.BookingRequest) (models.Booking, error)
	UpdateBooking(ctx context.Context, id string, booking models.Booking) (models.Booking, error)
	DeleteBooking(ctx context.Context, id string) (models.Booking, error)
}
