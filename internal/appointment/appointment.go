package appointment

import (
	"time"
)

type AppointmentDomain struct {
	date time.Time
}

func NewAppointmentDomain(date time.Time) (*AppointmentDomain, error) {
	currentDatetime := time.Now()
	if time.Now().Compare(currentDatetime) >= 0 {
		return nil, ErrAppointmentInvalidDate
	}

	return &AppointmentDomain{
		date: date,
	}, nil
}

func (a *AppointmentDomain) Date() time.Time {
	return a.date
}
