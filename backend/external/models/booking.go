package models

import "time"

type Booking struct {
	BookID      string    `json:"book_id"`
	SessionID   string    `json:"session_id"`
	UserID      string    `json:"user_id"`
	CarID       string    `json:"car_id"`
	AffiliateID string    `json:"affiliator_id"`
	TotalPrice  float64   `json:"total_price"`
	PickupDate  string    `json:"pickup_date"`
	ReturnDate  string    `json:"return_date"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BookingDetailResponse struct {
	BookingID  string       `json:"booking_id"`
	Car        CarDetail    `json:"car"`
	Customer   CustomerInfo `json:"customer"`
	PickupDate string       `json:"pickup_date"`
	ReturnDate string       `json:"return_date"`
	TotalPrice float64      `json:"total_price"`
	Status     string       `json:"status"`
}

type BookingRequest struct {
	BookID     string  `json:"book_id"`
	UserID     string  `json:"user_id"`
	CarID      string  `json:"car_id"`
	SessionID  string  `json:"session_id"`
	TotalPrice float64 `json:"total_price"`
	PickupDate string  `json:"pickup_date,omitempty"`
	ReturnDate string  `json:"return_date,omitempty"`
}
