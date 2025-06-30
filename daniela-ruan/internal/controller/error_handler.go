package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// APIError é a nossa estrutura de erro padrão para respostas JSON.
type APIError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// CustomHTTPErrorHandler intercepta os erros HTTP do Echo e os formata.
func CustomHTTPErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	he, ok := err.(*echo.HTTPError)
	if ok {

		c.JSON(he.Code, APIError{Message: he.Message.(string), Code: he.Code})
	} else {
		c.JSON(http.StatusInternalServerError, APIError{Message: "Erro interno do servidor", Code: http.StatusInternalServerError})
	}
}
