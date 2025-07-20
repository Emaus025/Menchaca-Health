package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"menchaca-health/internal/database"
	"menchaca-health/internal/models"
)

type AppointmentHandler struct {
	repo *database.AppointmentRepository
}

func NewAppointmentHandler(repo *database.AppointmentRepository) *AppointmentHandler {
	return &AppointmentHandler{repo: repo}
}

func (h *AppointmentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var appointment models.Appointment
	if err := json.NewDecoder(r.Body).Decode(&appointment); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.repo.Create(&appointment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(appointment)
}

func (h *AppointmentHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	appointment, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, "Appointment not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(appointment)
}

func (h *AppointmentHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var appointment models.Appointment
	if err := json.NewDecoder(r.Body).Decode(&appointment); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	appointment.ID = id
	if err := h.repo.Update(&appointment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(appointment)
}

func (h *AppointmentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *AppointmentHandler) List(w http.ResponseWriter, r *http.Request) {
	appointments, err := h.repo.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(appointments)
}

func (h *AppointmentHandler) GetByPatientID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	patientID := vars["patientId"]

	appointments, err := h.repo.GetByPatientID(patientID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(appointments)
}

func (h *AppointmentHandler) GetByDate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	date := vars["date"]

	appointments, err := h.repo.GetByDate(date)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(appointments)
}