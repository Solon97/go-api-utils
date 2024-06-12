package customerrors

import (
	"errors"
	"fmt"
)

var (
	ErrRepository = errors.New("[repository error]")
	ErrValidation = errors.New("[validation error]")
	ErrNotFound   = errors.New("[not found error]")
)

func NewRepositoryError(err error) error {
	return newCustomError(err, ErrRepository)
}

func NewValidationError(err error) error {
	return newCustomError(err, ErrValidation)
}

func NewNotFoundError(err error) error {
	return newCustomError(err, ErrNotFound)
}

func newCustomError(err error, wrapError error) error {
	return fmt.Errorf("%w %v", wrapError, err)
}
