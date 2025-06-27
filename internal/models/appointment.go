package models

import "time"

type Appointment struct {
	ID         int       `json:"id"`
	PacienteID int       `json:"paciente_id"`
	Datetime   time.Time `json:"datetime"`
}

// ListaConsultasPorPaciente retorna todas as consultas do paciente autenticado
func ListaConsultasPorPaciente(db *sql.DB, pacienteID int) ([]Appointment, error) {
	rows, err := db.Query(`
		SELECT id, paciente_id, datetime 
		FROM consultas 
		WHERE paciente_id = $1 
		ORDER BY datetime`, pacienteID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var consultas []Appointment
	for rows.Next() {
		var c Appointment
		if err := rows.Scan(&c.ID, &c.PacienteID, &c.Datetime); err != nil {
			return nil, err
		}
		consultas = append(consultas, c)
	}
	return consultas, nil
}
