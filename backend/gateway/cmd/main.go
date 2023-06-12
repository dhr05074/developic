package main

import (
	"code-connect/pkg/log"
	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var l = log.NewZap().With("service", "gateway")

func main() {
	app := echo.New()

	app.Use(middleware.CORS())
	app.Use(middleware.Logger())

	l.Infow("starting server", "port", 3000)
	if err := app.Start(":3000"); err != nil {
		panic(err)
	}
}

func openAPIRequestValidator(swagger *openapi3.T) echo.MiddlewareFunc {
	return oapimiddleware.OapiRequestValidatorWithOptions(swagger, &oapimiddleware.Options{
		Options: openapi3filter.Options{},
	})
}
