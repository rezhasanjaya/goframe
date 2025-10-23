package routes

import (
	"goframe/internal/app/http/controllers/api"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {
	// Initialize controllers
	// pingController := new(api.PingController)
	userController := api.NewUserController()

	// Initialize services

 	//route group for API
	apiGroup := r.Group("/api")
	{
		// apiGroup.GET("/ping", pingController.Ping)
		userGroup := apiGroup.Group("/users")
		{
			userGroup.GET("/", userController.Index)       // GET /api/users
			userGroup.GET("/:id", userController.Show)     // GET /api/users/:id
			userGroup.POST("/", userController.Store)      // POST /api/users
			userGroup.PUT("/:id", userController.Update)   // PUT /api/users/:id
			userGroup.DELETE("/:id", userController.Delete) // DELETE /api/users/:id
		}
	}
}
