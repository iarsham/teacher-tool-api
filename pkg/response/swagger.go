package response

import "github.com/iarsham/teacher-tool-api/internal/models"

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
