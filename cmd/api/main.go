package main

import (
	"log"

	"azevedoruan.github/lbd-gorm/internal/controller"
	"azevedoruan.github/lbd-gorm/internal/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=admin123 dbname=lbd_trabalho port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar com o Banco de Dados: ", err)
	}

	userRepo := repository.UserRepository{DB: db}
	artistRepo := repository.ArtistRepository{DB: db}
	musicRepo := repository.MusicRepository{DB: db}
	controller.GetUserByID(&userRepo)
	controller.GetArtistByID(&artistRepo)
	controller.GetMusicByID(&musicRepo)
}
