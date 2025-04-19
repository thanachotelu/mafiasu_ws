package models

type Data struct {
	ID      int    `db:"id"`
	Content string `db:"content"`
}

type DataResponse struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

type CreateDataRequest struct {
	Content string `json:"content"`
}
