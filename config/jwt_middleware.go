package config

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func JWTMiddleware(secret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			header := c.Request().Header.Get("Authorization")
			if header == "" || !strings.HasPrefix(header, "Bearer ") {
				return ResponseMessageJSON(c, http.StatusUnauthorized, "Token inválido ou expirado")
			}
			tokenStr := strings.TrimPrefix(header, "Bearer ")
			claims, err := ParseJWT(tokenStr, secret)
			if err != nil {
				return ResponseMessageJSON(c, http.StatusUnauthorized, "Token inválido ou expirado")
			}
			c.Set("patient_id", claims.PatientID)
			return next(c)
		}
	}
}
