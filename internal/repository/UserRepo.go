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

func (r *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *UserRepository) AddNew(newUser models.User) (models.User, error) {
	result := r.DB.Create(&newUser)
	return newUser, result.Error
}

func (r *UserRepository) Update(updatedUser models.User) error {
	err := r.DB.Save(&updatedUser).Error
	return err
}

func (r *UserRepository) DeleteById(id uint) error {
	err := r.DB.Delete(&models.User{}, id).Error
	return err
}
