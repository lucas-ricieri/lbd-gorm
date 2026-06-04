package controller

import (
	"fmt"
	"log"

	"azevedoruan.github/lbd-gorm/internal/repository"
)

// Uma função hipotética para requisição HTTP
func GetPlaylistByID(repo *repository.PlaylistRepository) {
	id := 2
	playlist, err := repo.FindByID(uint(id))
	if err != nil {
		log.Fatalf("Não foi possível buscar a playlist pelo ID %d: %s\n", id, err)
	}
	fmt.Printf("playlist buscada: '%s' de %s\n", playlist.Name, playlist.User.Name)
}
