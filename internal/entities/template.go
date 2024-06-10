package entities

import "mime/multipart"

type TemplateRequest struct {
	File   *multipart.FileHeader `form:"file" validate:"required,file"`
	UserID uint64
}
