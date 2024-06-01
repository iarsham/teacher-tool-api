package helpers

import "errors"

var (
	ErrInternalServer = errors.New("the server encountered a problem and could not process your request")
	ErrInvalidClaims  = errors.New("claims are invalid and could not be processed")
)
