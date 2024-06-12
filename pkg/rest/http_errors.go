package rest

import (
	internalerrors "emailn/internal/errors"
	customerrors "emailn/pkg/api_utils/custom_errors"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

func handleErrorResponse(err error) (body any, status int) {
	if err == nil {
		return nil, http.StatusOK
	}
	wrapErrorMessage, errorMessage := getErrorMessage(err)
	logrus.WithFields(logrus.Fields{
		"error_type": wrapErrorMessage,
	}).Error(errorMessage)
	if errors.Is(err, customerrors.ErrValidation) {
		return map[string]string{"error": errorMessage}, http.StatusBadRequest
	}
	if errors.Is(err, customerrors.ErrNotFound) {
		return map[string]string{"error": errorMessage}, http.StatusNotFound
	}

	return map[string]string{"error": internalerrors.ErrInternalServer.Error()}, http.StatusInternalServerError
}

func getErrorMessage(err error) (wrapErrorMessage string, errorMessage string) {
	errWrap := errors.Unwrap(err)
	if errWrap != nil {
		wrapString := fmt.Sprintf("%s ", errWrap.Error())
		return errWrap.Error(), strings.Replace(err.Error(), wrapString, "", 1)
	}
	return "", err.Error()
}
