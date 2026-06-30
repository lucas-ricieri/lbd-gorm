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
	err := r.DB.Raw(`
		SELECT
			p.nome AS playlist_name,
			u.username AS username,
			COALESCE(SUM(m.duracao_segundos), 0) AS total_duration_seconds
		FROM playlist p
		INNER JOIN usuario u ON u.id = p.usuario_id
		LEFT JOIN musica_playlist mp
			ON mp.playlist_id = p.playlist_id
			AND mp.usuario_id = p.usuario_id
		LEFT JOIN musica m ON m.id = mp.musica_id
		GROUP BY p.playlist_id, p.usuario_id, p.nome, u.username
		ORDER BY total_duration_seconds DESC, p.nome
	`).Scan(&playlists).Error

	return playlists, err
}

func (r *OtimizationDetailRepository) FindMusicsShorterThanArtistAverage() ([]MusicShorterThanArtistAverageResult, error) {
	var musics []MusicShorterThanArtistAverageResult
	err := r.DB.Raw(`
		SELECT
			m.id AS music_id,
			m.titulo AS music_title,
			m.duracao_segundos AS duration_seconds,
			a.nome AS artist_name,
			artist_average.average_seconds AS artist_average_seconds
		FROM musica m
		INNER JOIN artista a ON a.id = m.artista_id
		INNER JOIN (
			SELECT
				artista_id,
				AVG(duracao_segundos) AS average_seconds
			FROM musica
			GROUP BY artista_id
		) artist_average ON artist_average.artista_id = m.artista_id
		WHERE m.duracao_segundos < artist_average.average_seconds
		ORDER BY a.nome, m.duracao_segundos, m.titulo
	`).Scan(&musics).Error

	return musics, err
}
