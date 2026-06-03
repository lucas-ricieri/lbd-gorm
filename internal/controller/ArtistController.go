package controller

import (
	"fmt"
	"log"

	"azevedoruan.github/lbd-gorm/internal/repository"
)

// Uma função hipotética para requisição HTTP
func GetArtistByID(repo *repository.ArtistRepository) {
	artist, err := repo.FindByID(2)
	if err != nil {
		log.Fatalf("Não foi possível buscar o artista pelo ID %d\n", 2)
	}
	fmt.Printf("Artista buscado: %d, %s, %s\n", artist.ID, artist.Name, artist.Nacionality)
}
