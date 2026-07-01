package main

import (
	"fmt"
	"log"
	"net/http"

	"azevedoruan.github/lbd-gorm/cmd/config"
	"azevedoruan.github/lbd-gorm/internal/controller"
	"azevedoruan.github/lbd-gorm/internal/repository"
	"azevedoruan.github/lbd-gorm/internal/service"
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
	relantionshipRepo := repository.RelantionshipRepository{DB: db}
	otimizationDetailRepo := repository.OtimizationDetailRepository{DB: db}
	keysJunctionRepo := repository.KeysJunctionRepository{DB: db}
	functionRepo := repository.FunctionRepository{DB: db}
	transactionRepo := repository.TransactionRepository{DB: db}

	// Add new services here
	userService := service.UserService{Repo: &userRepo}
	artistService := service.ArtistService{Repo: &artistRepo}
	musicService := service.MusicService{Repo: &musicRepo, ArtistRepo: &artistRepo}
	playlistService := service.PlaylistService{Repo: &playlistRepo, UserRepo: &userRepo, MusicRepo: &musicRepo}
	relantionshipService := service.RelantionshipService{Repo: &relantionshipRepo}
	otimizationDetailService := service.OtimizationDetailService{Repo: &otimizationDetailRepo}
	keysJunctionService := service.KeysJunctionService{Repo: &keysJunctionRepo}
	functionService := service.FunctionService{Repo: &functionRepo}
	transactionService := service.TransactionService{Repo: &transactionRepo}

	// Add new controllers here
	userContr := controller.UserController{Service: &userService}
	artistContr := controller.ArtistController{Service: &artistService}
	musicContr := controller.MusicController{Service: &musicService}
	playlistContr := controller.PlaylistController{Service: &playlistService}
	relantionshipContr := controller.RelantionshipController{Service: &relantionshipService}
	otimizationDetailContr := controller.OtimizationDetailController{Service: &otimizationDetailService}
	keysJunctionContr := controller.KeysJunctionController{Service: &keysJunctionService}
	functionContr := controller.FunctionController{Service: &functionService}
	transactionContr := controller.TransactionController{Service: &transactionService}

	// Must to setup method in the mux for each controllers
	artistContr.Setup(mux)
	musicContr.Setup(mux)
	userContr.Setup(mux)
	playlistContr.Setup(mux)
	relantionshipContr.Setup(mux)
	otimizationDetailContr.Setup(mux)
	keysJunctionContr.Setup(mux)
	functionContr.Setup(mux)
	transactionContr.Setup(mux)

	fmt.Println("Starting listen and server...")

	log.Fatal(http.ListenAndServe(":8080", mux))
}
