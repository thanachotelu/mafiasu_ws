package service

import (
    "context"
    "crypto/rand"
    "encoding/hex"
    "fmt"
    "mafiasu_ws/internal/models"
    "mafiasu_ws/internal/repository"
)

type AffiliatorService struct {
    repo *repository.AffiliatorRepository
}

func NewAffiliatorService(repo *repository.AffiliatorRepository) *AffiliatorService {
    return &AffiliatorService{repo: repo}
}

func generateAPIKey() string {
    bytes := make([]byte, 16)
    rand.Read(bytes)
    return hex.EncodeToString(bytes)
}

func generateAffiliateCode() string {
    bytes := make([]byte, 4)
    rand.Read(bytes)
    return hex.EncodeToString(bytes)
}

func generateReferralLink() string {
    code := generateAffiliateCode()
    return fmt.Sprintf("https://yourdomain.com/referral/%s", code)
}

func (s *AffiliatorService) CreateAffiliatorForUser(ctx context.Context, userID string) error {
    apiKey := generateAPIKey()
    affiliator := &models.Affiliator{
        UserID:        userID,
        AffiliateCode: generateAffiliateCode(),
        ReferralLink:  generateReferralLink(),
        ApiKey:        apiKey,
    }
    return s.repo.CreateAffiliator(ctx, affiliator)
}
