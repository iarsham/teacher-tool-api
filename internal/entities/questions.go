package entities

type QuestionRequest struct {
	Lesson uint64 `form:"lesson" validate:"required"`
	Title  string `form:"title" validate:"required"`
	Grade  uint64 `form:"grade" validate:"required,oneof=0 1 2 3 4 5"`
	Level  uint64 `form:"level" validate:"required,oneof=0 1 2"`
}
