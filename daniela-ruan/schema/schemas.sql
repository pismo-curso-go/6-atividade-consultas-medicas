CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Tabela para armazenar os dados dos pacientes
CREATE TABLE patients (
                          id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                          name VARCHAR(255) NOT NULL,
                          email VARCHAR(255) NOT NULL UNIQUE,
                          password_hash VARCHAR(255) NOT NULL,
                          created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
                          updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_patients_email ON patients(email);

CREATE TYPE appointment_status AS ENUM ('SCHEDULED', 'CANCELED', 'COMPLETED');

CREATE TABLE appointments (
                              id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                              patient_id UUID NOT NULL,
                              appointment_date TIMESTAMP WITH TIME ZONE NOT NULL,
                              status appointment_status NOT NULL DEFAULT 'SCHEDULED',
                              created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
                              updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),

                              CONSTRAINT fk_patient
                                  FOREIGN KEY(patient_id)
                                      REFERENCES patients(id)
                                      ON DELETE CASCADE -- Se o paciente for deletado, seus agendamentos também são.
);

CREATE INDEX idx_appointments_patient_id ON appointments(patient_id);

CREATE INDEX idx_appointments_date ON appointments(appointment_date);