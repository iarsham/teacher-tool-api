package helpers

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func CreateAccessToken(userID uint64, phone string, secretKey string, expire int) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expire)).Unix()
	claims := jwt.MapClaims{
		"user_id": userID,
		"phone":   phone,
		"exp":     exp,
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
}

func CreateRefreshToken(userID uint64, secretKey string, expire int) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expire)).Unix()
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": exp,
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
}

func IsTokenValid(reqToken string, secretKey string) (*jwt.Token, error) {
	token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func GetClaims(token *jwt.Token) (M, error) {
	claimsMap := make(M)
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		for k, v := range claims {
			claimsMap[k] = v
		}
		return claimsMap, nil
	}
	return nil, ErrInvalidClaims
}
