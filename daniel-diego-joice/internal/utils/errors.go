package utils

import "net/http"

type APIError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e APIError) Error() string {
	return e.Message
}

var (
	ErrPatientAlreadyExists = APIError{
		Message: "Paciente já cadastrado",
		Code:    http.StatusConflict,
	}

	ErrInvalidCredentials = APIError{
		Message: "Credenciais inválidas",
		Code:    http.StatusUnauthorized,
	}

	ErrInvalidToken = APIError{
		Message: "Token inválido ou expirado",
		Code:    http.StatusUnauthorized,
	}

	ErrValidationFailed = APIError{
		Message: "Dados de entrada inválidos",
		Code:    http.StatusBadRequest,
	}
)
