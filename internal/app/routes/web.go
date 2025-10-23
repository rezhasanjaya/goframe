package routes

import (
	"path/filepath"
	"runtime"

	"goframe/internal/app/http/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterWebRoutes(r *gin.Engine) {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Join(filepath.Dir(b), "../views")

	r.LoadHTMLGlob(filepath.Join(basePath, "*.html"))
	r.Static("/images", "./public/images")
	r.GET("/", controllers.WelcomePage)
}
