package patients

import "errors"

var (
	ErrPatientInvalidName  = errors.New("nome do paciente nao pode ser vazio")
	ErrPatientInvalidEmail = errors.New("email do paciente nao pode ser vazio")
)
