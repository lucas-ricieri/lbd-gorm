package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"azevedoruan.github/lbd-gorm/internal/models"
	"azevedoruan.github/lbd-gorm/internal/service"
)

type PlaylistController struct {
	Service *service.PlaylistService
}

func (e *PlaylistController) Setup(mux *http.ServeMux) {
	mux.HandleFunc("/playlist/usuario/{id}", e.GetAll)
	mux.HandleFunc("/playlist/{playlist_id}/usuario/{user_id}", e.GetById)
	mux.HandleFunc("/playlist/create/usuario/{id}", e.Create)
	mux.HandleFunc("/playlist/update", e.Update)
	mux.HandleFunc("/playlist/delete", e.Delete)
	mux.HandleFunc("/playlist/addmusic", e.AddMusic)
	mux.HandleFunc("/playlist/removemusic", e.RemoveMusic)
}

func (e *PlaylistController) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed.", http.StatusBadRequest)
		return
	}
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Error to get user ID", http.StatusBadRequest)
		return
	}

	objs, err := e.Service.FindAllFromUserId(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(objs) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(objs)
}

func (e *PlaylistController) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed. Require POST", http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content type not allowed. Require JSON.", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Error to get user ID", http.StatusBadRequest)
		return
	}

	var newPlaylist models.Playlist
	if err := json.NewDecoder(r.Body).Decode(&newPlaylist); err != nil {
		http.Error(w, "Invalid JSON payload. Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := e.Service.Create(uint(id), &newPlaylist); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPlaylist)
}

func (e *PlaylistController) GetById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed. Require GET", http.StatusMethodNotAllowed)
		return
	}

	playlistId, err := strconv.ParseUint(r.PathValue("playlist_id"), 10, 64)
	if err != nil {
		http.Error(w, "Error to get Playlist ID", http.StatusBadRequest)
		return
	}
	userId, err := strconv.ParseUint(r.PathValue("user_id"), 10, 64)
	if err != nil {
		http.Error(w, "Error to get Playlist ID", http.StatusBadRequest)
		return
	}

	obj, err := e.Service.GetById(uint(playlistId), uint(userId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(obj)
}

func (e *PlaylistController) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed. Require PUT", http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content type not allowed. Require JSON.", http.StatusBadRequest)
		return
	}
	var playlist models.Playlist
	if err := json.NewDecoder(r.Body).Decode(&playlist); err != nil {
		http.Error(w, "Invalid JSON payload. Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := e.Service.Update(&playlist); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(playlist)
}

func (e *PlaylistController) AddMusic(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed. Require POST", http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content type not allowed. Require JSON.", http.StatusBadRequest)
		return
	}
	var newMusic models.MusicPlaylist
	if err := json.NewDecoder(r.Body).Decode(&newMusic); err != nil {
		http.Error(w, "Invalid JSON payload. Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := e.Service.AddMusic(&newMusic); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newMusic)
}

func (e *PlaylistController) RemoveMusic(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed. Require PUT", http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content type not allowed. Require JSON.", http.StatusBadRequest)
		return
	}
	var musicToRemove models.MusicPlaylist
	if err := json.NewDecoder(r.Body).Decode(&musicToRemove); err != nil {
		http.Error(w, "Invalid JSON payload. Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := e.Service.RemoveMusic(musicToRemove); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (e *PlaylistController) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed. Require DELETE", http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content type not allowed. Require JSON.", http.StatusBadRequest)
		return
	}
	var playlist models.Playlist
	if err := json.NewDecoder(r.Body).Decode(&playlist); err != nil {
		http.Error(w, "Invalid JSON payload. Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := e.Service.Delete(&playlist); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
