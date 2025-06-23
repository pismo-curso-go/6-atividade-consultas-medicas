package main

import (
	"healthclinic/config"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	env := config.InitEnvVariables()
	db := config.InitDB(env)

	diContainer := config.NewDIContainer(
		e,
		db,
		env,
	)

	err := initRoutes(e, *diContainer)
	if err != nil {
		log.Fatal("Erro ao iniciar as rotas da API")
	}

	e.Logger.Fatal(e.Start(":" + env.Port()))
}
