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
