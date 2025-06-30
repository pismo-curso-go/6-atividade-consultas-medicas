package model

import "time"

// AppointmentStatus define os possíveis status de um agendamento.
type AppointmentStatus string

const (
	StatusScheduled AppointmentStatus = "SCHEDULED"
	StatusCanceled  AppointmentStatus = "CANCELED"
	StatusCompleted AppointmentStatus = "COMPLETED"
)

// Appointment representa a estrutura de um agendamento no sistema.
type Appointment struct {
	ID              string            `json:"id"`
	PatientID       string            `json:"patient_id"`
	AppointmentDate time.Time         `json:"appointment_date"`
	Status          AppointmentStatus `json:"status"`
	CreatedAt       time.Time         `json:"created_at"`
	UpdatedAt       time.Time         `json:"updated_at"`
}
