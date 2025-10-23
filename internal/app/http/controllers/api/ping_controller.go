package api

import (
	"goframe/internal/app/http/controllers"

	"github.com/gin-gonic/gin"
)

type PingController struct {
	controllers.BaseController
}

func (p *PingController) Ping(c *gin.Context) {
	p.Success(c, "Pong", gin.H{
		"framework": "GoFrame",
		"version":   "1.0.0",
	})
}
