package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID              uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	UUID            string         `json:"uuid" gorm:"type:varchar(255);uniqueIndex"`
	Name            string         `json:"name"`
	Email           string         `json:"email" gorm:"uniqueIndex"`
	Password        string         `json:"-"` 
    RefreshTokenHash string        `json:"-"`
	Token           *string        `json:"token,omitempty"`
	APIKey          *string        `json:"api_key,omitempty"`
	APIKeyHash      *string        `json:"api_key_hash,omitempty"`
	EmailVerifiedAt *time.Time     `json:"email_verified_at,omitempty"`
	RememberToken   *string        `json:"remember_token,omitempty"`
	FlagActive      int            `json:"flag_active" gorm:"default:1"`
	Level           int            `json:"level" gorm:"default:1"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.UUID == "" {
		u.UUID = uuid.NewString()
	}
	if u.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashed)
	}
	return
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}
