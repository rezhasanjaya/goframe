package bootstrap

import (
	"log"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func InitValidator() {
	Validate = validator.New()
	log.Println("✅ Validator initialized")
}
