package models

import "time"

type Questions struct {
	ID        uint64    `json:"id"`
	Lesson    string    `json:"lesson"`
	Title     string    `json:"title"`
	Grade     Grade     `json:"grade"`
	Level     Level     `json:"level"`
	Views     uint64    `json:"views"`
	Used      uint64    `json:"used"`
	File      string    `json:"file"`
	UserID    uint64    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
