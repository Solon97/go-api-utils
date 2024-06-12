package validator

import (
	customerrors "emailn/pkg/api_utils/custom_errors"
	"encoding/json"
	"errors"
	"io"

	"github.com/xeipuuv/gojsonschema"
)

var (
	ErrEmptyBody = errors.New("empty request body")
)

func ValidateJSON[T any](requestBody io.ReadCloser, jsonSchema string) (structBody *T, err error) {
	if requestBody == nil {
		return nil, customerrors.NewValidationError(ErrEmptyBody)
	}
	if jsonSchema == "" {
		return nil, errors.New("json schema is empty")
	}

	content, err := io.ReadAll(requestBody)
	if err != nil {
		return
	}

	loader := gojsonschema.NewStringLoader(string(content))
	schemaLoader := gojsonschema.NewStringLoader(jsonSchema)

	result, err := gojsonschema.Validate(schemaLoader, loader)
	if err != nil {
		return nil, err
	}
	if !result.Valid() {
		for _, desc := range result.Errors() {
			validationError := errors.New(desc.String())
			return nil, customerrors.NewValidationError(validationError)
		}
	}

	structBody = new(T)
	err = json.Unmarshal([]byte(content), structBody)
	if err != nil {
		return nil, err
	}
	return structBody, nil
}
