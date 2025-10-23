package routes

import (
	"goframe/internal/app/http/controllers/api"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {
	// Initialize controllers
	pingController := new(api.PingController)

	// Initialize services

 	//route group for API
	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/ping", pingController.Ping)
	}
}
