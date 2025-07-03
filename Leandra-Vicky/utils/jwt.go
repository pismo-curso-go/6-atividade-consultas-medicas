package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("your-secret-key")

func GenerateJwt(PatientID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"patient_id": PatientID,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString(jwtKey)
}

func ParseJwt(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if claim, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claim["patient_id"].(string), nil
	}

	return "", err
}
