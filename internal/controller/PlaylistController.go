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
	mux.HandleFunc("/playlist/{playlist_id}/usuario/{user_id}", e.GetAllMusics)
	mux.HandleFunc("/playlist/addmusic", e.AddMusic)
	mux.HandleFunc("/playlist/removemusic", e.RemoveMusic)
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
	if err := json.NewDecoder(r.Body).Decode(&newPlaylist); err != nil {
		http.Error(w, "Invalid JSON payload. Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	newPlaylist.UserId = uint(id)

	if err := e.Repos.AddNewForUser(&newPlaylist); err != nil {
		http.Error(w, "Error to create playlist: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newPlaylist)
}

func (e *PlaylistController) GetAllMusics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed. Require GET", http.StatusMethodNotAllowed)
		return
	}

	// Get IDs
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

	obj, err := e.Repos.GetAllMusics(uint(playlistId), uint(userId))
	if err != nil {
		http.Error(w, "Could not found. "+err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(obj)
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

	// START SERVICE - define a ordem na playlist
	obj, err := e.Repos.GetLastSortedMusic(newMusic.PlaylistId, newMusic.UserId)
	if err != nil {
		http.Error(w, "Error to get the last music in the playlist: "+err.Error(), http.StatusBadRequest)
		return
	}
	newMusic.PlaylistOrder = obj.PlaylistOrder + 1
	// END

	if err := e.Repos.AddMusic(&newMusic); err != nil {
		http.Error(w, "Error to add new music in the playlist: "+err.Error(), http.StatusInternalServerError)
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
	var newMusic models.MusicPlaylist
	if err := json.NewDecoder(r.Body).Decode(&newMusic); err != nil {
		http.Error(w, "Invalid JSON payload. Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	// START SERVICE Validações
	if newMusic.MusicId == 0 {
		http.Error(w, "Music ID is required.", http.StatusBadRequest)
		return
	}
	if newMusic.PlaylistId == 0 {
		http.Error(w, "Playlist ID is required.", http.StatusBadRequest)
		return
	}
	if newMusic.UserId == 0 {
		http.Error(w, "User ID is required.", http.StatusBadRequest)
		return
	}
	// END

	if err := e.Repos.RemoveMusic(newMusic); err != nil {
		http.Error(w, "Error to remove music from playlist: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
