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
	err := r.DB.Preload("Artist").First(&music, id).Error
	return music, err
}

func (r *MusicRepository) FindAll() ([]models.Music, error) {
	var musics []models.Music
	err := r.DB.Preload("Artist").Find(&musics).Error
	return musics, err
}

func (r *MusicRepository) AddNew(newMusic *models.Music) error {
	result := r.DB.Create(&newMusic)
	return result.Error
}

func (r *MusicRepository) Update(updatedMusic models.Music) error {
	err := r.DB.Save(&updatedMusic).Error
	return err
}

func (r *MusicRepository) DeleteById(id uint) error {
	err := r.DB.Delete(&models.Music{}, id).Error
	return err
}
