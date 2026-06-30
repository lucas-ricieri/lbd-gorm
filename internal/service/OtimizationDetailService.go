package service

import (
	"azevedoruan.github/lbd-gorm/internal/models"
	"azevedoruan.github/lbd-gorm/internal/repository"
)

type OtimizationDetailService struct {
	Repo *repository.OtimizationDetailRepository
}

func (s *OtimizationDetailService) FindMusicWithArtistById(id uint) (models.Music, error) {
	return s.Repo.FindMusicWithArtistById(id)
}

func (s *OtimizationDetailService) TotalDurationByPlaylist() ([]repository.PlaylistTotalDurationResult, error) {
	return s.Repo.TotalDurationByPlaylist()
}

func (s *OtimizationDetailService) FindMusicsShorterThanArtistAverage() ([]repository.MusicShorterThanArtistAverageResult, error) {
	return s.Repo.FindMusicsShorterThanArtistAverage()
}
