package validation

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error, req interface{}) map[string]string {
	errors := map[string]string{}

	if errs, ok := err.(validator.ValidationErrors); ok {

		t := reflect.TypeOf(req)
		if t.Kind() == reflect.Pointer {
			t = t.Elem()
		}

		for _, e := range errs {
			field, _ := t.FieldByName(e.Field())
			jsonKey := field.Tag.Get("json")
			if jsonKey == "" {
				jsonKey = e.Field()
			}

			errors[jsonKey] = msgForTag(e.Tag())
		}

		return errors
	}

	errors["error"] = err.Error()
	return errors
}

func msgForTag(tag string) string {
	switch tag {
	case "required":
		return "field is required"
	case "email":
		return "must be a valid email"
	case "min":
		return "value is too short"
	case "max":
		return "value is too long"
	}
	return "invalid value"
}
