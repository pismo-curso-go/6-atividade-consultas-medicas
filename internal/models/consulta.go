package models

import (
	"database/sql"
	"time"
)

type Consulta struct {
	ID       int
	PatientID int
	Datetime time.Time
}

// Função que busca as consultas de um paciente pelo ID dele
func ListaConsultasPorPaciente(db *sql.DB, patientID int) ([]Consulta, error) {
	rows, err := db.Query("SELECT id, patient_id, datetime FROM consultas WHERE patient_id = $1 ORDER BY datetime", patientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var consultas []Consulta
	for rows.Next() {
		var c Consulta
		if err := rows.Scan(&c.ID, &c.PatientID, &c.Datetime); err != nil {
			return nil, err
		}
		consultas = append(consultas, c)
	}

	return consultas, nil
}
