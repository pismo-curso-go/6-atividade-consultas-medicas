package config

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	PatientID int    `json:"patient_id"`
	Email     string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(patientID int, email string, secret string) (string, error) {
	claims := JwtCustomClaims{
		PatientID: patientID,
		Email:     email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ParseJWT(tokenStr string, secret string) (*JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
