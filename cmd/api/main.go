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

	fmt.Println("===== Trabalho de LDB. Grupo 7 =====\nPor Ruan Azevedo e Lucas Ricieri\n")

	fmt.Println("Starting BD...")

	db := config.DBSetup()
	if db.Error != nil {
		log.Fatalf("Error to start BD: %s", db.Error.Error())
	}

	mux := http.NewServeMux()

	fmt.Println("Starting resources...")

	// Add new repositories here
	userRepo := repository.UserRepository{DB: db}
	artistRepo := repository.ArtistRepository{DB: db}

	// Add new controllers here
	userContr := controller.UserController{Respo: &userRepo}
	artistContr := controller.ArtistController{Respo: &artistRepo}

	// Must to setup method in the mux for each controllers
	artistContr.Setup(mux)
	userContr.Setup(mux)

	fmt.Println("Starting listen and server...")

	log.Fatal(http.ListenAndServe(":8080", mux))
}
