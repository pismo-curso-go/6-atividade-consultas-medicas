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

func initRoutes(e *echo.Echo, diContainer config.DIContainer) error {
	pRepo := patientRepository.NewPatientRepository(diContainer.DbInstance())
	pUseCase := patientUseCase.NewPatientUseCase(pRepo)
	pController := patientController.NewPatientController(pUseCase)
	e.POST("register", pController.RegisterPacient)

	aRepo := appointmentRepository.NewAppointmentRepository(diContainer.DbInstance())
	aUseCase := appointmentUseCase.NewAppointmentUseCase(aRepo)
	aController := appointmentController.NewAppointmentController(aUseCase)
	e.GET("/:patient_id/all", aController.ListByPatientID)
	e.POST("/:patient_id", aController.Create)

	return nil
}
