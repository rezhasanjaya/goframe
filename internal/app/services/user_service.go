package services

import (
	"goframe/internal/app/models"
	"goframe/internal/core/bootstrap"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService() *UserService {
	return &UserService{db: bootstrap.DB}
}

func (s *UserService) GetAll() ([]models.User, error) {
	var users []models.User
	err := s.db.Find(&users).Error
	return users, err
}

func (s *UserService) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := s.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) Create(user *models.User) error {
	return s.db.Create(user).Error
}

func (s *UserService) Update(id uint, updateData map[string]interface{}) error {
	return s.db.Model(&models.User{}).Where("id = ?", id).Updates(updateData).Error
}

func (s *UserService) Delete(id uint) error {
	return s.db.Delete(&models.User{}, id).Error
}
