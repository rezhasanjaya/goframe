package api

import (
	"goframe/internal/app/http/controllers"
	"goframe/internal/app/models"
	"goframe/internal/app/services"
	"goframe/internal/app/validators"
	"goframe/internal/core/validation"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	controllers.BaseController
	Service *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		Service: services.NewUserService(),
	}
}
func (uc *UserController) Index(c *gin.Context) {
	users, err := uc.Service.Fetch()
	if err != nil {
		uc.Error(c, "Failed to fetch users", err.Error())
		return
	}
	uc.Success(c, "Users retrieved", users)
}

func (uc *UserController) Show(c *gin.Context) {
	uuid := c.Param("uuid")
	user, err := uc.Service.Get(uuid)
	if err != nil {
		uc.HandleServiceError(c, err)
		return
	}
	uc.Success(c, "User retrieved", user)
}

func (uc *UserController) Store(c *gin.Context) {
	var req validators.CreateValidator

	if errs, err := validation.BindAndValidate(c, &req); err != nil {
		uc.ValidationError(c, errs)
		return
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	resp, err := uc.Service.Create(user)
	if err != nil {
		uc.Error(c, "Failed to create user", err.Error())
		return
	}

	uc.Created(c, "User created", resp)
}

func (uc *UserController) Update(c *gin.Context) {
	uuid := c.Param("uuid")
	var req validators.UpdateValidator

	if errs, err := validation.BindAndValidate(c, &req); err != nil {
		uc.ValidationError(c, errs)
		return
	}

	user, err := uc.Service.Update(uuid, req)
	if err != nil {
		uc.HandleServiceError(c, err)
		return
	}

	uc.Success(c, "User updated", user)
}

func (uc *UserController) Delete(c *gin.Context) {
	uuid := c.Param("uuid")

	resp, err := uc.Service.Delete(uuid)
	if err != nil {
		uc.HandleServiceError(c, err)
		return
	}

	uc.Success(c, "User deleted", resp)
}
