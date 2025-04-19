package model

type Client struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	APIKey string `json:"api_key"`
}
