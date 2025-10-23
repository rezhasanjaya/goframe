package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterWebRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to GoFrame 🚀")
	})
}
