package models

import "time"

type Receta struct {
	ID            int       `json:"id"`
	IDConsulta    int       `json:"id_consulta"`
	FechaEmision  time.Time `json:"fecha_emision"`
	Medicamentos  string    `json:"medicamentos"`
	Observaciones string    `json:"observaciones"`
}