package service

import "azevedoruan.github/lbd-gorm/internal/repository"

type KeysJunctionService struct {
	Repo *repository.KeysJunctionRepository
}

func (s *KeysJunctionService) FindMusicsOrderByPlaylistName(playlistName string) ([]repository.PlaylistMusicOrderResult, error) {
	return s.Repo.FindMusicsOrderByPlaylistName(playlistName)
}

func (s *KeysJunctionService) FindPlaylistOwnersByMusicTitle(musicTitle string) ([]repository.MusicPlaylistOwnerResult, error) {
	return s.Repo.FindPlaylistOwnersByMusicTitle(musicTitle)
}
