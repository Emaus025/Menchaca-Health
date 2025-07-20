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

type HorarioHandler struct {
	repo *database.HorarioRepository
}

func NewHorarioHandler(repo *database.HorarioRepository) *HorarioHandler {
	return &HorarioHandler{repo: repo}
}

func (h *HorarioHandler) Create(w http.ResponseWriter, r *http.Request) {
	var horario models.Horario
	if err := json.NewDecoder(r.Body).Decode(&horario); err != nil {
		log.Printf("[ERROR] Error decodificando horario: %v", err)
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	if err := h.repo.Create(&horario); err != nil {
		log.Printf("[ERROR] Error creando horario: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(horario)
}

func (h *HorarioHandler) List(w http.ResponseWriter, r *http.Request) {
	horarios, err := h.repo.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(horarios)
}

func (h *HorarioHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	horario, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(horario)
}

func (h *HorarioHandler) GetByMedico(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	medicoID, err := strconv.Atoi(vars["medicoId"])
	if err != nil {
		http.Error(w, "ID de médico inválido", http.StatusBadRequest)
		return
	}

	horarios, err := h.repo.GetByMedico(medicoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(horarios)
}

func (h *HorarioHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var horario models.Horario
	if err := json.NewDecoder(r.Body).Decode(&horario); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	horario.ID = id
	if err := h.repo.Update(&horario); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(horario)
}

func (h *HorarioHandler) Delete(w http.ResponseWriter, r *http.Request) {
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