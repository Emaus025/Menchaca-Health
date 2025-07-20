package models

import "time"

type Horario struct {
	ID           int       `json:"id"`
	IDMedico     int       `json:"id_medico"`
	IDConsultorio int       `json:"id_consultorio"`
	Turno        string    `json:"turno"`
	HoraInicio   time.Time `json:"hora_inicio"`
	HoraFin      time.Time `json:"hora_fin"`
}