package handlers

import (
	"fmt"
	"net/http"
	"saude-mais/internal/application/dto"
	"saude-mais/internal/application/services"
	"saude-mais/internal/utils"

	"github.com/labstack/echo/v4"
)

type PatientHandler struct {
	patientService *services.PatientService
	authService    *services.AuthService
}

func NewPatientHandler(patientService *services.PatientService, authService *services.AuthService) *PatientHandler {
	return &PatientHandler{
		patientService: patientService,
		authService:    authService,
	}
}

func (h *PatientHandler) Register(c echo.Context) error {
	var req dto.RegisterPatientRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(utils.ErrValidationFailed.Code, utils.ErrValidationFailed)
	}

	if req.Name == "" || req.Email == "" || req.Password == "" {
		return c.JSON(utils.ErrValidationFailed.Code, utils.ErrValidationFailed)
	}

	if len(req.Password) < 6 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Password must be at least 6 characters",
			"code":    http.StatusBadRequest,
		})
	}

	err := h.patientService.Register(c.Request().Context(), &req)
	if err != nil {
		if apiErr, ok := err.(utils.APIError); ok {
			return c.JSON(apiErr.Code, apiErr)
		}
		return err
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Paciente cadastrado com sucesso",
	})
}

func (h *PatientHandler) Login(c echo.Context) error {
	var req dto.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	loginResp, err := h.authService.Login(c.Request().Context(), &req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
	}

	fmt.Println("🚀 Token gerado:", loginResp.Token)

	return c.JSON(http.StatusOK, loginResp)
}
