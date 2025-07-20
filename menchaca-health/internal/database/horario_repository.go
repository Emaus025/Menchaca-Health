package database

import (
	"context"
	"fmt"
	"menchaca-health/internal/models"
)

type HorarioRepository struct {
	db *Database
}

func NewHorarioRepository(db *Database) *HorarioRepository {
	return &HorarioRepository{db: db}
}

func (r *HorarioRepository) Create(horario *models.Horario) error {
	query := `
		INSERT INTO horarios (id_medico, dia_semana, hora_inicio, hora_fin, disponible)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	err := r.db.DB.QueryRow(context.Background(), query,
		horario.IDMedico,
		horario.DiaSemana,
		horario.HoraInicio,
		horario.HoraFin,
		horario.Disponible,
	).Scan(&horario.ID)

	if err != nil {
		return fmt.Errorf("error al crear horario: %v", err)
	}

	return nil
}

func (r *HorarioRepository) List() ([]*models.Horario, error) {
	query := `
		SELECT id, id_medico, dia_semana, hora_inicio, hora_fin, disponible
		FROM horarios ORDER BY id_medico, dia_semana, hora_inicio`

	rows, err := r.db.DB.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("error al listar horarios: %v", err)
	}
	defer rows.Close()

	var horarios []*models.Horario
	for rows.Next() {
		horario := &models.Horario{}
		err := rows.Scan(
			&horario.ID,
			&horario.IDMedico,
			&horario.DiaSemana,
			&horario.HoraInicio,
			&horario.HoraFin,
			&horario.Disponible,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear horario: %v", err)
		}
		horarios = append(horarios, horario)
	}

	return horarios, nil
}

func (r *HorarioRepository) GetByID(id int) (*models.Horario, error) {
	query := `
		SELECT id, id_medico, dia_semana, hora_inicio, hora_fin, disponible
		FROM horarios WHERE id = $1`

	horario := &models.Horario{}
	err := r.db.DB.QueryRow(context.Background(), query, id).Scan(
		&horario.ID,
		&horario.IDMedico,
		&horario.DiaSemana,
		&horario.HoraInicio,
		&horario.HoraFin,
		&horario.Disponible,
	)

	if err != nil {
		return nil, fmt.Errorf("error al obtener horario: %v", err)
	}

	return horario, nil
}

func (r *HorarioRepository) GetByMedico(medicoID int) ([]*models.Horario, error) {
	query := `
		SELECT id, id_medico, dia_semana, hora_inicio, hora_fin, disponible
		FROM horarios WHERE id_medico = $1 ORDER BY dia_semana, hora_inicio`

	rows, err := r.db.DB.Query(context.Background(), query, medicoID)
	if err != nil {
		return nil, fmt.Errorf("error al obtener horarios por médico: %v", err)
	}
	defer rows.Close()

	var horarios []*models.Horario
	for rows.Next() {
		horario := &models.Horario{}
		err := rows.Scan(
			&horario.ID,
			&horario.IDMedico,
			&horario.DiaSemana,
			&horario.HoraInicio,
			&horario.HoraFin,
			&horario.Disponible,
		)
		if err != nil {
			return nil, fmt.Errorf("error al escanear horario: %v", err)
		}
		horarios = append(horarios, horario)
	}

	return horarios, nil
}

func (r *HorarioRepository) Update(horario *models.Horario) error {
	query := `
		UPDATE horarios 
		SET dia_semana = $2, hora_inicio = $3, hora_fin = $4, disponible = $5
		WHERE id = $1`

	_, err := r.db.DB.Exec(context.Background(), query,
		horario.ID,
		horario.DiaSemana,
		horario.HoraInicio,
		horario.HoraFin,
		horario.Disponible,
	)

	if err != nil {
		return fmt.Errorf("error al actualizar horario: %v", err)
	}

	return nil
}

func (r *HorarioRepository) Delete(id int) error {
	query := `DELETE FROM horarios WHERE id = $1`

	_, err := r.db.DB.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar horario: %v", err)
	}

	return nil
}