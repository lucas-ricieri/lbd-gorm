package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"azevedoruan.github/lbd-gorm/internal/models"
	"azevedoruan.github/lbd-gorm/internal/repository"
)

type PlaylistController struct {
	Repos *repository.PlaylistRepository
}

func (e *PlaylistController) Setup(mux *http.ServeMux) {
	mux.HandleFunc("/playlist/usuario/{id}", e.GetAllFromUser)
	mux.HandleFunc("/playlist/create/usuario/{id}", e.AddNewForUser)
}

func (e *PlaylistController) GetAllFromUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed.", http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Error to get user ID", http.StatusBadRequest)
		return
	}

	objs, err := e.Repos.FindAllFromUserId(uint(id))
	if err != nil {
		http.Error(w, "Could not found", http.StatusNotFound)
		return
	}

	if len(objs) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(objs)
}

func (e *PlaylistController) AddNewForUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed. Require POST", http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content type not allowed. Require JSON.", http.StatusBadRequest)
		return
	}

	// Get User ID
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Error to get user ID", http.StatusBadRequest)
		return
	}

	var newPlaylist models.Playlist
	newPlaylist.UserId = uint(id)
	if err := json.NewDecoder(r.Body).Decode(&newPlaylist); err != nil {
		http.Error(w, "Invalid JSON payload. Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := e.Repos.AddNewForUser(&newPlaylist); err != nil {
		http.Error(w, "Error to create playlist: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPlaylist)
}
