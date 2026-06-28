package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"azevedoruan.github/lbd-gorm/internal/models"
	"azevedoruan.github/lbd-gorm/internal/service"
)

type UserController struct {
	Service *service.UserService
}

func (e *UserController) Setup(mux *http.ServeMux) {
	mux.HandleFunc("/usuario/{id}", e.GetById)
	mux.HandleFunc("/usuarios", e.GetAll)
	mux.HandleFunc("/usuario/create", e.Create)
	mux.HandleFunc("/usuario/update", e.Update)
	mux.HandleFunc("/usuario/delete/{id}", e.DeleteById)
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
	obj, err := e.Service.FindByID(uint(id))
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(obj)
}

func (e *UserController) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed. Require GET", http.StatusBadRequest)
		return
	}
	objs, err := e.Service.FindAll()
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(objs)
}

func (e *UserController) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed. Require POST", http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content type not allowed. Require JSON.", http.StatusConflict)
		return
	}
	var newUser models.User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid JSON payload.", http.StatusBadRequest)
		return
	}
	if err := e.Service.Create(&newUser); err != nil {
		http.Error(w, "Error to create user.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func (e *UserController) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed. Require PUT", http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content type not allowed. Require JSON.", http.StatusConflict)
		return
	}
	var updatedUser models.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, "Invalid JSON payload.", http.StatusBadRequest)
		return
	}

	if err := e.Service.Update(updatedUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (e *UserController) DeleteById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed. Require DELETE", http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "Error to get ID", http.StatusBadRequest)
		return
	}

	if err := e.Service.DeleteByID(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
