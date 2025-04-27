package repository

import (
	"context"
	"fmt"

	"mafiasu_ws/internal/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type BookingRepository interface {
	GetBookingByID(ctx context.Context, id string) (models.Booking, error)
	GetAllBooking(ctx context.Context) ([]models.Booking, error)
	AddBooking(ctx context.Context, booking models.BookingRequest) (models.Booking, error)
	UpdateBooking(ctx context.Context, id string, booking models.Booking) (models.Booking, error)
	DeleteBooking(ctx context.Context, id string) (models.Booking, error)
}

type bookingRepository struct {
	db *pgxpool.Pool
}

func NewBookingRepository(db *pgxpool.Pool) BookingRepository {
	return &bookingRepository{db: db}
}

func (r *bookingRepository) GetBookingByID(ctx context.Context, id string) (models.Booking, error) {
	var booking models.Booking
	err := r.db.QueryRow(ctx, `
		SELECT book_id, user_id, car_id, affiliator_id, total_price, created_at, updated_at
		FROM booking
		WHERE book_id = $1
	`, id).Scan(
		&booking.BookID,
		&booking.UserID,
		&booking.CarID,
		&booking.AffiliatorID,
		&booking.TotalPrice,
		&booking.CreatedAt,
		&booking.UpdatedAt,
	)

	if err != nil {
		return models.Booking{}, fmt.Errorf("failed to get booking: %w", err)
	}
	return booking, nil
}

func (r *bookingRepository) GetAllBooking(ctx context.Context) ([]models.Booking, error) {
	rows, err := r.db.Query(ctx, `
		SELECT book_id, user_id, car_id, affiliator_id, total_price, created_at, updated_at
		FROM booking
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to get bookings: %w", err)
	}
	defer rows.Close()

	var bookings []models.Booking
	for rows.Next() {
		var booking models.Booking
		if err := rows.Scan(
			&booking.BookID,
			&booking.UserID,
			&booking.CarID,
			&booking.AffiliatorID,
			&booking.TotalPrice,
			&booking.CreatedAt,
			&booking.UpdatedAt,
		); err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

func (r *bookingRepository) AddBooking(ctx context.Context, booking models.BookingRequest) (models.Booking, error) {
	var newBooking models.Booking
	err := r.db.QueryRow(ctx, `
		INSERT INTO booking (user_id, car_id, affiliator_id, total_price, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		RETURNING book_id, user_id, car_id, affiliator_id, total_price, created_at, updated_at
	`,
		booking.UserID,
		booking.CarID,
		booking.AffiliatorID,
		booking.TotalPrice,
	).Scan(
		&newBooking.BookID,
		&newBooking.UserID,
		&newBooking.CarID,
		&newBooking.AffiliatorID,
		&newBooking.TotalPrice,
		&newBooking.CreatedAt,
		&newBooking.UpdatedAt,
	)

	if err != nil {
		return models.Booking{}, fmt.Errorf("failed to add booking: %w", err)
	}
	return newBooking, nil
}

func (r *bookingRepository) UpdateBooking(ctx context.Context, id string, booking models.Booking) (models.Booking, error) {
	var updatedBooking models.Booking
	err := r.db.QueryRow(ctx, `
		UPDATE booking
		SET user_id=$1, car_id=$2, affiliator_id=$3, total_price=$4, updated_at=NOW()
		WHERE book_id=$5
		RETURNING book_id, user_id, car_id, affiliator_id, total_price, created_at, updated_at
	`,
		booking.UserID,
		booking.CarID,
		booking.AffiliatorID,
		booking.TotalPrice,
		id,
	).Scan(
		&updatedBooking.BookID,
		&updatedBooking.UserID,
		&updatedBooking.CarID,
		&updatedBooking.AffiliatorID,
		&updatedBooking.TotalPrice,
		&updatedBooking.CreatedAt,
		&updatedBooking.UpdatedAt,
	)

	if err != nil {
		return models.Booking{}, fmt.Errorf("failed to update booking: %w", err)
	}
	return updatedBooking, nil
}

func (r *bookingRepository) DeleteBooking(ctx context.Context, id string) (models.Booking, error) {
	var deletedBooking models.Booking
	err := r.db.QueryRow(ctx, `
		DELETE FROM booking
		WHERE book_id = $1
		RETURNING book_id, user_id, car_id, affiliator_id, total_price, created_at, updated_at
	`, id).Scan(
		&deletedBooking.BookID,
		&deletedBooking.UserID,
		&deletedBooking.CarID,
		&deletedBooking.AffiliatorID,
		&deletedBooking.TotalPrice,
		&deletedBooking.CreatedAt,
		&deletedBooking.UpdatedAt,
	)

	if err != nil {
		return models.Booking{}, fmt.Errorf("failed to delete booking: %w", err)
	}
	return deletedBooking, nil
}
