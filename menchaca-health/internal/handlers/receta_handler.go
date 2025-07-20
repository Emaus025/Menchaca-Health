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

type RecetaHandler struct {
	repo *database.RecetaRepository
}

func NewRecetaHandler(repo *database.RecetaRepository) *RecetaHandler {
	return &RecetaHandler{repo: repo}
}

func (h *RecetaHandler) Create(w http.ResponseWriter, r *http.Request) {
	var receta models.Receta
	if err := json.NewDecoder(r.Body).Decode(&receta); err != nil {
		log.Printf("[ERROR] Error decodificando receta: %v", err)
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	if err := h.repo.Create(&receta); err != nil {
		log.Printf("[ERROR] Error creando receta: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(receta)
}

func (h *RecetaHandler) List(w http.ResponseWriter, r *http.Request) {
	recetas, err := h.repo.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recetas)
}

func (h *RecetaHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	receta, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(receta)
}

func (h *RecetaHandler) GetByConsulta(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	consultaID, err := strconv.Atoi(vars["consultaId"])
	if err != nil {
		http.Error(w, "ID de consulta inválido", http.StatusBadRequest)
		return
	}

	recetas, err := h.repo.GetByConsulta(consultaID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recetas)
}

func (h *RecetaHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var receta models.Receta
	if err := json.NewDecoder(r.Body).Decode(&receta); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	receta.ID = id
	if err := h.repo.Update(&receta); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(receta)
}

func (h *RecetaHandler) Delete(w http.ResponseWriter, r *http.Request) {
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