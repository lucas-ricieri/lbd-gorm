package service

import (
	"errors"
	"fmt"

	"azevedoruan.github/lbd-gorm/internal/models"
	"azevedoruan.github/lbd-gorm/internal/repository"
)

type PlaylistService struct {
	Repo      *repository.PlaylistRepository
	UserRepo  *repository.UserRepository
	MusicRepo *repository.MusicRepository
}

func (s *PlaylistService) FindByID(id uint) (models.Playlist, error) {
	return s.Repo.FindByID(id)
}

func (s *PlaylistService) FindAllFromUserId(userId uint) ([]models.Playlist, error) {
	user, err := s.UserRepo.FindByID(userId)
	if err != nil || user.ID == 0 {
		return nil, fmt.Errorf("Could not found User with ID %d.", userId)
	}

	return s.Repo.FindAllFromUserId(userId)
}

func (s *PlaylistService) Create(userId uint, playlist *models.Playlist) error {
	playlist.UserId = userId

	user, err := s.UserRepo.FindByID(userId)
	if err != nil || user.ID == 0 {
		return fmt.Errorf("Could not found User with ID %d.", userId)
	}

	return s.Repo.AddNewForUser(playlist)
}

func (s *PlaylistService) GetById(playlistId uint, userId uint) (models.Playlist, error) {
	playlist, err := s.Repo.FindByID(playlistId)
	if err != nil || playlist.PlaylistId == 0 {
		return models.Playlist{}, fmt.Errorf("Could not found Playlist with ID %d.", playlistId)
	}

	user, err := s.UserRepo.FindByID(userId)
	if err != nil || user.ID == 0 {
		return models.Playlist{}, fmt.Errorf("Could not found User with ID %d.", userId)
	}

	return s.Repo.GetAllMusics(playlistId, userId)
}

func (s *PlaylistService) Update(playlist *models.Playlist) error {
	if playlist.PlaylistId <= 0 {
		return errors.New("the Playlist ID is required")
	}
	if playlist.UserId <= 0 {
		return errors.New("the User ID is required")
	}

	obj, err := s.Repo.FindByID(playlist.PlaylistId)
	if err != nil || obj.PlaylistId == 0 {
		return fmt.Errorf("Could not found Playlist with ID %d.", playlist.PlaylistId)
	}

	user, err := s.UserRepo.FindByID(playlist.UserId)
	if err != nil || user.ID == 0 {
		return fmt.Errorf("Could not found User with ID %d.", playlist.UserId)
	}

	// Não se atualiza data de criacao
	playlist.CreactionDate = obj.CreactionDate

	return s.Repo.Update(playlist)
}

func (s *PlaylistService) AddMusic(newMusic *models.MusicPlaylist) error {
	if newMusic.MusicId <= 0 {
		return errors.New("the Music ID is required")
	}
	if newMusic.PlaylistId <= 0 {
		return errors.New("the Playlist ID is required")
	}
	if newMusic.UserId <= 0 {
		return errors.New("the User ID is required")
	}

	music, err := s.MusicRepo.FindByID(newMusic.MusicId)
	if err != nil || music.ID == 0 {
		return fmt.Errorf("Could not found Music with ID %d.", newMusic.MusicId)
	}

	playlist, err := s.Repo.FindByID(newMusic.PlaylistId)
	if err != nil || playlist.PlaylistId == 0 {
		return fmt.Errorf("Could not found Playlist with ID %d.", newMusic.PlaylistId)
	}

	user, err := s.UserRepo.FindByID(newMusic.UserId)
	if err != nil || user.ID == 0 {
		return fmt.Errorf("Could not found User with ID %d.", newMusic.UserId)
	}

	obj, err := s.Repo.GetLastSortedMusic(newMusic.PlaylistId, newMusic.UserId)
	if err != nil {
		return fmt.Errorf("Error to get the last music in the playlist: %w", err)
	}
	newMusic.PlaylistOrder = obj.PlaylistOrder + 1

	return s.Repo.AddMusic(newMusic)
}

func (s *PlaylistService) RemoveMusic(musicToRemove models.MusicPlaylist) error {
	if musicToRemove.MusicId <= 0 {
		return errors.New("the Music ID is required")
	}
	if musicToRemove.PlaylistId <= 0 {
		return errors.New("the Playlist ID is required")
	}
	if musicToRemove.UserId <= 0 {
		return errors.New("the User ID is required")
	}

	music, err := s.MusicRepo.FindByID(musicToRemove.MusicId)
	if err != nil || music.ID == 0 {
		return fmt.Errorf("Could not found Music with ID %d.", musicToRemove.MusicId)
	}

	playlist, err := s.Repo.FindByID(musicToRemove.PlaylistId)
	if err != nil || playlist.PlaylistId == 0 {
		return fmt.Errorf("Could not found Playlist with ID %d.", musicToRemove.PlaylistId)
	}

	user, err := s.UserRepo.FindByID(musicToRemove.UserId)
	if err != nil || user.ID == 0 {
		return fmt.Errorf("Could not found User with ID %d.", musicToRemove.UserId)
	}

	obj, err := s.Repo.GetLastSortedMusic(musicToRemove.PlaylistId, musicToRemove.UserId)
	if err != nil {
		return fmt.Errorf("Error to get the last music in the playlist: %w", err)
	}
	musicToRemove.PlaylistOrder = obj.PlaylistOrder + 1

	return s.Repo.RemoveMusic(musicToRemove)
}

func (s *PlaylistService) Delete(playlist *models.Playlist) error {
	if playlist.PlaylistId <= 0 {
		return errors.New("the Playlist ID is required")
	}
	if playlist.UserId <= 0 {
		return errors.New("the User ID is required")
	}

	obj, err := s.Repo.FindByID(playlist.PlaylistId)
	if err != nil || obj.PlaylistId == 0 {
		return fmt.Errorf("Could not found Playlist with ID %d.", playlist.PlaylistId)
	}

	user, err := s.UserRepo.FindByID(playlist.UserId)
	if err != nil || user.ID == 0 {
		return fmt.Errorf("Could not found User with ID %d.", playlist.UserId)
	}

	return s.Repo.Delete(playlist)
}
