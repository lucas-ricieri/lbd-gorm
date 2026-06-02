package controller

import (
	"fmt"
	"log"

	"azevedoruan.github/lbd-gorm/internal/repository"
)

// Uma função hipotética para requisição HTTP
func GetUserByID(repo *repository.UserRepository) {
	user, err := repo.FindByID(2)
	if err != nil {
		log.Fatalf("Não foi possível buscar o usuário pelo ID %d\n", 2)
	}
	fmt.Printf("Usuario buscado: %d, %s, %s\n", user.ID, user.Name, user.Email)
}
