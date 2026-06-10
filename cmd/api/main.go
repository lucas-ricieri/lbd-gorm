package main

import (
	"fmt"
	"log"
	"net/http"

	"azevedoruan.github/lbd-gorm/cmd/config"
	"azevedoruan.github/lbd-gorm/internal/controller"
	"azevedoruan.github/lbd-gorm/internal/repository"
)

func main() {

	fmt.Println("===== Trabalho de LBD. Grupo 7 =====\nPor Ruan Azevedo e Lucas Ricieri")

	fmt.Println("\nStarting BD...")

	db := config.DBSetup()
	mux := http.NewServeMux()

	fmt.Println("Starting resources...")

	// Add new repositories here
	userRepo := repository.UserRepository{DB: db}
	artistRepo := repository.ArtistRepository{DB: db}
	musicRepo := repository.MusicRepository{DB: db}
	playlistRepo := repository.PlaylistRepository{DB: db}

	// Add new controllers here
	userContr := controller.UserController{Respo: &userRepo}
	artistContr := controller.ArtistController{Respo: &artistRepo}
	musicContr := controller.MusicController{Respo: &musicRepo, ArtistFinder: &artistRepo}
	playlistContr := controller.PlaylistController{Repos: &playlistRepo, UserFinder: &userRepo, MusicFinder: &musicRepo}

	// Must to setup method in the mux for each controllers
	artistContr.Setup(mux)
	musicContr.Setup(mux)
	userContr.Setup(mux)
	playlistContr.Setup(mux)

	fmt.Println("Starting listen and server...")

	log.Fatal(http.ListenAndServe(":8080", mux))
}
