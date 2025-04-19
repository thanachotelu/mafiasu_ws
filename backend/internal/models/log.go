package models

import "time"

type Log struct {
	ID        int       `db:"id"`
	ClientID  int       `db:"client_id"`
	Endpoint  string    `db:"endpoint"`
	Method    string    `db:"method"`
	Timestamp time.Time `db:"timestamp"`
}

type LogResponse struct {
	ID        int    `json:"id"`
	ClientID  int    `json:"client_id"`
	Endpoint  string `json:"endpoint"`
	Method    string `json:"method"`
	Timestamp string `json:"timestamp"`
}
