package entities

type UserRequest struct {
	Phone    string `json:"phone" validate:"required,min=10,max=15" example:"+9891154326250"`
	Password string `json:"password" validate:"required,min=8,max=32" example:"1qaz2wsx"`
}

type UpdateUserRequest struct {
	Phone string `json:"phone" validate:"required,min=10,max=15" example:"+9891154326250"`
	Role  int    `json:"role"  validate:"oneof=0 1" example:"1" enums:"0,1"`
}
