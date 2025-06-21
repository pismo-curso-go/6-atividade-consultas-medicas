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
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database connection
	db, err := database.NewPostgresConnection(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize repositories
	patientRepo := repositories.NewPatientRepository(db)

	// Initialize services
	patientService := services.NewPatientService(patientRepo)
	authService := services.NewAuthService(patientRepo, cfg.JWT.Secret)

	// Initialize handlers
	patientHandler := handlers.NewPatientHandler(patientService, authService)

	// Initialize Echo
	e := echo.New()

	// Global middleware
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.Recover())
	e.Use(echomiddleware.CORS())

	// Custom error handler
	e.HTTPErrorHandler = middleware.ErrorHandler

	// Setup routes
	routes.InitRoutes(e, patientHandler, cfg.JWT.Secret)

	// Start server
	log.Printf("Server starting on port %s", cfg.Server.Port)
	if err := e.Start(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	// Test database connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	log.Println("✅ Successfully connected to database!")
}
