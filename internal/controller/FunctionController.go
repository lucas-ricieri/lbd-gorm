package controller

import (
	"encoding/json"
	"net/http"

	"azevedoruan.github/lbd-gorm/internal/service"
)

type FunctionController struct {
	Service *service.FunctionService
}

func (e *FunctionController) Setup(mux *http.ServeMux) {
	mux.HandleFunc("/funcoes/artistas/ranking-popularidade", e.GetArtistPopularityRank)
	mux.HandleFunc("/funcoes/musicas/artista/{artist_name}/maiores-que-maior-musica/{comparison_artist_name}", e.GetMusicsLongerThanArtistMax)
}

func (e *FunctionController) GetArtistPopularityRank(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed. Require GET", http.StatusMethodNotAllowed)
		return
	}

	objs, err := e.Service.RankArtistsByPlaylistPopularity()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(objs)
}

func (e *FunctionController) GetMusicsLongerThanArtistMax(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed. Require GET", http.StatusMethodNotAllowed)
		return
	}

	objs, err := e.Service.FindMusicsLongerThanArtistMax(r.PathValue("artist_name"), r.PathValue("comparison_artist_name"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(objs)
}
