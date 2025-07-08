package config

import "github.com/labstack/echo/v4"

type dataWrapper struct {
	Data any `json:"data"`
}

type messageWrapper struct {
	Message string `json:"message"`
}

// Send a JSON response with a body value using echo context
func ResponseJSON(echoContext echo.Context, status int, body any) error {
	bodyWrapper := dataWrapper{
		Data: body,
	}
	return echoContext.JSON(status, bodyWrapper)
}

// Send a JSON response with a message field into body using echo context
func ResponseMessageJSON(echoContext echo.Context, status int, message string) error {
	errBody := dataWrapper{
		Data: messageWrapper{
			Message: message,
		},
	}
	return echoContext.JSON(status, errBody)
}
