package bootstrap

import (
	"goframe/internal/core/config"
	"log"

	"github.com/gin-gonic/gin"
)

type Application struct {
	Config *config.Config
	Router *gin.Engine
}

func NewApp() *Application {
	cfg := config.LoadConfig()

	InitDB(cfg)
	InitLogger()
	InitValidator()
	InitRedis(cfg)

	r := gin.Default()
	log.Println("✅ Application initialized")

	return &Application{
		Config: cfg,
		Router: r,
	}
}
