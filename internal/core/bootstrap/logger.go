package bootstrap

import (
	"log"

	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger() {
	var err error
	Logger, err = zap.NewProduction()
	if err != nil {
		log.Fatalf("❌ Failed to init zap logger: %v", err)
	}
	log.Println("✅ Logger initialized")
}
