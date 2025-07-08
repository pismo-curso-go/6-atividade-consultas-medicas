package appointment

import "errors"

var (
	ErrAppointmentInvalidDate            = errors.New("Consulta no passado não é permitida")
	ErrAppointmentNotFoundOrUnauthorized = errors.New("Consulta não encontrada ou acesso não autorizado")
	ErrAppointmentDoubleBooking          = errors.New("Duas consultas no mesmo horário para o mesmo paciente não são permitidas")
)
