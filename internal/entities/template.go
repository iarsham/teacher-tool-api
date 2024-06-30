package entities

import "mime/multipart"

type TemplateRequest struct {
	File   *multipart.FileHeader `form:"file"`
	UserID uint64                `swaggerignore:"true"`
}
