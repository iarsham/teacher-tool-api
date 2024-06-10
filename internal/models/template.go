package models

type Templates struct {
	ID        uint64 `json:"id" example:"1"`
	UserID    uint64 `json:"user_id" example:"1"`
	File      string `json:"file" example:"domain.com/media/file.docx"`
	CreatedAt string `json:"created_at" example:"2024-01-29T03:09:00+03:30"`
}
