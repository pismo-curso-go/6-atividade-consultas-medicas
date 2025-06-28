package repository

import "errors"

var (
	ErrFailedQueryExec  = errors.New("erro ao executar query")
	ErrFailedRowScan    = errors.New("erro ao scanear a linha retornada")
	ErrInvalidIteration = errors.New("erro durante a iteracao nas linhas retornadas")
)
