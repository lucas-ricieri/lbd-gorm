package repository

import (
	"azevedoruan.github/lbd-gorm/internal/models"

	"gorm.io/gorm"
)

type PlaylistRepository struct {
	DB *gorm.DB
}

func (r *PlaylistRepository) FindByID(id uint) (models.Playlist, error) {
	var playlist models.Playlist
	err := r.DB.Preload("User").Find(&playlist, id).Error
	return playlist, err
}

func (r *PlaylistRepository) FindAllFromUserId(id uint) ([]models.Playlist, error) {
	var playlists []models.Playlist
	err := r.DB.Where("usuario_id = ?", id).Find(&playlists).Error
	return playlists, err
}

func (r *PlaylistRepository) AddNewForUser(playlist *models.Playlist) error {
	result := r.DB.Create(&playlist)
	return result.Error
}

func (r *PlaylistRepository) GetAllMusics(playlistId uint, userId uint) (models.Playlist, error) {
	var playlistWithMusics models.Playlist
	playlistWithMusics.PlaylistId = playlistId
	playlistWithMusics.UserId = userId
	err := r.DB.
		Preload("User").
		Preload("MusicPlaylists.Music").
		Find(&playlistWithMusics).Error

	return playlistWithMusics, err
}

func (r *PlaylistRepository) AddMusic(musicPlaylist models.MusicPlaylist) (models.MusicPlaylist, error) {
	// TODO...
	return models.MusicPlaylist{}, nil
}

func (r *PlaylistRepository) DeleteMusic(musicPlaylist models.MusicPlaylist) error {
	// TODO...
	return nil
}
