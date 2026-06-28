package service

import "azevedoruan.github/lbd-gorm/internal/repository"

type RelantionshipService struct {
	Repo *repository.RelantionshipRepository
}

func (s *RelantionshipService) FindPlaylistsByUsername(username string) ([]repository.UserPlaylistResult, error) {
	return s.Repo.FindPlaylistsByUsername(username)
}

func (s *RelantionshipService) FindMusicsInUserPlaylistsByArtist(username string, artistName string) ([]repository.UserPlaylistArtistMusicResult, error) {
	return s.Repo.FindMusicsInUserPlaylistsByArtist(username, artistName)
}

func (s *RelantionshipService) CountMusicsByPlaylist() ([]repository.PlaylistMusicCountResult, error) {
	return s.Repo.CountMusicsByPlaylist()
}

func (s *RelantionshipService) FindArtistsWithoutMusicsInPlaylists() ([]repository.ArtistWithoutPlaylistMusicResult, error) {
	return s.Repo.FindArtistsWithoutMusicsInPlaylists()
}
