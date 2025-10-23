package api

import (
	"goframe/internal/app/http/controllers"
	"goframe/internal/app/models"
	"goframe/internal/app/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	controllers.BaseController // embed langsung
	Service *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		Service: services.NewUserService(),
	}
}

// GET /api/users
func (uc *UserController) Index(c *gin.Context) {
	users, err := uc.Service.Fetch()
	if err != nil {
		uc.Error(c, "Failed to fetch users", err.Error())
		return
	}
	uc.Success(c, "Users retrieved", users)
}

// GET /api/users/:uuid
func (uc *UserController) Show(c *gin.Context) {
	uuid := c.Param("uuid")
	user, err := uc.Service.Get(uuid)
	if err != nil {
		uc.Error(c, "User not found", err.Error())
		return
	}
	uc.Success(c, "User retrieved", user)
}

// POST /api/users
func (uc *UserController) Store(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		uc.ValidationError(c, err.Error())
		return
	}

	resp, err := uc.Service.Create(&input)
	if err != nil {
		uc.Error(c, "Failed to create user", err.Error())
		return
	}
	uc.Created(c, "User created", resp)
}

// PUT /api/users/:uuid
func (uc *UserController) Update(c *gin.Context) {
	uuid := c.Param("uuid")
	var input map[string]interface{}
	if err := c.ShouldBindJSON(&input); err != nil {
		uc.ValidationError(c, err.Error())
		return
	}

	user, err := uc.Service.Update(uuid, input)
	if err != nil {
		uc.Error(c, "Failed to update user", err.Error())
		return
	}
	uc.Success(c, "User updated", user)
}

// DELETE /api/users/:uuid
func (uc *UserController) Delete(c *gin.Context) {
	uuid := c.Param("uuid")

	resp, err := uc.Service.Delete(uuid)
	if err != nil {
		uc.Error(c, "Failed to delete user", err.Error())
		return
	}

	uc.Success(c, "User deleted", resp)
}

