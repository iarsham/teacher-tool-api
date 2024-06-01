package helpers

import (
	"net/http"
	"regexp"
)

func GetUserID(r *http.Request) uint64 {
	return uint64(r.Context().Value("user_id").(float64))
}

func IsPhoneValid(phone string) bool {
	regex := regexp.MustCompile(`^(\+98|0|98|0098)?( |-|[()]){0,2}9[0-9]( |-|[()]){0,2}(?:[0-9]( |-|[()]){0,2}){8}`)
	return regex.MatchString(phone)
}
