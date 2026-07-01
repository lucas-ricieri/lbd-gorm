package service

import (
	"errors"

	"azevedoruan.github/lbd-gorm/internal/repository"
)

type TransactionService struct {
	Repo *repository.TransactionRepository
}

func (s *TransactionService) TransferMusicBetweenPlaylists(request repository.MusicTransferRequest) error {
	if request.MusicId <= 0 {
		return errors.New("the Music ID is required")
	}
	if request.FromPlaylistId <= 0 {
		return errors.New("the source Playlist ID is required")
	}
	if request.ToPlaylistId <= 0 {
		return errors.New("the destination Playlist ID is required")
	}
	if request.UserId <= 0 {
		return errors.New("the User ID is required")
	}
	if request.FromPlaylistId == request.ToPlaylistId {
		return errors.New("source and destination playlists must be different")
	}

	return s.Repo.TransferMusicBetweenPlaylists(request)
}
