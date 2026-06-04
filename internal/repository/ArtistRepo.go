package repository

import (
	"azevedoruan.github/lbd-gorm/internal/models"

	"gorm.io/gorm"
)

type ArtistRepository struct {
	DB *gorm.DB
}

func (r *ArtistRepository) FindByID(id uint) (models.Artist, error) {
	var artist models.Artist
	err := r.DB.First(&artist, id).Error
	return artist, err
}

func (r *ArtistRepository) FindAll() ([]models.Artist, error) {
	var artists []models.Artist
	err := r.DB.Find(&artists).Error
	return artists, err
}
