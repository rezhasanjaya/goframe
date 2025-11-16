package validation

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func BindAndValidate(c *gin.Context, req interface{}) (map[string]string, error) {

	if err := c.ShouldBindJSON(req); err != nil {
		return FormatValidationError(err, req), err
	}

	if err := Validate.Struct(req); err != nil {
		return FormatValidationError(err, req), err
	}

	return nil, nil
}
