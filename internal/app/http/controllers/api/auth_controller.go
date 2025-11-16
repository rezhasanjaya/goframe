package api

import (
	"goframe/internal/app/http/controllers"
	"goframe/internal/app/models"
	"goframe/internal/app/services"
	"goframe/internal/app/validators"
	"goframe/internal/core/config"
	"goframe/internal/core/validation"

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

func (ac *AuthController) Register(c *gin.Context) {
	var req validators.RegisterValidator

	if errs, err := validation.BindAndValidate(c, &req); err != nil {
		ac.ValidationError(c, errs)
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

	ac.Created(c, "User registered", gin.H{
		"uuid":  user.UUID,
		"email": user.Email,
	})
}

func (ac *AuthController) Login(c *gin.Context) {
	var req validators.LoginValidator

	if errs, err := validation.BindAndValidate(c, &req); err != nil {
		ac.ValidationError(c, errs)
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

func (ac *AuthController) Refresh(c *gin.Context) {
	var req validators.RefreshValidator

	if errs, err := validation.BindAndValidate(c, &req); err != nil {
		ac.ValidationError(c, errs)
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
	var req validators.LogoutValidator

	if errs, err := validation.BindAndValidate(c, &req); err != nil {
		ac.ValidationError(c, errs)
		return
	}

	if err := ac.service.Logout(req.Email); err != nil {
		ac.Error(c, "Logout failed", err.Error())
		return
	}

	ac.Success(c, "Logged out", nil)
}
