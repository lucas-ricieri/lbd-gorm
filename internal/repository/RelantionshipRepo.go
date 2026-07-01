package repository

import (
	"time"

	"azevedoruan.github/lbd-gorm/internal/models"

	"gorm.io/gorm"
)

type RelantionshipRepository struct {
	DB *gorm.DB
}

type UserPlaylistResult struct {
	PlaylistName string    `gorm:"column:playlist_name"`
	CreationDate time.Time `gorm:"column:creation_date"`
}

type UserPlaylistArtistMusicResult struct {
	MusicTitle   string `gorm:"column:music_title"`
	PlaylistName string `gorm:"column:playlist_name"`
	ArtistName   string `gorm:"column:artist_name"`
}

type PlaylistMusicCountResult struct {
	PlaylistName string `gorm:"column:playlist_name"`
	MusicCount   int64  `gorm:"column:music_count"`
}

type ArtistWithoutPlaylistMusicResult struct {
	ArtistId   uint   `gorm:"column:artist_id"`
	ArtistName string `gorm:"column:artist_name"`
}

func (r *RelantionshipRepository) FindPlaylistsByUsername(username string) ([]UserPlaylistResult, error) {
	var playlists []UserPlaylistResult
	err := r.DB.Model(&models.Playlist{}).
		Select(`
			playlist.nome AS playlist_name,
			playlist.data_criacao AS creation_date
		`).
		Joins("INNER JOIN usuario ON usuario.id = playlist.usuario_id").
		Where("usuario.username = ?", username).
		Order("playlist.data_criacao DESC").
		Scan(&playlists).Error

	return playlists, err
}

func (r *RelantionshipRepository) FindMusicsInUserPlaylistsByArtist(username string, artistName string) ([]UserPlaylistArtistMusicResult, error) {
	var musics []UserPlaylistArtistMusicResult
	err := r.DB.Model(&models.Music{}).
		Distinct(`
			musica.titulo AS music_title,
			playlist.nome AS playlist_name,
			artista.nome AS artist_name
		`).
		Joins("INNER JOIN artista ON artista.id = musica.artista_id").
		Joins("INNER JOIN musica_playlist ON musica_playlist.musica_id = musica.id").
		Joins(`
			INNER JOIN playlist
				ON playlist.playlist_id = musica_playlist.playlist_id
				AND playlist.usuario_id = musica_playlist.usuario_id
		`).
		Joins("INNER JOIN usuario ON usuario.id = playlist.usuario_id").
		Where("usuario.username = ?", username).
		Where("artista.nome = ?", artistName).
		Order("playlist.nome, musica.titulo").
		Scan(&musics).Error

	return musics, err
}

func (r *RelantionshipRepository) CountMusicsByPlaylist() ([]PlaylistMusicCountResult, error) {
	var playlists []PlaylistMusicCountResult
	err := r.DB.Model(&models.Playlist{}).
		Select(`
			playlist.nome AS playlist_name,
			COUNT(musica_playlist.musica_id) AS music_count
		`).
		Joins(`
			LEFT JOIN musica_playlist
				ON musica_playlist.playlist_id = playlist.playlist_id
				AND musica_playlist.usuario_id = playlist.usuario_id
		`).
		Group("playlist.playlist_id, playlist.usuario_id, playlist.nome").
		Order("music_count DESC, playlist.nome").
		Scan(&playlists).Error

	return playlists, err
}

func (r *RelantionshipRepository) FindArtistsWithoutMusicsInPlaylists() ([]ArtistWithoutPlaylistMusicResult, error) {
	var artists []ArtistWithoutPlaylistMusicResult
	err := r.DB.Model(&models.Artist{}).
		Select(`
			artista.id AS artist_id,
			artista.nome AS artist_name
		`).
		Joins("LEFT JOIN musica ON musica.artista_id = artista.id").
		Joins("LEFT JOIN musica_playlist ON musica_playlist.musica_id = musica.id").
		Group("artista.id, artista.nome").
		Having("COUNT(musica_playlist.musica_id) = 0").
		Order("artista.nome").
		Scan(&artists).Error

	return artists, err
}
