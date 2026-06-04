package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"azevedoruan.github/lbd-gorm/internal/repository"
)

type ArtistController struct {
	Respo *repository.ArtistRepository
}

// Constructor
func (ArtistController) NewArtistController(respo *repository.ArtistRepository) *ArtistController {
	return &ArtistController{Respo: respo}
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
