package main

import (
	"context"
	"log"
	"net/http"

	"PacientesEAgendamento/config"
	"PacientesEAgendamento/internal/controller"
	"PacientesEAgendamento/internal/postgres"
	"PacientesEAgendamento/internal/router"
	"PacientesEAgendamento/internal/usecase"
	_ "github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.LoadConfig()

	dbpool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Não foi possível conectar ao banco de dados: %v\n", err)
	}
	defer dbpool.Close()
	log.Println("Conexão com o banco de dados estabelecida.")

	e := echo.New()
	// Adiciona o nosso handler de erro customizado
	e.HTTPErrorHandler = controller.CustomHTTPErrorHandler

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Configuração do Middleware JWT (sem o campo NewClaims)
	jwtConfig := echojwt.Config{
		SigningKey:  []byte(cfg.JWTSecretKey),
		TokenLookup: "header:Authorization:Bearer ",
		// Adiciona um handler de erro customizado para o JWT
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, controller.APIError{
				Message: "Token inválido ou expirado",
				Code:    http.StatusUnauthorized,
			})
		},
	}

	// Injeção de Dependências
	patientRepo := postgres.NewPostgresPatientRepository(dbpool)
	appointmentRepo := postgres.NewPostgresAppointmentRepository(dbpool)

	patientUsecase := usecase.NewPatientUsecase(patientRepo, cfg.JWTSecretKey)
	appointmentUsecase := usecase.NewAppointmentUsecase(appointmentRepo)

	patientController := controller.NewPatientController(patientUsecase)
	appointmentController := controller.NewAppointmentController(appointmentUsecase)

	// Configurar Rotas
	router.SetupRoutes(e, &jwtConfig, patientController, appointmentController)

	log.Println("Servidor iniciando na porta :8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
