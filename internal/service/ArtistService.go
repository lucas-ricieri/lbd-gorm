package service

import (
	"errors"
	"fmt"

	"azevedoruan.github/lbd-gorm/internal/models"
	"azevedoruan.github/lbd-gorm/internal/repository"
)

type ArtistService struct {
	Repo *repository.ArtistRepository
}

func (s *ArtistService) FindByID(id uint) (models.Artist, error) {
	return s.Repo.FindByID(id)
}

func (s *ArtistService) FindAll() ([]models.Artist, error) {
	return s.Repo.FindAll()
}

func (s *ArtistService) Create(newArtist *models.Artist) error {
	return s.Repo.AddNew(newArtist)
}

func (s *ArtistService) Update(updatedArtist models.Artist) error {
	if updatedArtist.ID <= 0 {
		return errors.New("the ID is required")
	}

	obj, err := s.Repo.FindByID(updatedArtist.ID)
	if err != nil || obj.ID == 0 {
		return fmt.Errorf("Could not found User with ID %d.", updatedArtist.ID)
	}

	return s.Repo.Update(updatedArtist)
}

func (s *ArtistService) DeleteByID(id uint) error {
	obj, err := s.Repo.FindByID(id)
	if err != nil || obj.ID == 0 {
		return fmt.Errorf("Could not found User with ID %d.", id)
	}

	return s.Repo.DeleteById(id)
}
