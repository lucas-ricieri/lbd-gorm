package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"azevedoruan.github/lbd-gorm/internal/repository"
)

type PlaylistController struct {
	Repos *repository.PlaylistRepository
}

func (e *PlaylistController) Setup(mux *http.ServeMux) {
	mux.HandleFunc("/playlist/usuario/{id}", e.GetAllFromUser)
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

// Uma função hipotética para requisição HTTP
func GetPlaylistByID(repo *repository.PlaylistRepository) {
	id := 2
	playlist, err := repo.FindByID(uint(id))
	if err != nil {
		log.Fatalf("Não foi possível buscar a playlist pelo ID %d: %s\n", id, err)
	}
	fmt.Printf("playlist buscada: '%s' de %s\n", playlist.Name, playlist.User.Name)
}
