package store

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http/httptest"
	"testing"
)

type TestHandler struct{}

func (h *TestHandler) Test(ctx context.Context) (string, error) {
	username, ok := UsernameFromContext(ctx)
	if !ok {
		return "", echo.ErrUnauthorized
	}

	return username, nil
}

func TestUsernameFromContext_WhenExtractedAtEchoRequestContext(t *testing.T) {
	app := echo.New()
	app.Use(
		func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				c.SetRequest(c.Request().WithContext(CtxWithUsername(c.Request().Context(), "test")))
				return next(c)
			}
		},
	)
	app.GET(
		"/", func(c echo.Context) error {
			username, _ := new(TestHandler).Test(c.Request().Context())

			return c.String(200, username)
		},
	)

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)
	if rec.Body.String() != "test" {
		t.Errorf("expected %s, got %s", "test", rec.Body.String())
	}
}
