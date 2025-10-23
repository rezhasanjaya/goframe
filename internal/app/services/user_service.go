package services

import (
	"errors"

	"goframe/internal/app/models"
	"goframe/internal/core/bootstrap"

	"gorm.io/gorm"
)

// ---------------- Response Structs ----------------
type UserResponse struct {
	UUID       string    `json:"uuid"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Level      int       `json:"level"`
	FlagActive int       `json:"flag_active"`
}

type UserCreateResponse struct {
	UUID  string `json:"uuid"`
	Email string `json:"email"`
}

type UserManageResponse struct {
	UUID  string `json:"uuid"`
	Email string `json:"email"`
}


// ---------------- Service ----------------
type UserService struct {
	db *gorm.DB
}

func NewUserService() *UserService {
	return &UserService{db: bootstrap.DB}
}

func sanitizeUser(user *models.User) *UserResponse {
	return &UserResponse{
		UUID:       user.UUID,
		Name:       user.Name,
		Email:      user.Email,
		Level:      user.Level,
		FlagActive: user.FlagActive,
	}
}

func (s *UserService) Fetch() ([]*UserResponse, error) {
	var users []models.User
	if err := s.db.Find(&users).Error; err != nil {
		return nil, err
	}

	resp := make([]*UserResponse, 0, len(users))
	for _, u := range users {
		resp = append(resp, sanitizeUser(&u))
	}
	return resp, nil
}

func (s *UserService) Get(uuid string) (*UserResponse, error) {
	var user models.User
	if err := s.db.Where("uuid = ?", uuid).First(&user).Error; err != nil {
		return nil, err
	}
	return sanitizeUser(&user), nil
}

func (s *UserService) Create(user *models.User) (*UserCreateResponse, error) {
	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}

	return &UserCreateResponse{
		UUID:  user.UUID,
		Email: user.Email,
	}, nil
}

func (s *UserService) Update(uuid string, updateData map[string]interface{}) (*UserManageResponse, error) {
	var user models.User
	if err := s.db.Where("uuid = ?", uuid).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	if err := s.db.Model(&user).Updates(updateData).Error; err != nil {
		return nil, err
	}

	return &UserManageResponse{
		UUID:  user.UUID,
		Email: user.Email,
	}, nil
}

func (s *UserService) Delete(uuid string) (*UserManageResponse, error) {
	var user models.User
	if err := s.db.Where("uuid = ?", uuid).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	if err := s.db.Delete(&user).Error; err != nil {
		return nil, err
	}

	return &UserManageResponse{
		UUID:  user.UUID,
		Email: user.Email,
	}, nil
}