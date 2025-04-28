package interfaces

import (
	"context"

	"mafiasu_ws/external/models"
)

type AffiliateRepository interface {
	GetAllAffiliates(ctx context.Context) ([]models.Affiliator, error)
	GetBookingsByAffiliatorID(ctx context.Context, affiliatorID string) ([]models.Booking, error)
	GetAvailableCars(ctx context.Context, affiliateID string) ([]models.Car, error)
	GetAffiliateBookingDetail(ctx context.Context, affiliateID, bookingID string) (*models.BookingDetailResponse, error)
	TrackClick(ctx context.Context, affiliateID string, trackclick models.TrackClickRequest) (models.TrackClick, error)
	CreateAffiliateBooking(ctx context.Context, affiliateID string, booking models.BookingRequest) (models.Booking, error)
}
