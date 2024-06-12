package rest

import (
	customerrors "emailn/pkg/api_utils/custom_errors"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleErrorResponse(t *testing.T) {
	assert := assert.New(t)

	t.Run("Given validation error When calling handleErrorResponse Then return 400", func(t *testing.T) {
		err := customerrors.NewValidationError(errors.New("domain validation error"))
		_, status := handleErrorResponse(err)
		assert.Equal(status, http.StatusBadRequest)
	})

	t.Run("Given not found error When calling handleErrorResponse Then return 404", func(t *testing.T) {
		err := customerrors.NewNotFoundError(errors.New("not found error"))
		_, status := handleErrorResponse(err)
		assert.Equal(status, http.StatusNotFound)
	})

	t.Run("Given generic error When calling handleErrorResponse Then return 500", func(t *testing.T) {
		err := errors.New("generic error")
		_, status := handleErrorResponse(err)
		assert.Equal(status, http.StatusInternalServerError)
	})

	t.Run("Given nil error When calling handleErrorResponse Then return 200", func(t *testing.T) {
		_, status := handleErrorResponse(nil)
		assert.Equal(status, http.StatusOK)
	})
}
