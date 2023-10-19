package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type ValidationError struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

type XValidator struct {
	validator *validator.Validate
}

func NewValidator() *XValidator {
	return &XValidator{
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (v *XValidator) Validate(data interface{}) []ValidationError {
	var validationErrors []ValidationError

	errs := v.validator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ValidationError

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func FormatValidationErrors(errors []ValidationError) *fiber.Error {
	errMsgs := make([]string, 0)

	for _, err := range errors {
		errMsgs = append(errMsgs, fmt.Sprintf(
			"[%s]: '%v' | Needs to implement '%s'",
			err.FailedField,
			err.Value,
			err.Tag,
		))
	}

	return &fiber.Error{
		Code:    fiber.ErrBadRequest.Code,
		Message: strings.Join(errMsgs, " and "),
	}
}
