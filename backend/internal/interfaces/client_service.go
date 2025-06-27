package interfaces

import (
	"context"
)

type ClientService interface {
	CreateClient(ctx context.Context, name, email string) (string, error)
	RevokeClient(ctx context.Context, apiKey string) error
}
