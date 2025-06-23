package dto

import "time"

type Appointment struct {
	ID        int       `json:"id"`
	PatientID int       `json:"patient_id"`
	DateTime  time.Time `json:"date_time"`
}

type CreateAppointmentRequest struct {
	PatientID int       `json:"patient_id" validate:"required"`
	DateTime  time.Time `json:"date_time" validate:"required"`
}

type CreateAppointmentResponse struct {
	ID        int       `json:"id"`
	PatientID int       `json:"patient_id"`
	DateTime  time.Time `json:"date_time"`
}