package models

import "time"

type Log struct {
	UserID    string    `db:"user_id"`
	Endpoint  string    `db:"endpoint"`
	Method    string    `db:"method"`
	Timestamp time.Time `db:"timestamp"`
}
