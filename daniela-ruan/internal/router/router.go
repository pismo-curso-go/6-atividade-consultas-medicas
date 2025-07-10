package router

import (
	"net/http"

	"PacientesEAgendamento/internal/controller"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// SetupRoutes agora recebe a configuração do JWT como um ponteiro
func SetupRoutes(e *echo.Echo, jwtConfig *echojwt.Config, patientController *controller.PatientController, appointmentController *controller.AppointmentController) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "API SaúdeMais está no ar!")
	})

	apiV1 := e.Group("/api/v1")

	// Rotas Públicas
	apiV1.POST("/register", patientController.Register)
	apiV1.POST("/login", patientController.Login)

	// Rotas Protegidas de Agendamentos
	appointments := apiV1.Group("/appointments")
	appointments.Use(echojwt.WithConfig(*jwtConfig)) // Protege todo o grupo

	appointments.POST("", appointmentController.Create)
	appointments.GET("", appointmentController.List)
	appointments.DELETE("/:id", appointmentController.Delete)
}
