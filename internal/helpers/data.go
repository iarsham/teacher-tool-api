package helpers

import (
	"errors"
	"net/http"
)

type M map[string]interface{}

var (
	ErrInternalServer = errors.New("the server encountered a problem and could not process your request")
	ErrInvalidClaims  = errors.New("claims are invalid and could not be processed")
)

func GetUserID(r *http.Request) uint64 {
	return uint64(r.Context().Value("user_id").(float64))
}
