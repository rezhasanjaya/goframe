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

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Recovery())

	if err := r.SetTrustedProxies(nil); err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	log.Println("✅ Application initialized")

	return &Application{
		Config: cfg,
		Router: r,
	}
}
