package entities

import "mime/multipart"

type QuestionRequest struct {
	Lesson uint64 `form:"lesson" validate:"required" example:"1"`
	Title  string `form:"title" validate:"required" example:"Math Stuff"`
	Grade  uint64 `form:"grade" validate:"required,oneof=0 1 2 3 4 5" example:"5"`
	Level  uint64 `form:"level" validate:"required,oneof=0 1 2" example:"2"`
	File   multipart.FileHeader
}
