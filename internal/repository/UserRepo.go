package repository

import (
	"azevedoruan.github/lbd-gorm/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) FindByID(id uint) (models.User, error) {
	var user models.User
	err := r.DB.First(&user, id).Error
	return user, err
}
