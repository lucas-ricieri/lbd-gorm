package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"azevedoruan.github/lbd-gorm/internal/repository"
)

type UserController struct {
	Respo *repository.UserRepository
}

func (e *UserController) Setup(mux *http.ServeMux) {
	mux.HandleFunc("/usuario/{id}", e.GetById)
}

func (e *UserController) GetById(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(obj)
}
