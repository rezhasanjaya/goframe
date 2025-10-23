package api

import (
	"goframe/internal/app/http/controllers"
	"goframe/internal/app/models"
	"goframe/internal/app/services"
	"goframe/internal/core/config"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	controllers.BaseController
	service *services.AuthService
	cfg     *config.Config
}

func NewAuthController(cfg *config.Config) *AuthController {
	return &AuthController{
		service: services.NewAuthService(cfg),
		cfg:     cfg,
	}
}

type registerReq struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (ac *AuthController) Register(c *gin.Context) {
	var req registerReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ac.ValidationError(c, err.Error())
		return
	}
	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	if err := ac.service.Register(user); err != nil {
		ac.Error(c, "Failed to register", err.Error())
		return
	}
	ac.Created(c, "User registered", gin.H{"uuid": user.UUID, "email": user.Email})
}

type loginReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (ac *AuthController) Login(c *gin.Context) {
	var req loginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ac.ValidationError(c, err.Error())
		return
	}
	access, refresh, expiry, err := ac.service.Login(req.Email, req.Password)
	if err != nil {
		ac.Error(c, "Failed to login", err.Error())
		return
	}
	ac.Success(c, "Logged in", gin.H{
		"access_token":  access,
		"refresh_token": refresh,
		"expires_at":    expiry,
	})
}

type refreshReq struct {
	Email        string `json:"email" binding:"required,email"`
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func (ac *AuthController) Refresh(c *gin.Context) {
	var req refreshReq
	if err := c.ShouldBindJSON(&req); err != nil {
		ac.ValidationError(c, err.Error())
		return
	}
	access, refresh, expiry, err := ac.service.Refresh(req.Email, req.RefreshToken)
	if err != nil {
		ac.Error(c, "Failed to refresh token", err.Error())
		return
	}
	ac.Success(c, "Token refreshed", gin.H{
		"access_token":  access,
		"refresh_token": refresh,
		"expires_at":    expiry,
	})
}

func (ac *AuthController) Logout(c *gin.Context) {
	// read email from body or from jwt claims via middleware. here expect JSON
	var body struct {
		Email string `json:"email" binding:"required,email"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		ac.ValidationError(c, err.Error())
		return
	}
	if err := ac.service.Logout(body.Email); err != nil {
		ac.Error(c, "Logout failed", err.Error())
		return
	}
	ac.Success(c, "Logged out", nil)
}
