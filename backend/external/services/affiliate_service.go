package services

import (
	"context"

	"mafiasu_ws/external/interfaces"
	"mafiasu_ws/external/models"
)

type affiliateService struct {
	repo interfaces.AffiliateRepository
}

func NewAffiliateService(repo interfaces.AffiliateRepository) interfaces.AffiliateService {
	return &affiliateService{repo: repo}
}

// GetAllAffiliates
func (s *affiliateService) GetAllAffiliates(ctx context.Context) ([]models.Affiliator, error) {
	return s.repo.GetAllAffiliates(ctx)
}

// GetActiveCars
func (s *affiliateService) GetAvailableCars(ctx context.Context, affiliateID string) ([]models.Car, error) {
	return s.repo.GetAvailableCars(ctx, affiliateID)
}

// GetAffiliateBooking
func (s *affiliateService) GetAffiliateBookingDetail(ctx context.Context, affiliateID, bookingID string) (*models.BookingDetailResponse, error) {
	return s.repo.GetAffiliateBookingDetail(ctx, affiliateID, bookingID)
}

// TrackClick
func (s *affiliateService) TrackClick(ctx context.Context, affiliateID string, trackclick models.TrackClickRequest) (models.TrackClick, error) {
	return s.repo.TrackClick(ctx, affiliateID, trackclick)
}

// AddBookingbyAffiliator
func (s *affiliateService) CreateAffiliateBooking(ctx context.Context, affiliateID string, booking models.BookingRequest) (models.Booking, error) {
	return s.repo.CreateAffiliateBooking(ctx, affiliateID, booking)
}
