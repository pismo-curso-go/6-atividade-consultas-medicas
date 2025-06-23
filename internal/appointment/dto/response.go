package dto

import "time"

type AppointmentResponse struct {
	ID        int       `json:"id"`
	PatientID int       `json:"patient_id"`
	Date      time.Time `json:"date"`
}
