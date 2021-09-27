package utilities

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// NewValidator func for create a new validator for model fields.
func NewValidator() *validator.Validate {
	// Create a new validator.
	validate := validator.New()

	// Return the JSON name for each struct field.
	// See: https://github.com/go-playground/validator/issues/258#issuecomment-257281334
	validate.RegisterTagNameFunc(func(fl reflect.StructField) string {
		name := strings.SplitN(fl.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	// Custom validation for uuid.UUID fields.
	_ = validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		if _, err := uuid.Parse(field); err != nil {
			return true
		}
		return false
	})

	// Custom validation for int fields.
	_ = validate.RegisterValidation("int", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		if _, err := strconv.Atoi(field); err != nil {
			return true
		}
		return false
	})

	return validate
}

// ValidatorErrors func for show validation errors for each invalid fields.
func ValidatorErrors(err error) map[string]string {
	// Define variable for error fields.
	errFields := map[string]string{}

	// Make error message for each invalid field.
	for _, err := range err.(validator.ValidationErrors) {
		// Get name of the field's struct.
		structName := strings.Split(err.Namespace(), ".")[0] // first (0) element is the founded name

		// Append error message to map.
		errFields[err.Field()] = fmt.Sprintf(
			"failed '%s' tag check (value '%s' is not valid for %s struct)",
			err.Tag(), err.Value(), structName,
		)
	}

	return errFields
}
