package service

import "azevedoruan.github/lbd-gorm/internal/repository"

type FunctionService struct {
	Repo *repository.FunctionRepository
}

func (s *FunctionService) RankArtistsByPlaylistPopularity() ([]repository.ArtistPopularityRankResult, error) {
	return s.Repo.RankArtistsByPlaylistPopularity()
}

func (s *FunctionService) FindMusicsLongerThanArtistMax(artistName string, comparisonArtistName string) ([]repository.MusicLongerThanArtistMaxResult, error) {
	return s.Repo.FindMusicsLongerThanArtistMax(artistName, comparisonArtistName)
}
