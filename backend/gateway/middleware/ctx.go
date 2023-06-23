package middleware

import (
	"code-connect/pkg/store"
	"github.com/labstack/echo/v4"
)

func InjectUsernameToContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Request().Header.Get("Authorization")
		ctx := store.WithUsername(c.Request().Context(), username)
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
