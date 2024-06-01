package helpers

import "golang.org/x/crypto/bcrypt"

type M map[string]interface{}

func EncryptPass(plainPass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(plainPass), bcrypt.DefaultCost)
}
