package interfaces

import (
	"context"
)

type ClientRepository interface {
	CreateClient(ctx context.Context, name, email string) (string, error)
	RevokeClient(ctx context.Context, apiKey string) error
}
