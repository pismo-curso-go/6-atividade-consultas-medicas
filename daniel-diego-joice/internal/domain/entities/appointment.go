package entities

import "time"

type Appointment struct {
	ID        int       `json:"id"`
	PatientID int       `json:"patient_id"`
	DateTime  time.Time `json:"date_time"`
}