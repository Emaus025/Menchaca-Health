package database

import (
	"context"
	"fmt"
	"time"
	"menchaca-health/internal/models"
)

type RecetaRepository struct {
	db *Database
}

func NewRecetaRepository(db *Database) *RecetaRepository {
	return &RecetaRepository{db: db}
}

func (r *RecetaRepository) Create(receta *models.Receta) error {
	query := `
		INSERT INTO recetas (id_consulta, fecha_emision, medicamentos, observaciones)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	if receta.FechaEmision.IsZero() {
		receta.FechaEmision = time.Now()
	}

	err := r.db.DB.QueryRow(context.Background(), query,
		receta.IDConsulta,
		receta.FechaEmision,
		receta.Medicamentos,
		receta.Observaciones,
	).Scan(&receta.ID)

	if err != nil {
		return fmt.Errorf("error al crear receta: %v", err)
	}

	return nil
}

func (r *RecetaRepository) List() ([]*models.Receta, error) {
	query := `
		SELECT id, id_consulta, fecha_emision, medicamentos, observaciones
		FROM recetas ORDER BY fecha_emision DESC`

	rows, err := r.db.DB.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("error al listar recetas: %v", err)
	}
	defer rows.Close()

	var recetas []*models.Receta
	for rows.Next() {
		receta := &models.Receta{}
		err := rows.Scan(
			&receta.ID,
			&receta.IDConsulta,
			&receta.FechaEmision,
			&receta.Medicamentos,
			&receta.Observaciones,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear receta: %v", err)
		}
		recetas = append(recetas, receta)
	}

	return recetas, nil
}

func (r *RecetaRepository) GetByID(id int) (*models.Receta, error) {
	query := `
		SELECT id, id_consulta, fecha_emision, medicamentos, observaciones
		FROM recetas WHERE id = $1`

	receta := &models.Receta{}
	err := r.db.DB.QueryRow(context.Background(), query, id).Scan(
		&receta.ID,
		&receta.IDConsulta,
		&receta.FechaEmision,
		&receta.Medicamentos,
		&receta.Observaciones,
	)

	if err != nil {
		return nil, fmt.Errorf("error al obtener receta: %v", err)
	}

	return receta, nil
}

func (r *RecetaRepository) GetByConsulta(consultaID int) ([]*models.Receta, error) {
	query := `
		SELECT id, id_consulta, fecha_emision, medicamentos, observaciones
		FROM recetas WHERE id_consulta = $1 ORDER BY fecha_emision DESC`

	rows, err := r.db.DB.Query(context.Background(), query, consultaID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener recetas por consulta: %v", err)
	}
	defer rows.Close()

	var recetas []*models.Receta
	for rows.Next() {
		receta := &models.Receta{}
		err := rows.Scan(
			&receta.ID,
			&receta.IDConsulta,
			&receta.FechaEmision,
			&receta.Medicamentos,
			&receta.Observaciones,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear receta: %v", err)
		}
		recetas = append(recetas, receta)
	}

	return recetas, nil
}

func (r *RecetaRepository) Update(receta *models.Receta) error {
	query := `
		UPDATE recetas 
		SET medicamentos = $2, observaciones = $3
		WHERE id = $1`

	_, err := r.db.DB.Exec(context.Background(), query,
		receta.ID,
		receta.Medicamentos,
		receta.Observaciones,
	)

	if err != nil {
		return fmt.Errorf("error al actualizar receta: %v", err)
	}

	return nil
}

func (r *RecetaRepository) Delete(id int) error {
	query := `DELETE FROM recetas WHERE id = $1`

	_, err := r.db.DB.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar receta: %v", err)
	}

	return nil
}