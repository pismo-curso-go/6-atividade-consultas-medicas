
-- Tabela de pacientes
CREATE TABLE patient (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL
);

-- Tabela de consultas
CREATE TABLE appointment (
    id SERIAL PRIMARY KEY,
    patient_id SERIAL NOT NULL,
    date_time TIMESTAMP NOT NULL,
    FOREIGN KEY (patient_id) REFERENCES patient(id)
);

-- Inserindo dados de exemplo na tabela de pacientes
INSERT INTO patient (name, email) VALUES 
('Alice Silva', 'alice.silva@example.com'),
('Bruno Costa', 'bruno.costa@example.com'),
('Carla Mendes', 'carla.mendes@example.com');

-- Inserindo dados de exemplo na tabela de consultas
INSERT INTO appointment (patient_id, date_time) VALUES 
(1, CURRENT_DATE + INTERVAL '5 days' + TIME '10:00:00'),
(1, CURRENT_DATE + INTERVAL '15 days' + TIME '14:00:00'),
(1, CURRENT_DATE + INTERVAL '25 days' + TIME '09:00:00'),
(2, CURRENT_DATE + INTERVAL '6 days' + TIME '11:00:00'),
(2, CURRENT_DATE + INTERVAL '16 days' + TIME '15:00:00'),
(2, CURRENT_DATE + INTERVAL '26 days' + TIME '10:00:00'),
(3, CURRENT_DATE + INTERVAL '7 days' + TIME '12:00:00'),
(3, CURRENT_DATE + INTERVAL '17 days' + TIME '16:00:00'),
(3, CURRENT_DATE + INTERVAL '27 days' + TIME '11:00:00');
