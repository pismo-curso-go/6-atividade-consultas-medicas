package patients

import "errors"

var (
	ErrPatientInvalidName     = errors.New("nome do paciente nao pode ficar em branco")
	ErrPatientInvalidEmail    = errors.New("email do paciente nao pode ficar em branco")
	ErrPatientInvalidPassword = errors.New("senha do paciente nao pode ficar em branco")
)
