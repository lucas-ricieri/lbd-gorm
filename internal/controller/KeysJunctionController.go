package controller

import (
	"encoding/json"
	"net/http"

	"azevedoruan.github/lbd-gorm/internal/service"
)

type KeysJunctionController struct {
	Service *service.KeysJunctionService
}

func (e *KeysJunctionController) Setup(mux *http.ServeMux) {
	mux.HandleFunc("/chaves-juncao/playlists/{playlist_name}/musicas", e.GetMusicsOrderByPlaylistName)
	mux.HandleFunc("/chaves-juncao/musicas/{music_title}/donos-playlist", e.GetPlaylistOwnersByMusicTitle)
}

func (e *KeysJunctionController) GetMusicsOrderByPlaylistName(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed. Require GET", http.StatusMethodNotAllowed)
		return
	}

	objs, err := e.Service.FindMusicsOrderByPlaylistName(r.PathValue("playlist_name"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(objs)
}

func (e *KeysJunctionController) GetPlaylistOwnersByMusicTitle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed. Require GET", http.StatusMethodNotAllowed)
		return
	}

	objs, err := e.Service.FindPlaylistOwnersByMusicTitle(r.PathValue("music_title"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(objs)
}
