package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"menchaca-health/internal/database"
	"menchaca-health/internal/models"
	"github.com/gorilla/mux"
)

type ExpedienteHandler struct {
	repo *database.ExpedienteRepository
}

func NewExpedienteHandler(repo *database.ExpedienteRepository) *ExpedienteHandler {
	return &ExpedienteHandler{repo: repo}
}

func (h *ExpedienteHandler) Create(w http.ResponseWriter, r *http.Request) {
	var expediente models.Expediente
	if err := json.NewDecoder(r.Body).Decode(&expediente); err != nil {
		log.Printf("[ERROR] Error decodificando expediente: %v", err)
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	if err := h.repo.Create(&expediente); err != nil {
		log.Printf("[ERROR] Error creando expediente: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(expediente)
}

func (h *ExpedienteHandler) List(w http.ResponseWriter, r *http.Request) {
	expedientes, err := h.repo.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expedientes)
}

func (h *ExpedienteHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	expediente, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expediente)
}

func (h *ExpedienteHandler) GetByPaciente(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pacienteID, err := strconv.Atoi(vars["pacienteId"])
	if err != nil {
		http.Error(w, "ID de paciente inválido", http.StatusBadRequest)
		return
	}

	expediente, err := h.repo.GetByPaciente(pacienteID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expediente)
}

func (h *ExpedienteHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var expediente models.Expediente
	if err := json.NewDecoder(r.Body).Decode(&expediente); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	expediente.ID = id
	if err := h.repo.Update(&expediente); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expediente)
}

func (h *ExpedienteHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := h.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}