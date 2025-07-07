package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"saude-mais/internal/utils"
)

func ErrorHandler(err error, c echo.Context) {
	if apiErr, ok := err.(utils.APIError); ok {
		c.JSON(apiErr.Code, apiErr)
		return
	}

	if he, ok := err.(*echo.HTTPError); ok {
		c.JSON(he.Code, map[string]interface{}{
			"message": he.Message,
			"code":    he.Code,
		})
		return
	}

	// Default internal server error
	c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"message": "Internal server error",
		"code":    http.StatusInternalServerError,
	})
}
