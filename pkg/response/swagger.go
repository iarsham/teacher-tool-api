package response

import (
	"github.com/iarsham/teacher-tool-api/internal/models"
	"time"
)

type HealthCheck struct {
	Status string `example:"available"`
}

type UserCreated struct {
	Response string `example:"user created"`
}

type BadRequest struct {
	Error string `example:"bad request"`
}

type UserAlreadyExists struct {
	Error string `example:"user already exists"`
}

type InternalServerError struct {
	Error string `example:"internal server error"`
}

type LoginSuccess struct {
	AccessToken  string `example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"`
	RefreshToken string `example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"`
}

type UserNotFound struct {
	Error string `example:"user not found"`
}

type WrongPassword struct {
	Error string `example:"wrong password"`
}

type UserData models.Users

type PasswordChanged struct {
	Response string `example:"password changed successfully"`
}

type AllTemplates []models.Templates

type TemplateCreated struct {
	Response string `example:"template created"`
}

type TemplateExists struct {
	Error string `example:"template already exists"`
}

type AllQuestions []questionResponse

type QuestionNotFound struct {
	Error string `example:"question not found"`
}

type QuestionData questionResponse

type QuestionAlreadyExists struct {
	Error string `example:"question already exists"`
}

type QuestionCreated struct {
	Response string `example:"question created"`
}

type questionResponse struct {
	ID        uint64    `example:"1"`
	Lesson    string    `example:"12"`
	Title     string    `example:"Math statistics"`
	Grade     string    `example:"fifth"`
	Level     string    `example:"advanced"`
	Views     uint64    `example:"666"`
	Used      uint64    `example:"195"`
	File      string    `example:"domain.com/media/image.png"`
	UserID    uint64    `example:"10"`
	CreatedAt time.Time `example:"2024-01-29T03:09:00+03:30"`
}
