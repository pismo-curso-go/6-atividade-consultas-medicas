package config

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

type DIContainer struct {
	e   *echo.Echo
	db  *sql.DB
	env *EnvVariables
}

func NewDIContainer(e *echo.Echo, db *sql.DB, env *EnvVariables) *DIContainer {
	return &DIContainer{
		e:   e,
		db:  db,
		env: env,
	}
}

func (di *DIContainer) Echo() *echo.Echo {
	return di.e
}

func (di *DIContainer) DbInstance() *sql.DB {
	return di.db
}

func (di *DIContainer) Env() *EnvVariables {
	return di.env
}
