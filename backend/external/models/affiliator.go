package models

import (
	"time"
)

type Affiliator struct {
	AffiliatorID    string    `db:"affiliator_id" json:"affiliator_id"`
	UserID          string    `db:"user_id" json:"user_id"`
	AffiliateCode   string    `db:"affiliate_code" json:"affiliate_code"`
	ReferralLink    string    `db:"referral_link" json:"referral_link"`
	CommissionRate  float64   `db:"commission_rate" json:"commission_rate"`
	TotalCommission float64   `db:"total_commission" json:"total_commission"`
	Balance         float64   `db:"balance" json:"balance"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at"`
}

type AffiliateClick struct {
	ID          int       `db:"id"`
	AffiliateID string    `db:"affiliate_id"`
	ClickedAt   time.Time `db:"clicked_at"`
}

type AffiliateBookingRequest struct {
	SessionID  string       `json:"session_id"`
	UserID     string       `json:"user_id"`
	CarID      string       `json:"car_id"`
	Customer   CustomerInfo `json:"customer"`
	PickupDate string       `json:"pickup_date"`
	ReturnDate string       `json:"return_date"`
}
