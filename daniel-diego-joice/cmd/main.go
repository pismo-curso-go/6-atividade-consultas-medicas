package main

import (
	"log"
	"saude-mais/internal/application/services"
	"saude-mais/internal/config"
	"saude-mais/internal/infrastructure/database"
	"saude-mais/internal/infrastructure/repositories"
	"saude-mais/internal/interfaces/handlers"
	"saude-mais/internal/interfaces/middleware"
	"saude-mais/internal/interfaces/routes"

	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := database.NewPostgresConnection(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	patientRepo := repositories.NewPatientRepository(db)
	appointmentRepo := repositories.NewAppointmentRepository(db)
	patientService := services.NewPatientService(patientRepo)
	authService := services.NewAuthService(patientRepo, cfg.JWT.Secret)
	appointmentService := services.NewAppointmentService(appointmentRepo)
	patientHandler := handlers.NewPatientHandler(patientService, authService)
	appointmentHandler := handlers.NewAppointmentHandler(appointmentService)

	e := echo.New()

	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.Use(echomiddleware.CORS())

	e.HTTPErrorHandler = middleware.ErrorHandler

	routes.InitRoutes(e, patientHandler, appointmentHandler, cfg.JWT.Secret)

	log.Printf("Server starting on port %s", cfg.Server.Port)
	if err := e.Start(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	log.Println(" Successfully connected to database!")
}
