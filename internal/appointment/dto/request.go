package dto

import "time"

type CreateAppointmentRequest struct {
	Date      time.Time `json:"date"`
	PatientID int       `param:"patient_id"`
}

type ListAppointmentRequest struct {
	PatientID int `param:"patient_id"`
}
