package main

import (
	"healthclinic/config"
	patientController "healthclinic/internal/patients/controller"
	patientRepository "healthclinic/internal/patients/repository"
	patientUseCase "healthclinic/internal/patients/usecase"

	appointmentController "healthclinic/internal/appointment/controller"
	appointmentRepository "healthclinic/internal/appointment/repository"
	appointmentUseCase "healthclinic/internal/appointment/usecase"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, diContainer config.DIContainer) error {
	pRepo := patientRepository.NewPatientRepository(diContainer.DbInstance())
	pUseCase := patientUseCase.NewPatientUseCase(pRepo)
	pController := patientController.NewPatientController(pUseCase, diContainer.Env())
	e.POST("register", pController.RegisterPacient)
	e.POST("login", pController.Login)

	aRepo := appointmentRepository.NewAppointmentRepository(diContainer.DbInstance())
	aUseCase := appointmentUseCase.NewAppointmentUseCase(aRepo)
	aController := appointmentController.NewAppointmentController(aUseCase)

	jwtMiddleware := config.JWTMiddleware(diContainer.Env().JwtSecret())

	appointments := e.Group("/appointments", jwtMiddleware)
	appointments.GET("", aController.ListByPatientID)
	appointments.POST("", aController.Create)
	appointments.DELETE("/:id", aController.Cancel)

	return nil
}
