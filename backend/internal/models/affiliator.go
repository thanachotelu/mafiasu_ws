package models

import "time"

type Affiliator struct {
	AffiliatorID    string    `db:"affiliator_id"`
	UserID          string    `db:"user_id"`
	AffiliateCode   string    `db:"affiliate_code"`
	ReferralLink    string    `db:"referral_link"`
	CommissionRate  float64   `db:"commission_rate"`
	TotalCommission float64   `db:"total_commission"`
	Balance         float64   `db:"balance"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}

type AffiliatorResponse struct {
	AffiliatorID    string  `json:"affiliator_id"`
	UserID          string  `json:"user_id"`
	AffiliateCode   string  `json:"affiliate_code"`
	ReferralLink    string  `json:"referral_link"`
	CommissionRate  float64 `json:"commission_rate"`
	TotalCommission float64 `json:"total_commission"`
	Balance         float64 `json:"balance"`
}
