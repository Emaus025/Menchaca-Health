package models

type Expediente struct {
	ID                 int    `json:"id"`
	IDPaciente         int    `json:"id_paciente"`
	AntecedentesClinic string `json:"antecedentes_clinicos"`
	HistorialClinico   string `json:"historial_clinico"`
	SeguroMedico       string `json:"seguro_medico"`
}