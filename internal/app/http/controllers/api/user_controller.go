package api

import (
	"goframe/internal/app/http/controllers"
	"goframe/internal/app/models"
	"goframe/internal/app/services"
	"strconv"

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
	users, err := uc.Service.GetAll()
	if err != nil {
		uc.Error(c, "Failed to fetch users", err.Error())
		return
	}
	uc.Success(c, "Users retrieved", users)
}

// GET /api/users/:id
func (uc *UserController) Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := uc.Service.GetByID(uint(id))
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
	if err := uc.Service.Create(&input); err != nil {
		uc.Error(c, "Failed to create user", err.Error())
		return
	}
	uc.Created(c, "User created", input)
}

// PUT /api/users/:id
func (uc *UserController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input map[string]interface{}
	if err := c.ShouldBindJSON(&input); err != nil {
		uc.ValidationError(c, err.Error())
		return
	}
	if err := uc.Service.Update(uint(id), input); err != nil {
		uc.Error(c, "Failed to update user", err.Error())
		return
	}
	uc.Success(c, "User updated", input)
}

// DELETE /api/users/:id
func (uc *UserController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := uc.Service.Delete(uint(id)); err != nil {
		uc.Error(c, "Failed to delete user", err.Error())
		return
	}
	uc.NoContent(c)
}
