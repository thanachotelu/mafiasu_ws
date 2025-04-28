package repository

import (
    "context"
    "mafiasu_ws/internal/models"
    "github.com/jackc/pgx/v4/pgxpool"
)

type AffiliatorRepository struct {
    db *pgxpool.Pool
}

func NewAffiliatorRepository(db *pgxpool.Pool) *AffiliatorRepository {
    return &AffiliatorRepository{db}
}

func (r *AffiliatorRepository) CreateAffiliator(ctx context.Context, affiliator *models.Affiliator) error {
    query := `INSERT INTO affiliator (user_id, affiliate_code, referral_link, api_key) VALUES ($1, $2, $3, $4)`
    _, err := r.db.Exec(ctx, query, affiliator.UserID, affiliator.AffiliateCode, affiliator.ReferralLink, affiliator.ApiKey)
    return err
}

