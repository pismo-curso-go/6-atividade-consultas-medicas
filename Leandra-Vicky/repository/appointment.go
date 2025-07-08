package repository

import (
	"api/banco"
	"api/model"
	"errors"
	"log"
	"time"
)

func CreateAppointment(patientID string, datetime time.Time) (model.Appointment, error) {
	if datetime.IsZero() {
		return model.Appointment{}, errors.New("Horário do agentamento inválido")
	}
	if datetime.Before(time.Now()) {
		return model.Appointment{}, errors.New("Agendamentos não podem ser no passado.")
	}

	var count int
	err := banco.DB.Get(&count, `SELECT COUNT(*) FROM appointments WHERE patient_id = $1 AND datetime = $2`, patientID, datetime)
	if err != nil {
		return model.Appointment{}, err
	}
	if count > 0 {
		return model.Appointment{}, errors.New("Horário já reservado para esse paciente.")
	}

	_, err = banco.DB.Exec(`INSERT INTO appointments (patient_id, datetime) VALUES ($1, $2)`, patientID, datetime)
	if err != nil {
		return model.Appointment{}, err
	}

	return model.Appointment{
		PatientID: patientID,
		Datetime:  datetime,
	}, nil
}

func GetAppointmentsByPatientID(patientID string) ([]model.Appointment, error) {
	var appointment []model.Appointment
	err := banco.DB.Select(&appointment, `SELECT * FROM appointments WHERE patient_id = $1`, patientID)
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}

	return appointment, nil
}

func DeleteAppointment(patientID, id string) error {
	result, err := banco.DB.Exec(`DELETE FROM appointments WHERE id = $1 AND patient_id = $2`, id, patientID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("Nenhum agendamento encontrado para esse paciente")
	}
	return nil
}
