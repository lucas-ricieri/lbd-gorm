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

	repository := repository.UserRepository{}
	repository.DB = db
	controller.GetUserByID(&repository)
}
