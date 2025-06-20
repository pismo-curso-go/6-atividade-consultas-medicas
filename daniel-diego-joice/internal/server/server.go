package server

import (
	"database/sql"
	"os"
	"saude-mais/internal/router"

	"github.com/labstack/echo/v4"
)

type Server struct {
	httpServer *echo.Echo
	DB         *sql.DB
	Port       string
}

const (
	DefaultPort = "3000"
)

func NewServer(db *sql.DB) *Server {

	port := os.Getenv("PORT")
	
	if port == "" {
		port = DefaultPort
	}

	httpServer := echo.New()

	router.InitRoutes(db, httpServer)

	return &Server{
		httpServer: httpServer,
		DB:         db,
		Port:       port,
	}
}

func (s *Server) Start() error {
	s.httpServer.Logger.Fatal(s.httpServer.Start(":" + s.Port))
	return nil
}
