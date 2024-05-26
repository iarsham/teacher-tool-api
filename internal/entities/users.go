package entities

type UserRequest struct {
	Phone    string `json:"phone" validate:"required,min=10,max=15"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}
