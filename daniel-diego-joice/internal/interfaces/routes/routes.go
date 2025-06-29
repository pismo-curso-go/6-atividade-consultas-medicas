package routes

import (
	"saude-mais/internal/interfaces/handlers"
	"saude-mais/internal/interfaces/middleware"

	"github.com/labstack/echo/v4"
)

func InitRoutes(
	e *echo.Echo,
	patientHandler *handlers.PatientHandler,
	appointmentHandler *handlers.AppointmentHandler,
	jwtSecret string,
) {
	e.POST("/register", patientHandler.Register)
	e.POST("/login", patientHandler.Login)

	protected := e.Group("")
	protected.Use(middleware.JWTMiddleware(jwtSecret))

	protected.POST("/appointments", appointmentHandler.Create)
	protected.GET("/appointments", appointmentHandler.GetAllAppointmentByPatientID)
	protected.GET("/appointments/:id", appointmentHandler.GetByID)
	protected.PUT("/appointments/:id", appointmentHandler.Update)
	protected.PATCH("/appointments/:id", appointmentHandler.Update)
	protected.DELETE("/appointments/:id", appointmentHandler.Delete)
}
