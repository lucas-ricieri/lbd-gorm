package repository

import (
	"azevedoruan.github/lbd-gorm/internal/models"

	"gorm.io/gorm"
)

type FunctionRepository struct {
	DB *gorm.DB
}

type ArtistPopularityRankResult struct {
	ArtistId       uint   `gorm:"column:artist_id"`
	ArtistName     string `gorm:"column:artist_name"`
	PlaylistCount  int64  `gorm:"column:playlist_count"`
	PopularityRank int64  `gorm:"column:popularity_rank"`
}

type MusicLongerThanArtistMaxResult struct {
	MusicId          uint   `gorm:"column:music_id"`
	MusicTitle       string `gorm:"column:music_title"`
	DurationSec      int32  `gorm:"column:duration_seconds"`
	ArtistName       string `gorm:"column:artist_name"`
	ComparedToArtist string `gorm:"column:compared_to_artist"`
}

func (r *FunctionRepository) RankArtistsByPlaylistPopularity() ([]ArtistPopularityRankResult, error) {
	var artists []ArtistPopularityRankResult
	err := r.DB.Model(&models.Artist{}).
		Select(`
			artista.id AS artist_id,
			artista.nome AS artist_name,
			COUNT(DISTINCT (musica_playlist.playlist_id, musica_playlist.usuario_id)) AS playlist_count,
			RANK() OVER (
				ORDER BY COUNT(DISTINCT (musica_playlist.playlist_id, musica_playlist.usuario_id)) DESC
			) AS popularity_rank
		`).
		Joins("LEFT JOIN musica ON musica.artista_id = artista.id").
		Joins("LEFT JOIN musica_playlist ON musica_playlist.musica_id = musica.id").
		Group("artista.id, artista.nome").
		Order("popularity_rank, artista.nome").
		Scan(&artists).Error

	return artists, err
}

func (r *FunctionRepository) FindMusicsLongerThanArtistMax(artistName string, comparisonArtistName string) ([]MusicLongerThanArtistMaxResult, error) {
	var musics []MusicLongerThanArtistMaxResult
	comparisonArtistId := r.DB.Model(&models.Artist{}).
		Select("id").
		Where("nome = ?", comparisonArtistName).
		Limit(1)
	maxComparisonDuration := r.DB.Model(&models.Music{}).
		Select("MAX(duracao_segundos)").
		Where("artista_id = (?)", comparisonArtistId)

	err := r.DB.Model(&models.Music{}).
		Select(`
			musica.id AS music_id,
			musica.titulo AS music_title,
			musica.duracao_segundos AS duration_seconds,
			artista.nome AS artist_name,
			? AS compared_to_artist
		`, comparisonArtistName).
		Joins("INNER JOIN artista ON artista.id = musica.artista_id").
		Where("artista.nome = ?", artistName).
		Where("musica.duracao_segundos > (?)", maxComparisonDuration).
		Order("musica.duracao_segundos DESC, musica.titulo").
		Scan(&musics).Error

	return musics, err
}
