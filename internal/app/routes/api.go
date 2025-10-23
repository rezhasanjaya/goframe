package routes

import (
	"goframe/internal/app/http/controllers/api"
	authctl "goframe/internal/app/http/controllers/api"
	"goframe/internal/app/http/middleware"
	"goframe/internal/core/config"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine,  cfg *config.Config) {
	// Initialize controllers
	// pingController := new(api.PingController)
	userController := api.NewUserController()
	authController := authctl.NewAuthController(cfg)

	// Initialize services

 	//route group for API
	apiGroup := r.Group("/api")
	{
		
		apiGroup.POST("/auth/register", authController.Register)
		apiGroup.POST("/auth/login", authController.Login)
		apiGroup.POST("/auth/refresh", authController.Refresh)
		apiGroup.POST("/auth/logout", authController.Logout)

		// apiGroup.GET("/ping", pingController.Ping)
		userGroup := apiGroup.Group("/users")
		userGroup.Use(middleware.JWTAuth(cfg))
		{
			userGroup.GET("/", userController.Index)          
			userGroup.POST("/", userController.Store)     
			userGroup.GET("/:uuid", userController.Show)      
			userGroup.PUT("/:uuid", userController.Update)   
			userGroup.DELETE("/:uuid", userController.Delete) 
		}
	}
}
