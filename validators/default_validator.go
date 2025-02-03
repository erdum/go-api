package validators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field string
	Tag string
}

func (ve *ValidationError) Error() string {
	return fmt.Sprintf("The field %s %s", ve.Field, ve.Tag)
}

type DefaultValidator struct {
	validator *validator.Validate
}

func NewDefaultValidator() *DefaultValidator {
	return &DefaultValidator{validator: validator.New()}
}

func (cv *DefaultValidator) Validate(i interface{}) error {
	err := cv.validator.Struct(i)

	if err == nil {
		return nil
	}
	last_err := err.(validator.ValidationErrors)[0]

	return &ValidationError{
		Field: last_err.Field(),
		Tag: last_err.Tag(),
	}
}
