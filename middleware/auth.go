package middleware

import (
	"net/http"
	"os"
	"saudemais-api/auth"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY")),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(auth.JwtCustomClaims)
		},
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"message": "Token inválido ou expirado",
				"code":    http.StatusUnauthorized,
			})
		},
	})
}
