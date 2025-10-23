package plugins

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func PrintRegisteredRoutes(r *gin.Engine) {
	routes := r.Routes()
	fmt.Println("\n📋 Registered Routes:")
	fmt.Println("----------------------------")
	for _, route := range routes {
		fmt.Printf("%-6s %s\n", route.Method, route.Path)
	}
	fmt.Println("----------------------------\n")
}
