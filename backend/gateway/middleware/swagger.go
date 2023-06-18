package middleware

import (
	"code-connect/pkg/store"
	"context"
	"errors"
	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/getkin/kin-openapi/openapi3filter"
)

func Authenticate(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
	if input.SecuritySchemeName == "" {
		return nil
	}

	username := input.RequestValidationInput.Request.Header.Get("Authorization")
	if username == "" {
		return errors.New("authorization header is empty")
	}

	echoCtx := oapimiddleware.GetEchoContext(ctx)
	httpReq := echoCtx.Request()
	echoCtx.SetRequest(httpReq.WithContext(store.WithUsername(httpReq.Context(), username)))

	return nil
}
