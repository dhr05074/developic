package middleware

import (
	"code-connect/pkg/store"
	"github.com/labstack/echo/v4"
)

const authorizationHeaderKey = "Authorization"

func InjectUsernameToContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Request().Header.Get(authorizationHeaderKey)

		ctx := store.CtxWithUsername(c.Request().Context(), username)
		ctx = store.CtxWithIP(ctx, c.RealIP())

		c.SetRequest(c.Request().WithContext(ctx))

		return next(c)
	}
}
