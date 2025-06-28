package repository

import (
	"context"
	"fmt"
	"mafiasu_ws/internal/models"
	"mafiasu_ws/internal/utils"

	"github.com/jackc/pgx/v4/pgxpool"
)

type clientRepository struct {
	db *pgxpool.Pool
}

func NewClientRepository(db *pgxpool.Pool) *clientRepository {
	return &clientRepository{db: db}
}

func (r *clientRepository) CreateClient(ctx context.Context, name, email string) (string, error) {
	if r.db == nil {
		return "", fmt.Errorf("database connection is not initialized")
	}

	apiKey, err := utils.GenerateAPIKey()
	if err != nil {
		return "", err
	}

	var id int
	err = r.db.QueryRow(ctx,
		`INSERT INTO clients (name, email, api_key) VALUES ($1, $2, $3) RETURNING id`,
		name, email, apiKey).Scan(&id)

	if err != nil {
		return "", err
	}

	return apiKey, nil
}

func (r *clientRepository) RevokeClient(ctx context.Context, apiKey string) error {
	if r.db == nil {
		return fmt.Errorf("database connection is not initialized")
	}
	cmd, err := r.db.Exec(ctx, `UPDATE clients SET revoked = true WHERE api_key = $1`, apiKey)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return fmt.Errorf("API key not found")
	}
	return nil
}

func (r *clientRepository) GetLogs(ctx context.Context, userID string) ([]models.Log, error) {
	rows, err := r.db.Query(ctx, `
		select user_id, endpoint, method, timestamp
		from logs
		where user_id = $1
		order by timestamp desc
		limit 10
	`, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []models.Log
	for rows.Next() {
		var log models.Log
		if err := rows.Scan(&log.Endpoint, &log.Method, &log.Timestamp); err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}
