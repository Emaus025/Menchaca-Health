package models

import (
	"time"
)

type Appointment struct {
	ID        string    `json:"id" db:"id"`
	PatientID string    `json:"patientId" db:"patient_id"`
	Date      string    `json:"date" db:"date"`
	Time      string    `json:"time" db:"time"`
	Duration  int       `json:"duration" db:"duration"`
	Type      string    `json:"type" db:"type"`
	Status    string    `json:"status" db:"status"`
	Reason    string    `json:"reason" db:"reason"`
	Notes     string    `json:"notes" db:"notes"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}