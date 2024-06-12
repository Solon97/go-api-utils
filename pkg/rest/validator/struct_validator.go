package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator"
)

func ValidateStruct(s interface{}) error {
	validate := validator.New()
	err := validate.Struct(s)
	if err == nil {
		return nil
	}
	validationError := err.(validator.ValidationErrors)[0]
	field := strings.ToLower(validationError.Field())
	switch validationError.Tag() {
	case "required":
		return fmt.Errorf(ErrRequiredFieldPattern, field)
	case "min":
		return fmt.Errorf(ErrMinFieldPattern, field, validationError.Param())
	case "max":
		return fmt.Errorf(ErrMaxFieldPattern, field, validationError.Param())
	case "email":
		return fmt.Errorf(ErrEmailFieldPattern, field)
	}

	return err
}
