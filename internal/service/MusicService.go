package service

import (
	"errors"
	"fmt"

	"azevedoruan.github/lbd-gorm/internal/models"
	"azevedoruan.github/lbd-gorm/internal/repository"
)

type MusicService struct {
	Repo       *repository.MusicRepository
	ArtistRepo *repository.ArtistRepository
}

func (s *MusicService) FindByID(id uint) (models.Music, error) {
	return s.Repo.FindByID(id)
}

func (s *MusicService) FindAll() ([]models.Music, error) {
	return s.Repo.FindAll()
}

func (s *MusicService) Create(newMusic *models.Music) error {
	artist, err := s.ArtistRepo.FindByID(uint(newMusic.ArtistId))
	if err != nil || artist.ID == 0 {
		return fmt.Errorf("Could not found Artist with ID %d.", newMusic.ArtistId)
	}

	return s.Repo.AddNew(newMusic)
}

func (s *MusicService) Update(updatedMusic models.Music) error {
	if updatedMusic.ID <= 0 {
		return errors.New("the ID is required")
	}
	if updatedMusic.ArtistId <= 0 {
		return errors.New("the Artist ID is required")
	}

	music, err := s.Repo.FindByID(uint(updatedMusic.ID))
	if err != nil || music.ID == 0 {
		return fmt.Errorf("Could not found Music with ID %d.", updatedMusic.ID)
	}

	artist, err := s.ArtistRepo.FindByID(uint(updatedMusic.ArtistId))
	if err != nil || artist.ID == 0 {
		return fmt.Errorf("Could not found Artist with ID %d.", updatedMusic.ArtistId)
	}

	return s.Repo.Update(updatedMusic)
}

func (s *MusicService) DeleteByID(id uint) error {
	music, err := s.Repo.FindByID(id)
	if err != nil || music.ID == 0 {
		return fmt.Errorf("Could not found Music with ID %d.", id)
	}

	return s.Repo.DeleteById(id)
}
