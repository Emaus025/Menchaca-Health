package database

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"menchaca-health/internal/models"
)

type AppointmentRepository struct {
	db *Database
}

func NewAppointmentRepository(db *Database) *AppointmentRepository {
	return &AppointmentRepository{db: db}
}

func (r *AppointmentRepository) Create(appointment *models.Appointment) error {
	appointment.ID = uuid.New().String()
	appointment.CreatedAt = time.Now()
	appointment.UpdatedAt = time.Now()

	query := `
		INSERT INTO appointments (
			id, patient_id, date, time, duration, type,
			status, reason, notes, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`

	_, err := r.db.DB.Exec(context.Background(), query,
		appointment.ID, appointment.PatientID, appointment.Date,
		appointment.Time, appointment.Duration, appointment.Type,
		appointment.Status, appointment.Reason, appointment.Notes,
		appointment.CreatedAt, appointment.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("error creating appointment: %v", err)
	}

	return nil
}

func (r *AppointmentRepository) GetByID(id string) (*models.Appointment, error) {
	query := `
		SELECT * FROM appointments WHERE id = $1
	`

	appointment := &models.Appointment{}
	err := r.db.DB.QueryRow(context.Background(), query, id).Scan(
		&appointment.ID, &appointment.PatientID, &appointment.Date,
		&appointment.Time, &appointment.Duration, &appointment.Type,
		&appointment.Status, &appointment.Reason, &appointment.Notes,
		&appointment.CreatedAt, &appointment.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error getting appointment: %v", err)
	}

	return appointment, nil
}

func (r *AppointmentRepository) Update(appointment *models.Appointment) error {
	appointment.UpdatedAt = time.Now()

	query := `
		UPDATE appointments SET
			patient_id = $1, date = $2, time = $3, duration = $4,
			type = $5, status = $6, reason = $7, notes = $8,
			updated_at = $9
		WHERE id = $10
	`

	ct, err := r.db.DB.Exec(context.Background(), query,
		appointment.PatientID, appointment.Date, appointment.Time,
		appointment.Duration, appointment.Type, appointment.Status,
		appointment.Reason, appointment.Notes, appointment.UpdatedAt,
		appointment.ID,
	)

	if err != nil {
		return fmt.Errorf("error updating appointment: %v", err)
	}

	if ct.RowsAffected() == 0 {
		return fmt.Errorf("appointment not found")
	}

	return nil
}

func (r *AppointmentRepository) Delete(id string) error {
	query := `DELETE FROM appointments WHERE id = $1`

	ct, err := r.db.DB.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("error deleting appointment: %v", err)
	}

	if ct.RowsAffected() == 0 {
		return fmt.Errorf("appointment not found")
	}

	return nil
}

func (r *AppointmentRepository) List() ([]models.Appointment, error) {
	query := `SELECT * FROM appointments ORDER BY date ASC, time ASC`

	rows, err := r.db.DB.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("error listing appointments: %v", err)
	}
	defer rows.Close()

	var appointments []models.Appointment
	for rows.Next() {
		var appointment models.Appointment
		err := rows.Scan(
			&appointment.ID, &appointment.PatientID, &appointment.Date,
			&appointment.Time, &appointment.Duration, &appointment.Type,
			&appointment.Status, &appointment.Reason, &appointment.Notes,
			&appointment.CreatedAt, &appointment.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning appointment: %v", err)
		}
		appointments = append(appointments, appointment)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating appointments: %v", err)
	}

	return appointments, nil
}

func (r *AppointmentRepository) GetByPatientID(patientID string) ([]models.Appointment, error) {
	query := `
		SELECT * FROM appointments
		WHERE patient_id = $1
		ORDER BY date ASC, time ASC
	`

	rows, err := r.db.DB.Query(context.Background(), query, patientID)
	if err != nil {
		return nil, fmt.Errorf("error getting patient appointments: %v", err)
	}
	defer rows.Close()

	var appointments []models.Appointment
	for rows.Next() {
		var appointment models.Appointment
		err := rows.Scan(
			&appointment.ID, &appointment.PatientID, &appointment.Date,
			&appointment.Time, &appointment.Duration, &appointment.Type,
			&appointment.Status, &appointment.Reason, &appointment.Notes,
			&appointment.CreatedAt, &appointment.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning appointment: %v", err)
		}
		appointments = append(appointments, appointment)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating appointments: %v", err)
	}

	return appointments, nil
}

func (r *AppointmentRepository) GetByDate(date string) ([]models.Appointment, error) {
	query := `
		SELECT * FROM appointments
		WHERE date = $1
		ORDER BY time ASC
	`

	rows, err := r.db.DB.Query(context.Background(), query, date)
	if err != nil {
		return nil, fmt.Errorf("error getting appointments by date: %v", err)
	}
	defer rows.Close()

	var appointments []models.Appointment
	for rows.Next() {
		var appointment models.Appointment
		err := rows.Scan(
			&appointment.ID, &appointment.PatientID, &appointment.Date,
			&appointment.Time, &appointment.Duration, &appointment.Type,
			&appointment.Status, &appointment.Reason, &appointment.Notes,
			&appointment.CreatedAt, &appointment.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning appointment: %v", err)
		}
		appointments = append(appointments, appointment)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating appointments: %v", err)
	}

	return appointments, nil
}