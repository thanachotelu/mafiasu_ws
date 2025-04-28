package models

import "time"

type TrackClick struct {
	SessionID    string    `json:"session_id"`
	CarID        string    `json:"car_id"`
	AffiliatorID string    `json:"affiliator_id"`
	ReferralLink string    `json:"referral_link"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type TrackClickRequest struct {
	CarID        string `json:"car_id"`
	ReferralLink string `json:"referral_link"`
}

type TrackClickResponse struct {
	SessionID string `json:"session_id"`
	Message   string `json:"message"`
}
