package repository

import (
	"azevedoruan.github/lbd-gorm/internal/models"

	"gorm.io/gorm"
)

type OtimizationDetailRepository struct {
	DB *gorm.DB
}

type PlaylistTotalDurationResult struct {
	PlaylistName        string `gorm:"column:playlist_name"`
	Username            string `gorm:"column:username"`
	TotalDurationSecond int64  `gorm:"column:total_duration_seconds"`
}

type MusicShorterThanArtistAverageResult struct {
	MusicId             uint    `gorm:"column:music_id"`
	MusicTitle          string  `gorm:"column:music_title"`
	DurationSec         int32   `gorm:"column:duration_seconds"`
	ArtistName          string  `gorm:"column:artist_name"`
	ArtistAverageSecond float64 `gorm:"column:artist_average_seconds"`
}

func (r *OtimizationDetailRepository) FindMusicWithArtistById(id uint) (models.Music, error) {
	var music models.Music
	err := r.DB.Joins("Artist").First(&music, id).Error
	return music, err
}

func (r *OtimizationDetailRepository) TotalDurationByPlaylist() ([]PlaylistTotalDurationResult, error) {
	var playlists []PlaylistTotalDurationResult
	err := r.DB.Model(&models.Playlist{}).
		Select(`
			playlist.nome AS playlist_name,
			usuario.username AS username,
			COALESCE(SUM(musica.duracao_segundos), 0) AS total_duration_seconds
		`).
		Joins("INNER JOIN usuario ON usuario.id = playlist.usuario_id").
		Joins(`
			LEFT JOIN musica_playlist
				ON musica_playlist.playlist_id = playlist.playlist_id
				AND musica_playlist.usuario_id = playlist.usuario_id
		`).
		Joins("LEFT JOIN musica ON musica.id = musica_playlist.musica_id").
		Group("playlist.playlist_id, playlist.usuario_id, playlist.nome, usuario.username").
		Order("total_duration_seconds DESC, playlist.nome").
		Scan(&playlists).Error

	return playlists, err
}

func (r *OtimizationDetailRepository) FindMusicsShorterThanArtistAverage() ([]MusicShorterThanArtistAverageResult, error) {
	var musics []MusicShorterThanArtistAverageResult
	artistAverage := r.DB.Model(&models.Music{}).
		Select("artista_id, AVG(duracao_segundos) AS average_seconds").
		Group("artista_id")

	err := r.DB.Model(&models.Music{}).
		Select(`
			musica.id AS music_id,
			musica.titulo AS music_title,
			musica.duracao_segundos AS duration_seconds,
			artista.nome AS artist_name,
			artist_average.average_seconds AS artist_average_seconds
		`).
		Joins("INNER JOIN artista ON artista.id = musica.artista_id").
		Joins("INNER JOIN (?) artist_average ON artist_average.artista_id = musica.artista_id", artistAverage).
		Where("musica.duracao_segundos < artist_average.average_seconds").
		Order("artista.nome, musica.duracao_segundos, musica.titulo").
		Scan(&musics).Error

	return musics, err
}
