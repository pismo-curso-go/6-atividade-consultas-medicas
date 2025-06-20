package appointment

import "errors"

var (
	ErrAppointmentInvalidDate = errors.New("nao e possivel marcar uma consulta para uma data anterior a atual")
)
