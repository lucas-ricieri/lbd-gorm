package repository

import (
	"azevedoruan.github/lbd-gorm/internal/models"

	"gorm.io/gorm"
)

type MusicRepository struct {
	DB *gorm.DB
}

func (r *MusicRepository) FindByID(id uint) (models.Music, error) {
	var music models.Music
	err := r.DB.Preload("Artist").Find(&music, id).Error
	return music, err
}
