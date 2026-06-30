package service

import (
	"errors"
	"fmt"

	"azevedoruan.github/lbd-gorm/internal/models"
	"azevedoruan.github/lbd-gorm/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (s *UserService) FindByID(id uint) (models.User, error) {
	return s.Repo.FindByID(id)
}

func (s *UserService) FindAll() ([]models.User, error) {
	return s.Repo.FindAll()
}

func (s *UserService) Create(newUser *models.User) error {
	return s.Repo.AddNew(newUser)
}

func (s *UserService) Update(updatedUser models.User) error {
	if updatedUser.ID <= 0 {
		return errors.New("the ID is required")
	}

	obj, err := s.Repo.FindByID(updatedUser.ID)
	if err != nil || obj.ID == 0 {
		return fmt.Errorf("Could not found User with ID %d.", updatedUser.ID)
	}

	return s.Repo.Update(updatedUser)
}

func (s *UserService) DeleteByID(id uint) error {
	obj, err := s.Repo.FindByID(id)
	if err != nil || obj.ID == 0 {
		return fmt.Errorf("Could not found User with ID %d.", id)
	}

	return s.Repo.DeleteById(id)
}
