package controller

import (
	"encoding/json"
	"net/http"

	"azevedoruan.github/lbd-gorm/internal/repository"
	"azevedoruan.github/lbd-gorm/internal/service"
)

type TransactionController struct {
	Service *service.TransactionService
}

func (e *TransactionController) Setup(mux *http.ServeMux) {
	mux.HandleFunc("/transacoes/musicas/transferir", e.TransferMusicBetweenPlaylists)
}

func (e *TransactionController) TransferMusicBetweenPlaylists(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed. Require POST", http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content type not allowed. Require JSON.", http.StatusBadRequest)
		return
	}

	var request repository.MusicTransferRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid JSON payload. Error: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := e.Service.TransferMusicBetweenPlaylists(request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
