package interfaces

import (
	"context"
	"mafiasu_ws/internal/models"
)

type ClientRepository interface {
	CreateClient(ctx context.Context, name, email string) (string, error)
	RevokeClient(ctx context.Context, apiKey string) error
	GetLogs(ctx context.Context, userID string) ([]models.Log, error)
}
