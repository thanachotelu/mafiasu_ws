package models

import "time"

type Booking struct {
	BookID       string    `db:"book_id" json:"book_id"`
	UserID       string    `db:"user_id" json:"user_id"`
	CarID        string    `db:"car_id" json:"car_id"`
	AffiliatorID *string   `db:"affiliator_id" json:"affiliator_id"`
	TotalPrice   float64   `db:"total_price" json:"total_price"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}

type BookingRequest struct {
	UserID       string  `json:"user_id" binding:"required"`
	CarID        string  `json:"car_id" binding:"required"`
	AffiliatorID *string `json:"affiliator_id,omitempty"`
	TotalPrice   float64 `json:"total_price"`
}
