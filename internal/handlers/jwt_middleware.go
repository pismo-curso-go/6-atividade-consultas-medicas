package handlers

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// Middleware para autenticação JWT
func AuthMiddleware(db *sql.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"message": "Token não fornecido",
					"code":    401,
				})
			}

			tokenStr := authHeader[len("Bearer "):]
			secret := os.Getenv("JWT_SECRET")
			if secret == "" {
				secret = "segredo_padrao"
			}

			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, echo.NewHTTPError(http.StatusUnauthorized, "Método de assinatura inválido")
				}
				return []byte(secret), nil
			})

			if err != nil || !token.Valid {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"message": "Token inválido ou expirado",
					"code":    401,
				})
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"message": "Token malformado",
					"code":    401,
				})
			}

			email := claims["sub"].(string)
			var pacienteID int
			err = db.QueryRow("SELECT id FROM pacientes WHERE email = $1", email).Scan(&pacienteID)
			if err != nil {
				return c.JSON(http.StatusForbidden, map[string]interface{}{
					"message": "Paciente não encontrado",
					"code":    403,
				})
			}

			// Salva o paciente_id no contexto para uso nas rotas
			c.Set("paciente_id", pacienteID)

			return next(c)
		}
	}
}
