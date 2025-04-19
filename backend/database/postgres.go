package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"mafiasu_ws/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresDB struct {
	pool        *pgxpool.Pool
	databaseURL string
}

func NewPostgresDB(cfg *config.Config) (*PostgresDB, error) {
	db := &PostgresDB{
		databaseURL: cfg.DatabaseURL,
	}

	if err := db.connect(cfg); err != nil {
		return nil, err
	}

	return db, nil
}

func (db *PostgresDB) connect(cfg *config.Config) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	poolConfig, err := pgxpool.ParseConfig(db.databaseURL)
	if err != nil {
		return fmt.Errorf("failed to parse Postgres config: %w", err)
	}

	poolConfig.MaxConns = 25
	poolConfig.MinConns = 5
	poolConfig.MaxConnLifetime = 10 * time.Minute

	pool, err := pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to Postgres: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return fmt.Errorf("failed to ping Postgres: %w", err)
	}

	db.pool = pool
	log.Println("PostgreSQL connected successfully")
	return nil
}

func (db *PostgresDB) Reconnect() error {
	log.Println("Reconnecting to PostgreSQL")
	db.Close()
	return db.connect(&config.Config{DatabaseURL: db.databaseURL})
}

func (db *PostgresDB) Close() {
	if db.pool != nil {
		db.pool.Close()
		log.Println("PostgreSQL connection closed")
	}
}

func (db *PostgresDB) GetPool() *pgxpool.Pool {
	return db.pool
}
