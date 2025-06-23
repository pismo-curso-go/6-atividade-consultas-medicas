package main

import (
	"healthclinic/config"
	patientController "healthclinic/internal/patients/controller"
	patientRepository "healthclinic/internal/patients/repository"
	patientUseCase "healthclinic/internal/patients/usecase"

	"github.com/labstack/echo/v4"
)

func initRoutes(e *echo.Echo, diContainer config.DIContainer) error {
	pRepo := patientRepository.NewPatientRepository(diContainer.DbInstance())
	pUseCase := patientUseCase.NewPatientUseCase(pRepo)
	pController := patientController.NewPatientController(pUseCase)
	e.POST("register", pController.RegisterPacient)

	return nil
}
