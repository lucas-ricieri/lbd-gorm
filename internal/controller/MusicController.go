package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"azevedoruan.github/lbd-gorm/internal/models"
	"azevedoruan.github/lbd-gorm/internal/repository"
)

type ArtistFinder interface {
	FindByID(id uint) (models.Artist, error)
}

type MusicController struct {
	Respo        *repository.MusicRepository
	ArtistFinder ArtistFinder
}

func (e *MusicController) Setup(mux *http.ServeMux) {
	mux.HandleFunc("/musica/{id}", e.GetById)
	mux.HandleFunc("/musicas", e.GetAll)
	mux.HandleFunc("/musica/create", e.Create)
	mux.HandleFunc("/musica/update", e.Update)
	mux.HandleFunc("/musica/delete/{id}", e.DeleteById)
}

func (e *MusicController) GetById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed. Require GET", http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Error to get ID", http.StatusBadRequest)
		return
	}
	obj, err := e.Respo.FindByID(uint(id))
	if err != nil {
		http.Error(w, "Music not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(obj)
}

func (e *MusicController) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed. Require GET", http.StatusBadRequest)
		return
	}
	objs, err := e.Respo.FindAll()
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(objs)
}

func (e *MusicController) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed. Require POST", http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content type not allowed. Require JSON.", http.StatusConflict)
		return
	}

	var newMusic models.Music
	if err := json.NewDecoder(r.Body).Decode(&newMusic); err != nil {
		http.Error(w, "Invalid JSON payload.", http.StatusBadRequest)
		return
	}

	// START SERVICE - Verifica se artista existe
	obj, err := e.ArtistFinder.FindByID(uint(newMusic.ArtistId))
	if err != nil || obj.ID == 0 {
		http.Error(w, "Could not found Artist with ID "+strconv.FormatUint(uint64(newMusic.ArtistId), 10)+".", http.StatusBadRequest)
		return
	}
	// END

	if err := e.Respo.AddNew(&newMusic); err != nil {
		http.Error(w, "Error to create music.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMusic)
}

func (e *MusicController) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed. Require PUT", http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content type not allowed. Require JSON.", http.StatusConflict)
		return
	}
	var updatedMusic models.Music
	if err := json.NewDecoder(r.Body).Decode(&updatedMusic); err != nil {
		http.Error(w, "Invalid JSON payload.", http.StatusBadRequest)
		return
	}

	// START SERVICE
	// Verifica se IDs existe
	if updatedMusic.ID <= 0 {
		http.Error(w, "The ID is required.", http.StatusBadRequest)
		return
	}
	if updatedMusic.ArtistId <= 0 {
		http.Error(w, "The Artist ID is required.", http.StatusBadRequest)
		return
	}
	// Verifica se music existe
	music, err := e.Respo.FindByID(uint(updatedMusic.ID))
	if err != nil || music.ID == 0 {
		http.Error(w, "Could not found Music with ID "+strconv.FormatUint(uint64(updatedMusic.ID), 10)+".", http.StatusBadRequest)
		return
	}
	// Verifica se artist existe
	artist, err := e.ArtistFinder.FindByID(uint(updatedMusic.ArtistId))
	if err != nil || artist.ID == 0 {
		http.Error(w, "Could not found Artist with ID "+strconv.FormatUint(uint64(updatedMusic.ArtistId), 10)+".", http.StatusBadRequest)
		return
	}
	// END

	if err := e.Respo.Update(updatedMusic); err != nil {
		http.Error(w, "Error to update music "+strconv.FormatUint(uint64(updatedMusic.ID), 10), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (e *MusicController) DeleteById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed. Require DELETE", http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Error to get ID", http.StatusBadRequest)
		return
	}

	// START SERVICE
	// Verifica se music existe
	music, err := e.Respo.FindByID(uint(id))
	if err != nil || music.ID == 0 {
		http.Error(w, "Could not found Music with ID "+strconv.FormatUint(uint64(id), 10)+".", http.StatusBadRequest)
		return
	}
	// END

	if err := e.Respo.DeleteById(uint(id)); err != nil {
		http.Error(w, "Error to delete music "+strconv.FormatUint(id, 10), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
