package database

import (
	"context"
	"fmt"
	"menchaca-health/internal/models"
)

type ExpedienteRepository struct {
	db *Database
}

func NewExpedienteRepository(db *Database) *ExpedienteRepository {
	return &ExpedienteRepository{db: db}
}

func (r *ExpedienteRepository) Create(expediente *models.Expediente) error {
	query := `
		INSERT INTO expedientes (id_paciente, antecedentes_clinicos, historial_clinico, seguro_medico)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	err := r.db.DB.QueryRow(context.Background(), query,
		expediente.IDPaciente,
		expediente.AntecedentesClinic,
		expediente.HistorialClinico,
		expediente.SeguroMedico,
	).Scan(&expediente.ID)

	if err != nil {
		return fmt.Errorf("error al crear expediente: %v", err)
	}

	return nil
}

func (r *ExpedienteRepository) List() ([]*models.Expediente, error) {
	query := `
		SELECT id, id_paciente, antecedentes_clinicos, historial_clinico, seguro_medico
		FROM expedientes ORDER BY id`

	rows, err := r.db.DB.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("error al listar expedientes: %v", err)
	}
	defer rows.Close()

	var expedientes []*models.Expediente
	for rows.Next() {
		expediente := &models.Expediente{}
		err := rows.Scan(
			&expediente.ID,
			&expediente.IDPaciente,
			&expediente.AntecedentesClinic,
			&expediente.HistorialClinico,
			&expediente.SeguroMedico,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear expediente: %v", err)
		}
		expedientes = append(expedientes, expediente)
	}

	return expedientes, nil
}

func (r *ExpedienteRepository) GetByID(id int) (*models.Expediente, error) {
	query := `
		SELECT id, id_paciente, antecedentes_clinicos, historial_clinico, seguro_medico
		FROM expedientes WHERE id = $1`

	expediente := &models.Expediente{}
	err := r.db.DB.QueryRow(context.Background(), query, id).Scan(
		&expediente.ID,
		&expediente.IDPaciente,
		&expediente.AntecedentesClinic,
		&expediente.HistorialClinico,
		&expediente.SeguroMedico,
	)

	if err != nil {
		return nil, fmt.Errorf("error al obtener expediente: %v", err)
	}

	return expediente, nil
}

func (r *ExpedienteRepository) GetByPaciente(pacienteID int) (*models.Expediente, error) {
	query := `
		SELECT id, id_paciente, antecedentes_clinicos, historial_clinico, seguro_medico
		FROM expedientes WHERE id_paciente = $1`

	expediente := &models.Expediente{}
	err := r.db.DB.QueryRow(context.Background(), query, pacienteID).Scan(
		&expediente.ID,
		&expediente.IDPaciente,
		&expediente.AntecedentesClinic,
		&expediente.HistorialClinico,
		&expediente.SeguroMedico,
	)

	if err != nil {
		return nil, fmt.Errorf("error al obtener expediente por paciente: %v", err)
	}

	return expediente, nil
}

func (r *ExpedienteRepository) Update(expediente *models.Expediente) error {
	query := `
		UPDATE expedientes 
		SET antecedentes_clinicos = $2, historial_clinico = $3, seguro_medico = $4
		WHERE id = $1`

	_, err := r.db.DB.Exec(context.Background(), query,
		expediente.ID,
		expediente.AntecedentesClinic,
		expediente.HistorialClinico,
		expediente.SeguroMedico,
	)

	if err != nil {
		return fmt.Errorf("error al actualizar expediente: %v", err)
	}

	return nil
}

func (r *ExpedienteRepository) Delete(id int) error {
	query := `DELETE FROM expedientes WHERE id = $1`

	_, err := r.db.DB.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar expediente: %v", err)
	}

	return nil
}