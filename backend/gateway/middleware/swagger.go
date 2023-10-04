package middleware

import (
	"context"
	"errors"
	"github.com/getkin/kin-openapi/openapi3filter"
)

var errAuthorizationHeaderIsEmpty = errors.New("authorization header is empty")

func ValidateAuth(_ context.Context, input *openapi3filter.AuthenticationInput) error {
	if input.SecuritySchemeName == "" {
		return nil
	}

	username := input.RequestValidationInput.Request.Header.Get(authorizationHeaderKey)
	if username == "" {
		return errAuthorizationHeaderIsEmpty
	}

	return nil
}
