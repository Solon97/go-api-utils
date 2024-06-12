package rest

import (
	internalerrors "emailn/internal/errors"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HandleResponse(t *testing.T) {
	assert := assert.New(t)

	t.Run("Given status and empty error and body When calling HandleResponse Then return the status", func(t *testing.T) {
		endpointFunc := func(w http.ResponseWriter, r *http.Request) (Response, error) {
			return Response{
				Body:       nil,
				StatusCode: http.StatusOK,
			}, nil
		}
		handler := HandleResponse(endpointFunc)
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)
		assert.Equal(http.StatusOK, w.Code)
		assert.Empty(w.Body.String())
	})

	t.Run("Given body and status and no error When calling HandleResponse Then return the body and status", func(t *testing.T) {
		endpointFunc := func(w http.ResponseWriter, r *http.Request) (Response, error) {
			return Response{
				Body:       map[string]string{"key": "value"},
				StatusCode: http.StatusOK,
			}, nil
		}
		handler := HandleResponse(endpointFunc)
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)
		assert.Equal(http.StatusOK, w.Code)
		assert.Contains(w.Body.String(), "key")
	})

	t.Run("Given status and generic error When calling HandleResponse Then return the status and the error message", func(t *testing.T) {
		endpointFunc := func(w http.ResponseWriter, r *http.Request) (Response, error) {
			return Response{}, errors.New("some error")
		}
		handler := HandleResponse(endpointFunc)
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)
		assert.Equal(http.StatusInternalServerError, w.Code)
		assert.Contains(w.Body.String(), internalerrors.ErrInternalServer.Error())
	})

	t.Run("Given body and status and error When calling HandleResponse Then return status and the error message in body", func(t *testing.T) {
		endpointFunc := func(w http.ResponseWriter, r *http.Request) (Response, error) {
			return Response{
				Body:       map[string]string{"key": "value"},
				StatusCode: http.StatusBadRequest,
			}, errors.New("some error")
		}
		handler := HandleResponse(endpointFunc)
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)
		assert.Equal(http.StatusInternalServerError, w.Code)
		assert.Contains(w.Body.String(), internalerrors.ErrInternalServer.Error())
		assert.NotContains(w.Body.String(), "key")
	})

	t.Run("Given body and status and wrapped internal server error When calling HandleResponse Then return status and the internal server error in body", func(t *testing.T) {
		originalErrorMessage := "some error"
		endpointFunc := func(w http.ResponseWriter, r *http.Request) (Response, error) {
			err := errors.New(originalErrorMessage)
			return Response{
				Body:       nil,
				StatusCode: http.StatusBadRequest,
			}, fmt.Errorf("%w: %s", internalerrors.ErrInternalServer, err.Error())
		}
		handler := HandleResponse(endpointFunc)
		r := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)
		assert.Equal(http.StatusInternalServerError, w.Code)
		assert.Contains(w.Body.String(), internalerrors.ErrInternalServer.Error())
		assert.NotContains(w.Body.String(), originalErrorMessage)
	})
}
