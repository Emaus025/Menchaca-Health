package models

import "time"

type Consulta struct {
	ID            int       `json:"id"`
	IDPaciente    int       `json:"id_paciente"`
	IDMedico      int       `json:"id_medico"`
	IDConsultorio int       `json:"id_consultorio"`
	TipoConsulta  string    `json:"tipo_consulta"`
	FechaConsulta time.Time `json:"fecha_consulta"`
	HoraConsulta  time.Time `json:"hora_consulta"`
	Diagnostico   string    `json:"diagnostico"`
	Costo         float64   `json:"costo"`
}