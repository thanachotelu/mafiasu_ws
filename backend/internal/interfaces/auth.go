package interfaces

import(
	"context"
)



type AuthRepository interface {
	ValidateAPIKey(ctx context.Context, apiKey string) (int, error)
	LogRequest(ctx context.Context, clientID int, endpoint string, method string) error

	
}

