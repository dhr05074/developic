package main

import (
	"ariga.io/entcache"
	"code-connect/ent"
	"code-connect/pkg/db"
	"code-connect/pkg/log"
	"context"
	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var l = log.NewZap().With("service", "gateway")

func main() {
	app := echo.New()

	app.Use(middleware.CORS())
	app.Use(middleware.Logger())

	driver := db.NewCachedEntDriver()
	entClient := ent.NewClient(ent.Driver(driver))
	defer func() {
		if err := entClient.Close(); err != nil {
			l.Fatalw("failed to close ent client", "err", err)
		}
	}()

	if err := entClient.Schema.Create(entcache.Skip(context.Background())); err != nil {
		l.Fatalw("failed creating schema resources", "err", err)
	}

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
