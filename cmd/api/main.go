package main

import (
	"log"
	"net/http"

	"azevedoruan.github/lbd-gorm/cmd/config"
	"azevedoruan.github/lbd-gorm/internal/controller"
	"azevedoruan.github/lbd-gorm/internal/repository"
)

func main() {
	db := config.DBSetup()
	mux := http.NewServeMux()

	artistRepo := repository.ArtistRepository{DB: db}
	artistContr := controller.ArtistController{Respo: &artistRepo}
	artistContr.Setup(mux)

	log.Fatal(http.ListenAndServe(":8080", mux))

	//userRepo := repository.UserRepository{DB: db}
	//musicRepo := repository.MusicRepository{DB: db}
	//playlistRepo := repository.PlaylistRepository{DB: db}
	//controller.GetUserByID(&userRepo)
	//controller.GetArtistByID(&artistRepo)
	//controller.GetMusicByID(&musicRepo)
	//controller.GetPlaylistByID(&playlistRepo)
}
