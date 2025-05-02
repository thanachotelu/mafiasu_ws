package repository

import (
	"context"
	"fmt"
	"time"

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
		SELECT book_id, session_id, user_id, car_id, affiliator_id, total_price, pickup_date, return_date, status, created_at, updated_at
		FROM booking
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to get bookings: %w", err)
	}
	defer rows.Close()

	var bookings []models.Booking
	for rows.Next() {
		var booking models.Booking
		var pickupDate, returnDate time.Time

		err := rows.Scan(
			&booking.BookID,
			&booking.SessionID,
			&booking.UserID,
			&booking.CarID,
			&booking.AffiliatorID,
			&booking.TotalPrice,
			&pickupDate,
			&returnDate,
			&booking.Status,
			&booking.CreatedAt,
			&booking.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan booking: %w", err)
		}

		// แปลง Date -> String
		booking.PickupDate = pickupDate.Format("2006-01-02")
		booking.ReturnDate = returnDate.Format("2006-01-02")

		bookings = append(bookings, booking)
	}

	return bookings, nil
}

func (r *bookingRepository) AddBooking(ctx context.Context, booking models.BookingRequest) (models.Booking, error) {
	var newBooking models.Booking
	var pickupDate, returnDate time.Time

	err := r.db.QueryRow(ctx, `
		INSERT INTO booking (
			user_id, car_id, total_price, pickup_date, return_date, created_at, updated_at
		) 
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
		RETURNING 
			book_id, session_id, user_id, car_id, affiliator_id, total_price, pickup_date, return_date, status, created_at, updated_at
	`,
		booking.UserID,
		booking.CarID,
		booking.TotalPrice,
		booking.PickupDate,
		booking.ReturnDate,
	).Scan(
		&newBooking.BookID,
		&newBooking.SessionID,
		&newBooking.UserID,
		&newBooking.CarID,
		&newBooking.AffiliatorID,
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

	// แปลง Date -> String
	newBooking.PickupDate = pickupDate.Format("2006-01-02")
	newBooking.ReturnDate = returnDate.Format("2006-01-02")

	return newBooking, nil
}

func (r *bookingRepository) UpdateBooking(ctx context.Context, id string, booking models.Booking) (models.Booking, error) {
	var existing models.Booking
	var pickupDate, returnDate time.Time

	// 1. อ่านข้อมูลเดิมก่อน
	err := r.db.QueryRow(ctx, `
		SELECT book_id, session_id, user_id, car_id, affiliator_id, total_price, pickup_date, return_date, status, created_at, updated_at
		FROM booking
		WHERE book_id = $1
	`, id).Scan(
		&existing.BookID,
		&existing.SessionID,
		&existing.UserID,
		&existing.CarID,
		&existing.AffiliatorID,
		&existing.TotalPrice,
		&pickupDate,
		&returnDate,
		&existing.Status,
		&existing.CreatedAt,
		&existing.UpdatedAt,
	)
	if err != nil {
		return models.Booking{}, fmt.Errorf("booking not found: %w", err)
	}
	existing.PickupDate = pickupDate.Format("2006-01-02")
	existing.ReturnDate = returnDate.Format("2006-01-02")

	// 2. คงค่าเดิมไว้ถ้าไม่ได้ส่งค่ามาใหม่
	userID := booking.UserID
	if userID == "" {
		userID = existing.UserID
	}

	carID := booking.CarID
	if carID == "" {
		carID = existing.CarID
	}

	affiliatorID := booking.AffiliatorID
	if affiliatorID == nil {
		affiliatorID = existing.AffiliatorID
	}

	totalPrice := booking.TotalPrice
	if totalPrice == 0 {
		totalPrice = existing.TotalPrice
	}

	status := booking.Status
	if status == "" {
		status = existing.Status
	}

	// 3. อัปเดตด้วยค่าที่เลือกแล้ว
	var updatedBooking models.Booking
	err = r.db.QueryRow(ctx, `
		UPDATE booking
		SET user_id=$1, car_id=$2, affiliator_id=$3, total_price=$4, status=$5, updated_at=NOW()
		WHERE book_id=$6
		RETURNING 
			book_id, session_id, user_id, car_id, affiliator_id, total_price, pickup_date, return_date, status, created_at, updated_at
	`,
		userID,
		carID,
		affiliatorID,
		totalPrice,
		status,
		id,
	).Scan(
		&updatedBooking.BookID,
		&updatedBooking.SessionID,
		&updatedBooking.UserID,
		&updatedBooking.CarID,
		&updatedBooking.AffiliatorID,
		&updatedBooking.TotalPrice,
		&pickupDate,
		&returnDate,
		&updatedBooking.Status,
		&updatedBooking.CreatedAt,
		&updatedBooking.UpdatedAt,
	)
	if err != nil {
		return models.Booking{}, fmt.Errorf("failed to update booking: %w", err)
	}

	updatedBooking.PickupDate = pickupDate.Format("2006-01-02")
	updatedBooking.ReturnDate = returnDate.Format("2006-01-02")

	return updatedBooking, nil
}

func (r *bookingRepository) DeleteBooking(ctx context.Context, id string) (models.Booking, error) {
	var deletedBooking models.Booking
	var pickupDate, returnDate time.Time

	err := r.db.QueryRow(ctx, `
		DELETE FROM booking
		WHERE book_id = $1
		RETURNING 
			book_id, session_id, user_id, car_id, affiliator_id, total_price, pickup_date, return_date, status, created_at, updated_at
	`, id).Scan(
		&deletedBooking.BookID,
		&deletedBooking.SessionID,
		&deletedBooking.UserID,
		&deletedBooking.CarID,
		&deletedBooking.AffiliatorID,
		&deletedBooking.TotalPrice,
		&pickupDate,
		&returnDate,
		&deletedBooking.Status,
		&deletedBooking.CreatedAt,
		&deletedBooking.UpdatedAt,
	)

	if err != nil {
		return models.Booking{}, fmt.Errorf("failed to delete booking: %w", err)
	}

	// แปลง Date -> String
	deletedBooking.PickupDate = pickupDate.Format("2006-01-02")
	deletedBooking.ReturnDate = returnDate.Format("2006-01-02")

	return deletedBooking, nil
}
