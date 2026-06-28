package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"azevedoruan.github/lbd-gorm/internal/service"
)

type OtimizationDetailController struct {
	Service *service.OtimizationDetailService
}

func (e *OtimizationDetailController) Setup(mux *http.ServeMux) {
	mux.HandleFunc("/otimizacoes/musica/{id}/artista", e.GetMusicWithArtistById)
	mux.HandleFunc("/otimizacoes/playlists/tempo-total", e.TotalDurationByPlaylist)
	mux.HandleFunc("/otimizacoes/musicas/menores-que-media-artista", e.GetMusicsShorterThanArtistAverage)
}

func (e *OtimizationDetailController) GetMusicWithArtistById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed. Require GET", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Error to get music ID", http.StatusBadRequest)
		return
	}

	obj, err := e.Service.FindMusicWithArtistById(uint(id))
	if err != nil {
		http.Error(w, "Music not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(obj)
}

func (e *OtimizationDetailController) TotalDurationByPlaylist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed. Require GET", http.StatusMethodNotAllowed)
		return
	}

	objs, err := e.Service.TotalDurationByPlaylist()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(objs)
}

func (e *OtimizationDetailController) GetMusicsShorterThanArtistAverage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed. Require GET", http.StatusMethodNotAllowed)
		return
	}

	objs, err := e.Service.FindMusicsShorterThanArtistAverage()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(objs)
}
