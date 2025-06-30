package controller

import (
	"PacientesEAgendamento/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type PatientController struct {
	patientUsecase *usecase.PatientUsecase
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewPatientController(uc *usecase.PatientUsecase) *PatientController {
	return &PatientController{patientUsecase: uc}
}

func (pc *PatientController) Register(c echo.Context) error {
	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, APIError{Message: "Corpo da requisição inválido", Code: http.StatusBadRequest})
	}

	_, err := pc.patientUsecase.Register(c.Request().Context(), req.Name, req.Email, req.Password)
	if err != nil {
		if err.Error() == "email já cadastrado" {
			return c.JSON(http.StatusConflict, APIError{Message: "Paciente já cadastrado", Code: http.StatusConflict})
		}
		return c.JSON(http.StatusBadRequest, APIError{Message: err.Error(), Code: http.StatusBadRequest})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Paciente cadastrado com sucesso!"})
}

func (pc *PatientController) Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, APIError{Message: "Corpo da requisição inválido", Code: http.StatusBadRequest})
	}

	token, err := pc.patientUsecase.Login(c.Request().Context(), req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, APIError{Message: "Credenciais inválidas", Code: http.StatusUnauthorized})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": token})
}
