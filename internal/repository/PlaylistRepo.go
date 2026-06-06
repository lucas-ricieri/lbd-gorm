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
