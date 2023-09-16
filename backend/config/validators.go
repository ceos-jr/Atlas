package config

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type RequestBodyValidator struct {
	Validator *validator.Validate
}

type ValidationError struct {
	Message string
}

func NewValidator() *RequestBodyValidator {
	return &RequestBodyValidator{
		Validator: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (v *RequestBodyValidator) Validate(body interface{}) []ValidationError {
	var validationErrors []ValidationError

	errors := v.Validator.Struct(body)

	if errors != nil {
		for _, valError := range errors.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, ValidationError{
				Message: fmt.Sprintf(
					"Field %s has an invalid value: %s",
					valError.Field(), valError.Value(),
				),
			})
		}
	}

	return validationErrors
}
