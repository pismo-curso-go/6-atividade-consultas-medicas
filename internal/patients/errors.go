package patients

import "errors"

var (
	ErrPatientBlankName       = errors.New("nome do paciente nao pode ficar em branco")
	ErrPatientBlankEmail      = errors.New("email do paciente nao pode ficar em branco")
	ErrPatientBlankPassword   = errors.New("senha do paciente nao pode ficar em branco")
	ErrPatientInvalidPassword = errors.New("as credenciais do paciente nao puderam ser processadas")
	ErrPatientAlreadyExists   = errors.New("Paciente já cadastrado")
)
