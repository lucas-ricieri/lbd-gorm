package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"azevedoruan.github/lbd-gorm/internal/models"
	"azevedoruan.github/lbd-gorm/internal/repository"
)

type ArtistController struct {
	Respo *repository.ArtistRepository
}

func (e *ArtistController) Setup(mux *http.ServeMux) {
	mux.HandleFunc("/artista/{id}", e.GetByID)
	mux.HandleFunc("/artistas", e.GetAll)
	mux.HandleFunc("/artista/create", e.Create)
	mux.HandleFunc("/artista/update", e.Update)
	mux.HandleFunc("/artista/delete/{id}", e.DeleteById)
}

func (e *ArtistController) GetByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Error to get ID", http.StatusBadRequest)
		return
	}

	obj, err := e.Respo.FindByID(uint(id))
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(obj)
}

func (e *ArtistController) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusBadRequest)
		return
	}

	objs, err := e.Respo.FindAll()
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(objs)
}

func (e *ArtistController) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed. Require POST", http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content type not allowed. Require JSON.", http.StatusConflict)
		return
	}
	var newArtist models.Artist
	if err := json.NewDecoder(r.Body).Decode(&newArtist); err != nil {
		http.Error(w, "Invalid JSON payload.", http.StatusBadRequest)
		return
	}
	if err := e.Respo.AddNew(&newArtist); err != nil {
		http.Error(w, "Error to create artist.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newArtist)
}

func (e *ArtistController) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed. Require PUT", http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content type not allowed. Require JSON.", http.StatusConflict)
		return
	}
	var updatedArtist models.Artist
	if err := json.NewDecoder(r.Body).Decode(&updatedArtist); err != nil {
		http.Error(w, "Invalid JSON payload.", http.StatusBadRequest)
		return
	}
	if updatedArtist.ID <= 0 {
		http.Error(w, "The ID is required.", http.StatusBadRequest)
		return
	}
	if err := e.Respo.Update(updatedArtist); err != nil {
		http.Error(w, "Error to update artist "+strconv.FormatUint(uint64(updatedArtist.ID), 10), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (e *ArtistController) DeleteById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed. Require DELETE", http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Error to get ID", http.StatusBadRequest)
		return
	}
	if err := e.Respo.DeleteById(uint(id)); err != nil {
		http.Error(w, "Error to delete artist "+strconv.FormatUint(id, 10), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
