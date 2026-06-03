package controller

import (
	"fmt"
	"log"

	"azevedoruan.github/lbd-gorm/internal/repository"
)

// Uma função hipotética para requisição HTTP
func GetMusicByID(repo *repository.MusicRepository) {
	music, err := repo.FindByID(2)
	if err != nil {
		log.Fatalf("Não foi possível buscar o usuário pelo ID %d: %s\n", 2, err)
	}
	fmt.Printf("Musica buscada: %d, %s, %ds, artista: %s\n", music.ID, music.Title, music.DurationSec, music.Artist.Name)
}
