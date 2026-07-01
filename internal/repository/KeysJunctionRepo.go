package repository

import (
	"azevedoruan.github/lbd-gorm/internal/models"

	"gorm.io/gorm"
)

type KeysJunctionRepository struct {
	DB *gorm.DB
}

type PlaylistMusicOrderResult struct {
	MusicTitle    string `gorm:"column:music_title"`
	PlaylistName  string `gorm:"column:playlist_name"`
	PlaylistOrder uint   `gorm:"column:playlist_order"`
}

type MusicPlaylistOwnerResult struct {
	Username     string `gorm:"column:username"`
	PlaylistName string `gorm:"column:playlist_name"`
	MusicTitle   string `gorm:"column:music_title"`
}

func (r *KeysJunctionRepository) FindMusicsOrderByPlaylistName(playlistName string) ([]PlaylistMusicOrderResult, error) {
	var musics []PlaylistMusicOrderResult
	err := r.DB.Model(&models.MusicPlaylist{}).
		Select(`
			musica.titulo AS music_title,
			playlist.nome AS playlist_name,
			musica_playlist.ordem_na_playlist AS playlist_order
		`).
		Joins("INNER JOIN musica ON musica.id = musica_playlist.musica_id").
		Joins(`
			INNER JOIN playlist
				ON playlist.playlist_id = musica_playlist.playlist_id
				AND playlist.usuario_id = musica_playlist.usuario_id
		`).
		Where("playlist.nome = ?", playlistName).
		Order("musica_playlist.ordem_na_playlist, musica.titulo").
		Scan(&musics).Error

	return musics, err
}

func (r *KeysJunctionRepository) FindPlaylistOwnersByMusicTitle(musicTitle string) ([]MusicPlaylistOwnerResult, error) {
	var owners []MusicPlaylistOwnerResult
	err := r.DB.Model(&models.Music{}).
		Distinct(`
			usuario.username AS username,
			playlist.nome AS playlist_name,
			musica.titulo AS music_title
		`).
		Joins("INNER JOIN musica_playlist ON musica_playlist.musica_id = musica.id").
		Joins(`
			INNER JOIN playlist
				ON playlist.playlist_id = musica_playlist.playlist_id
				AND playlist.usuario_id = musica_playlist.usuario_id
		`).
		Joins("INNER JOIN usuario ON usuario.id = playlist.usuario_id").
		Where("musica.titulo = ?", musicTitle).
		Order("usuario.username, playlist.nome").
		Scan(&owners).Error

	return owners, err
}
