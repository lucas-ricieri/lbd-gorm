package main

import (
	"azevedoruan.github/lbd-gorm/cmd/config"
	"azevedoruan.github/lbd-gorm/internal/controller"
	"azevedoruan.github/lbd-gorm/internal/repository"
)

func main() {
	db := config.DBSetup()

	userRepo := repository.UserRepository{DB: db}
	artistRepo := repository.ArtistRepository{DB: db}
	musicRepo := repository.MusicRepository{DB: db}
	playlistRepo := repository.PlaylistRepository{DB: db}
	controller.GetUserByID(&userRepo)
	controller.GetArtistByID(&artistRepo)
	controller.GetMusicByID(&musicRepo)
	controller.GetPlaylistByID(&playlistRepo)
}
