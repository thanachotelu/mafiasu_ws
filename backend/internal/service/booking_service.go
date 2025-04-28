package service

import (
	"context"
	"mafiasu_ws/internal/interfaces"
	"mafiasu_ws/internal/models"
)

type bookingService struct {
	repo interfaces.BookingRepository
}

func NewBookingService(repo interfaces.BookingRepository) interfaces.BookingService {
	return &bookingService{repo: repo}
}

func (s *bookingService) GetBookingByID(ctx context.Context, id string) (models.Booking, error) {
	return s.repo.GetBookingByID(ctx, id)
}

func (s *bookingService) GetAllBooking(ctx context.Context) ([]models.Booking, error) {
	return s.repo.GetAllBooking(ctx)
}

func (s *bookingService) AddBooking(ctx context.Context, booking models.BookingRequest) (models.Booking, error) {
	return s.repo.AddBooking(ctx, booking)
}

func (s *bookingService) UpdateBooking(ctx context.Context, id string, booking models.Booking) (models.Booking, error) {
	return s.repo.UpdateBooking(ctx, id, booking)
}

func (s *bookingService) DeleteBooking(ctx context.Context, id string) (models.Booking, error) {
	return s.repo.DeleteBooking(ctx, id)
}
