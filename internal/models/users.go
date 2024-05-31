package models

type Users struct {
	ID        uint64 `json:"id" example:"1"`
	Phone     string `json:"phone" example:"+989029266610"`
	Password  string `json:"-"`
	Role      Role   `json:"role" example:"1"`
	CreatedAt string `json:"created_at" example:"2024-01-29T03:09:00+03:30"`
}
