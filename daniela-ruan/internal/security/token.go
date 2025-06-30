package security

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// GenerateJWT cria um novo token JWT para um paciente.
func GenerateJWT(patientID, secretKey string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["sub"] = patientID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
