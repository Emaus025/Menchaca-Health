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

type ConsultorioHandler struct {
	repo *database.ConsultorioRepository
}

func NewConsultorioHandler(repo *database.ConsultorioRepository) *ConsultorioHandler {
	return &ConsultorioHandler{repo: repo}
}

func (h *ConsultorioHandler) Create(w http.ResponseWriter, r *http.Request) {
	var consultorio models.Consultorio
	if err := json.NewDecoder(r.Body).Decode(&consultorio); err != nil {
		log.Printf("[ERROR] Error decodificando consultorio: %v", err)
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	if err := h.repo.Create(&consultorio); err != nil {
		log.Printf("[ERROR] Error creando consultorio: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(consultorio)
}

func (h *ConsultorioHandler) List(w http.ResponseWriter, r *http.Request) {
	consultorios, err := h.repo.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(consultorios)
}

func (h *ConsultorioHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	consultorio, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(consultorio)
}

func (h *ConsultorioHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var consultorio models.Consultorio
	if err := json.NewDecoder(r.Body).Decode(&consultorio); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	consultorio.ID = id
	if err := h.repo.Update(&consultorio); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(consultorio)
}

func (h *ConsultorioHandler) Delete(w http.ResponseWriter, r *http.Request) {
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

func (h *ConsultorioHandler) GetByMedico(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	medicoID, err := strconv.Atoi(vars["medicoId"])
	if err != nil {
		http.Error(w, "ID de médico inválido", http.StatusBadRequest)
		return
	}

	consultorios, err := h.repo.GetByMedico(medicoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(consultorios)
}