package model

type ContentReceive struct {
	Title     string `json:"title" binding:"required"`
	TopNumber string    `json:"top_number" binding:"required"`
	Content   string `json:"content" binding:"required"`
}
