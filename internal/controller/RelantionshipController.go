package controller

import (
	"encoding/json"
	"net/http"

	"azevedoruan.github/lbd-gorm/internal/service"
)

type RelantionshipController struct {
	Service *service.RelantionshipService
}

func (e *RelantionshipController) Setup(mux *http.ServeMux) {
	mux.HandleFunc("/relacionamentos/playlists/usuario/{username}", e.GetPlaylistsByUsername)
	mux.HandleFunc("/relacionamentos/musicas/usuario/{username}/artista/{artist_name}", e.GetMusicsInUserPlaylistsByArtist)
	mux.HandleFunc("/relacionamentos/playlists/contagem-musicas", e.CountMusicsByPlaylist)
	mux.HandleFunc("/relacionamentos/artistas/sem-musicas-playlists", e.GetArtistsWithoutMusicsInPlaylists)
}

func (e *RelantionshipController) GetPlaylistsByUsername(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed. Require GET", http.StatusMethodNotAllowed)
		return
	}

	objs, err := e.Service.FindPlaylistsByUsername(r.PathValue("username"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(objs)
}

func (e *RelantionshipController) GetMusicsInUserPlaylistsByArtist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed. Require GET", http.StatusMethodNotAllowed)
		return
	}

	objs, err := e.Service.FindMusicsInUserPlaylistsByArtist(r.PathValue("username"), r.PathValue("artist_name"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(objs)
}

func (e *RelantionshipController) CountMusicsByPlaylist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed. Require GET", http.StatusMethodNotAllowed)
		return
	}

	objs, err := e.Service.CountMusicsByPlaylist()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(objs)
}

func (e *RelantionshipController) GetArtistsWithoutMusicsInPlaylists(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed. Require GET", http.StatusMethodNotAllowed)
		return
	}

	objs, err := e.Service.FindArtistsWithoutMusicsInPlaylists()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(objs)
}
