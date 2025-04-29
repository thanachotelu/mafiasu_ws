package repository

import (
	"context"
	"fmt"
	"mafiasu_ws/external/interfaces"
	"mafiasu_ws/external/models"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type affiliateRepository struct {
	db *pgxpool.Pool
}

func NewAffiliateRepository(db *pgxpool.Pool) interfaces.AffiliateRepository {
	return &affiliateRepository{db: db}
}

// GetAllAffiliates
func (r *affiliateRepository) GetAllAffiliates(ctx context.Context) ([]models.Affiliator, error) {
	rows, err := r.db.Query(ctx, `
		SELECT 
			affiliator_id,
			user_id,
			affiliate_code,
			referral_link,
			commission_rate,
			total_commission,
			balance,
			created_at,
			updated_at
		FROM affiliator
	`)

	if err != nil {
		return nil, fmt.Errorf("failed to get affiliators: %w", err)
	}
	defer rows.Close()

	var affiliates []models.Affiliator
	for rows.Next() {
		var a models.Affiliator
		err := rows.Scan(
			&a.AffiliatorID,
			&a.UserID,
			&a.AffiliateCode,
			&a.ReferralLink,
			&a.CommissionRate,
			&a.TotalCommission,
			&a.Balance,
			&a.CreatedAt,
			&a.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		affiliates = append(affiliates, a)
	}

	return affiliates, nil
}

func (r *affiliateRepository) GetAffiliatorByID(ctx context.Context, id string) (*models.Affiliator, error) {
	return nil, nil
}

func (r *affiliateRepository) GetBookingsByAffiliatorID(ctx context.Context, affiliatorID string) ([]models.Booking, error) {
	return nil, nil
}

// GetCars
func (r *affiliateRepository) GetAvailableCars(ctx context.Context, affiliateID string) ([]models.Car, error) {
	query := `
		SELECT car_id, brand, model, license_plate, carType, seat, doors, gearType, fuelType, rental_price_per_day, status car_status
		FROM cars
		WHERE status = 'active'
	`
	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cars []models.Car
	for rows.Next() {
		var car models.Car
		if err := rows.Scan(&car.CarID, &car.Brand, &car.Model, &car.LicensePlate, &car.CarType, &car.Seat, &car.Doors, &car.GearType, &car.FuelType, &car.RentalPricePerDay, &car.Status); err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}

	return cars, nil
}

// GetAffiliateBooking
func (r *affiliateRepository) GetAffiliateBookingDetail(ctx context.Context, affiliateID, bookingID string) (*models.BookingDetailResponse, error) {
	var booking models.BookingDetailResponse
	var pickupDate, returnDate time.Time
	var firstname, lastname string

	err := r.db.QueryRow(ctx, `
		SELECT
			b.book_id,
			b.total_price,
			b.status,
			b.pickup_date,
			b.return_date,
			c.car_id,
			c.model,
			c.rental_price_per_day,
			u.firstname,
			u.lastname,
			u.email
		FROM booking b
		JOIN cars c ON b.car_id = c.car_id
		JOIN users u ON b.user_id = u.user_id
		WHERE b.affiliator_id = $1 AND b.book_id = $2
	`, affiliateID, bookingID).Scan(
		&booking.BookingID,
		&booking.TotalPrice,
		&booking.Status,
		&pickupDate,
		&returnDate,
		&booking.Car.CarID,
		&booking.Car.Model,
		&booking.Car.RentalPricePerDay,
		&firstname,
		&lastname,
		&booking.Customer.Email,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get booking detail: %w", err)
	}

	booking.PickupDate = pickupDate.Format("2006-01-02")
	booking.ReturnDate = returnDate.Format("2006-01-02")
	booking.Customer.Name = firstname + " " + lastname

	return &booking, nil
}

// TrackClick
func (r *affiliateRepository) TrackClick(ctx context.Context, affiliateID string, trackclick models.TrackClickRequest) (models.TrackClick, error) {
	var track models.TrackClick
	err := r.db.QueryRow(ctx, `
		INSERT INTO trackclicks (car_id, affiliator_id, referral_link, created_at, updated_at) 
		VALUES ($1, $2, $3, NOW(), NOW())
		RETURNING session_id, car_id, affiliator_id, referral_link, created_at, updated_at
	`, trackclick.CarID, affiliateID, trackclick.ReferralLink).Scan( // <<< ตรงนี้ สลับ affiliateID มาไว้ที่ 2
		&track.SessionID,
		&track.CarID,
		&track.AffiliatorID,
		&track.ReferralLink,
		&track.CreatedAt,
		&track.UpdatedAt,
	)

	if err != nil {
		return models.TrackClick{}, fmt.Errorf("failed to add track-click: %w", err)
	}
	return track, nil
}

// CreateBookings
func (r *affiliateRepository) CreateAffiliateBooking(ctx context.Context, affiliateID string, booking models.BookingRequest) (models.Booking, error) {
	var newBooking models.Booking
	var pickupDate time.Time
	var returnDate time.Time

	err := r.db.QueryRow(ctx, `
	INSERT INTO booking (
		session_id, user_id, car_id, affiliator_id, total_price, pickup_date, return_date, created_at, updated_at
	) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())
	RETURNING 
		book_id, session_id, user_id, car_id, affiliator_id, total_price, pickup_date, return_date, status, created_at, updated_at
`,
		booking.SessionID,
		booking.UserID,
		booking.CarID,
		affiliateID,
		booking.TotalPrice,
		booking.PickupDate,
		booking.ReturnDate,
	).Scan(
		&newBooking.BookID,
		&newBooking.SessionID,
		&newBooking.UserID,
		&newBooking.CarID,
		&newBooking.AffiliateID,
		&newBooking.TotalPrice,
		&pickupDate,
		&returnDate,
		&newBooking.Status,
		&newBooking.CreatedAt,
		&newBooking.UpdatedAt,
	)

	if err != nil {
		return models.Booking{}, fmt.Errorf("failed to add booking: %w", err)
	}

	newBooking.PickupDate = pickupDate.Format("2006-01-02")
	newBooking.ReturnDate = returnDate.Format("2006-01-02")

	return newBooking, nil
}
