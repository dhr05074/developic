package middleware

import (
	"context"
	"errors"
	"github.com/getkin/kin-openapi/openapi3filter"
)

func ValidateAuth(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
	if input.SecuritySchemeName == "" {
		return nil
	}

	username := input.RequestValidationInput.Request.Header.Get("Authorization")
	if username == "" {
		return errors.New("authorization header is empty")
	}

	return nil
}
