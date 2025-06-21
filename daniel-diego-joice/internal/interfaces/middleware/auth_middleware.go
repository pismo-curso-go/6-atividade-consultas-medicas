package middleware

import (
	"saude-mais/internal/utils"
	"strings"

	"github.com/labstack/echo/v4"
)

func JWTMiddleware(jwtSecret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(utils.ErrInvalidToken.Code, utils.ErrInvalidToken)
			}

			// Extract token from "Bearer <token>"
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.JSON(utils.ErrInvalidToken.Code, utils.ErrInvalidToken)
			}

			token := parts[1]
			claims, err := utils.ValidateJWT(token, jwtSecret)
			if err != nil {
				return c.JSON(utils.ErrInvalidToken.Code, utils.ErrInvalidToken)
			}

			// Store user ID in context
			c.Set("user_id", claims.UserID)
			return next(c)
		}
	}
}
