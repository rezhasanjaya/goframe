package services

import (
	"errors"
	"time"

	"goframe/internal/app/models"
	"goframe/internal/core/bootstrap"
	"goframe/internal/core/config"
	"goframe/internal/core/jwt"
	"goframe/internal/core/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	db  *gorm.DB
	cfg *config.Config
}

func NewAuthService(cfg *config.Config) *AuthService {
	return &AuthService{
		db:  bootstrap.DB,
		cfg: cfg,
	}
}

func (s *AuthService) Register(u *models.User) error {
	var exists int64
	s.db.Model(&models.User{}).Where("email = ?", u.Email).Count(&exists)
	if exists > 0 {
		return errors.New("email already taken")
	}

	if u.Password == "" {
		return errors.New("password required")
	}
	return s.db.Create(u).Error
}

func (s *AuthService) Login(email, password string) (accessToken string, refreshToken string, refreshExpiry time.Time, err error) {
	var user models.User
	if err = s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return "", "", time.Time{}, err
	}
	if !user.ComparePassword(password) {
		return "", "", time.Time{}, errors.New("invalid credentials")
	}

	accessToken, err = jwt.GenerateAccessToken(s.cfg.JWTSecret, s.cfg.AccessTokenTTLMin, user.UUID, user.Email)
	if err != nil {
		return "", "", time.Time{}, err
	}

	refreshToken, err = utils.RandomToken(32)
	if err != nil {
		return "", "", time.Time{}, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)
	if err != nil {
		return "", "", time.Time{}, err
	}

	expiry := time.Now().Add(time.Duration(s.cfg.RefreshTokenTTLH) * time.Hour)

	if err := s.db.Model(&user).Updates(map[string]interface{}{
		"refresh_token_hash": string(hash),
		"updated_at":         time.Now(),
	}).Error; err != nil {
		return "", "", time.Time{}, err
	}

	return accessToken, refreshToken, expiry, nil
}

func (s *AuthService) Refresh(email, providedRefresh string) (newAccessToken, newRefreshToken string, newExpiry time.Time, err error) {
	var user models.User
	if err = s.db.Where("email = ?", email).First(&user).Error; err != nil {
		return "", "", time.Time{}, err
	}
	if user.RefreshTokenHash == "" {
		return "", "", time.Time{}, errors.New("no refresh token stored")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.RefreshTokenHash), []byte(providedRefresh)) != nil {
		return "", "", time.Time{}, errors.New("invalid refresh token")
	}

	newAccessToken, err = jwt.GenerateAccessToken(s.cfg.JWTSecret, s.cfg.AccessTokenTTLMin, user.UUID, user.Email)
	if err != nil {
		return "", "", time.Time{}, err
	}

	newRefreshToken, err = utils.RandomToken(32)
	if err != nil {
		return "", "", time.Time{}, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(newRefreshToken), bcrypt.DefaultCost)
	if err != nil {
		return "", "", time.Time{}, err
	}

	newExpiry = time.Now().Add(time.Duration(s.cfg.RefreshTokenTTLH) * time.Hour)
	if err := s.db.Model(&user).Updates(map[string]interface{}{
		"refresh_token_hash": string(hash),
		"updated_at":         time.Now(),
	}).Error; err != nil {
		return "", "", time.Time{}, err
	}

	return newAccessToken, newRefreshToken, newExpiry, nil
}

func (s *AuthService) Logout(email string) error {
	return s.db.Model(&models.User{}).Where("email = ?", email).Updates(map[string]interface{}{
		"refresh_token_hash": nil,
		"updated_at":         time.Now(),
	}).Error
}
