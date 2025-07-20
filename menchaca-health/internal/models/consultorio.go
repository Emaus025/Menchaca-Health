package models

type Consultorio struct {
	ID              int    `json:"id"`
	TipoConsultorio string `json:"tipo_consultorio"`
	NombreConsultorio string `json:"nombre_consultorio"`
	Ubicacion       string `json:"ubicacion"`
	IDMedico        int    `json:"id_medico"`
}