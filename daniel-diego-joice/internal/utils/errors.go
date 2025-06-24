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
		Message: "Falha na validação",
		Code:    http.StatusBadRequest,
	}

	ErrInvalidDateTime = APIError{
		Message: "Data e hora inválidas",
		Code:    http.StatusBadRequest,
	}

	ErrDateTimeBefore = APIError{
		Message: "Data e hora devem ser no futuro",
		Code:    http.StatusBadRequest,
	}

	ErrPastAppointment = APIError{
		Message: "Consulta no passado não é permitida",
		Code:    http.StatusBadRequest,
	}

	ErrAppointmentConflict = APIError{
		Message: "Paciente já possui uma consulta nesse horário",
		Code:    http.StatusConflict,
	}

	ErrInternalServer = APIError{
		Message: "Erro interno do servidor",
		Code:    http.StatusInternalServerError,
	}
)
