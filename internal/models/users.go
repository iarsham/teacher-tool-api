package models

type Users struct {
	ID        uint64 `json:"id"`
	Phone     string `json:"phone"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
}
