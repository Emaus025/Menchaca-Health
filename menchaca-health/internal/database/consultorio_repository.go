package database

import (
	"context"
	"fmt"
	"menchaca-health/internal/models"
)

type ConsultorioRepository struct {
	db *Database
}

func NewConsultorioRepository(db *Database) *ConsultorioRepository {
	return &ConsultorioRepository{db: db}
}

func (r *ConsultorioRepository) Create(consultorio *models.Consultorio) error {
	query := `
		INSERT INTO consultorios (tipo_consultorio, nombre_consultorio, ubicacion, id_medico)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	err := r.db.DB.QueryRow(context.Background(), query,
		consultorio.TipoConsultorio,
		consultorio.NombreConsultorio,
		consultorio.Ubicacion,
		consultorio.IDMedico,
	).Scan(&consultorio.ID)

	if err != nil {
		return fmt.Errorf("error al crear consultorio: %v", err)
	}

	return nil
}

func (r *ConsultorioRepository) List() ([]*models.Consultorio, error) {
	query := `
		SELECT id, tipo_consultorio, nombre_consultorio, ubicacion, id_medico
		FROM consultorios ORDER BY nombre_consultorio`

	rows, err := r.db.DB.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("error al listar consultorios: %v", err)
	}
	defer rows.Close()

	var consultorios []*models.Consultorio
	for rows.Next() {
		consultorio := &models.Consultorio{}
		err := rows.Scan(
			&consultorio.ID,
			&consultorio.TipoConsultorio,
			&consultorio.NombreConsultorio,
			&consultorio.Ubicacion,
			&consultorio.IDMedico,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear consultorio: %v", err)
		}
		consultorios = append(consultorios, consultorio)
	}

	return consultorios, nil
}

func (r *ConsultorioRepository) GetByID(id int) (*models.Consultorio, error) {
	query := `
		SELECT id, tipo_consultorio, nombre_consultorio, ubicacion, id_medico
		FROM consultorios WHERE id = $1`

	consultorio := &models.Consultorio{}
	err := r.db.DB.QueryRow(context.Background(), query, id).Scan(
		&consultorio.ID,
		&consultorio.TipoConsultorio,
		&consultorio.NombreConsultorio,
		&consultorio.Ubicacion,
		&consultorio.IDMedico,
	)

	if err != nil {
		return nil, fmt.Errorf("error al obtener consultorio: %v", err)
	}

	return consultorio, nil
}

func (r *ConsultorioRepository) GetByMedico(medicoID int) ([]*models.Consultorio, error) {
	query := `
		SELECT id, tipo_consultorio, nombre_consultorio, ubicacion, id_medico
		FROM consultorios WHERE id_medico = $1 ORDER BY nombre_consultorio`

	rows, err := r.db.DB.Query(context.Background(), query, medicoID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener consultorios por médico: %v", err)
	}
	defer rows.Close()

	var consultorios []*models.Consultorio
	for rows.Next() {
		consultorio := &models.Consultorio{}
		err := rows.Scan(
			&consultorio.ID,
			&consultorio.TipoConsultorio,
			&consultorio.NombreConsultorio,
			&consultorio.Ubicacion,
			&consultorio.IDMedico,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear consultorio: %v", err)
		}
		consultorios = append(consultorios, consultorio)
	}

	return consultorios, nil
}

func (r *ConsultorioRepository) Update(consultorio *models.Consultorio) error {
	query := `
		UPDATE consultorios 
		SET tipo_consultorio = $2, nombre_consultorio = $3, ubicacion = $4, id_medico = $5
		WHERE id = $1`

	_, err := r.db.DB.Exec(context.Background(), query,
		consultorio.ID,
		consultorio.TipoConsultorio,
		consultorio.NombreConsultorio,
		consultorio.Ubicacion,
		consultorio.IDMedico,
	)

	if err != nil {
		return fmt.Errorf("error al actualizar consultorio: %v", err)
	}

	return nil
}

func (r *ConsultorioRepository) Delete(id int) error {
	query := `DELETE FROM consultorios WHERE id = $1`

	_, err := r.db.DB.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar consultorio: %v", err)
	}

	return nil
}