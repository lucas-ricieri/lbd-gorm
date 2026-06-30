package repository

import (
	"time"

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
	err := r.DB.Raw(`
		SELECT
			p.nome AS playlist_name,
			p.data_criacao AS creation_date
		FROM playlist p
		INNER JOIN usuario u ON u.id = p.usuario_id
		WHERE u.username = ?
		ORDER BY p.data_criacao DESC
	`, username).Scan(&playlists).Error

	return playlists, err
}

func (r *RelantionshipRepository) FindMusicsInUserPlaylistsByArtist(username string, artistName string) ([]UserPlaylistArtistMusicResult, error) {
	var musics []UserPlaylistArtistMusicResult
	err := r.DB.Raw(`
		SELECT DISTINCT
			m.titulo AS music_title,
			p.nome AS playlist_name,
			a.nome AS artist_name
		FROM musica m
		INNER JOIN artista a ON a.id = m.artista_id
		INNER JOIN musica_playlist mp ON mp.musica_id = m.id
		INNER JOIN playlist p
			ON p.playlist_id = mp.playlist_id
			AND p.usuario_id = mp.usuario_id
		INNER JOIN usuario u ON u.id = p.usuario_id
		WHERE u.username = ?
			AND a.nome = ?
		ORDER BY p.nome, m.titulo
	`, username, artistName).Scan(&musics).Error

	return musics, err
}

func (r *RelantionshipRepository) CountMusicsByPlaylist() ([]PlaylistMusicCountResult, error) {
	var playlists []PlaylistMusicCountResult
	err := r.DB.Raw(`
		SELECT
			p.nome AS playlist_name,
			COUNT(mp.musica_id) AS music_count
		FROM playlist p
		LEFT JOIN musica_playlist mp
			ON mp.playlist_id = p.playlist_id
			AND mp.usuario_id = p.usuario_id
		GROUP BY p.playlist_id, p.usuario_id, p.nome
		ORDER BY music_count DESC, p.nome
	`).Scan(&playlists).Error

	return playlists, err
}

func (r *RelantionshipRepository) FindArtistsWithoutMusicsInPlaylists() ([]ArtistWithoutPlaylistMusicResult, error) {
	var artists []ArtistWithoutPlaylistMusicResult
	err := r.DB.Raw(`
		SELECT
			a.id AS artist_id,
			a.nome AS artist_name
		FROM artista a
		WHERE NOT EXISTS (
			SELECT 1
			FROM musica m
			INNER JOIN musica_playlist mp ON mp.musica_id = m.id
			WHERE m.artista_id = a.id
		)
		ORDER BY a.nome
	`).Scan(&artists).Error

	return artists, err
}
