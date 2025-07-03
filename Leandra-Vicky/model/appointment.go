package model

import (
	"time"
)

type Appointment struct {
	ID        string    `db:"id" json:"id"`
	PatientID string    `db:"patient_id" json:"patient_id"`
	Datetime  time.Time `db:"datetime" json:"datetime"`
}
