package appointment

import (
	"time"
)

type AppointmentDomain struct {
	id        int
	date      time.Time
	patientID int
}

func NewAppointmentDomain(id int, date time.Time, patientID int) (*AppointmentDomain, error) {
	currentDatetime := time.Now()
	if time.Now().Compare(currentDatetime) >= 0 {
		return nil, ErrAppointmentInvalidDate
	}

	return &AppointmentDomain{
		id:        id,
		date:      date,
		patientID: patientID,
	}, nil
}

func (p *AppointmentDomain) ID() int {
	return p.id
}

func (a *AppointmentDomain) Date() time.Time {
	return a.date
}

func (a *AppointmentDomain) PatientID() int {
	return a.patientID
}
