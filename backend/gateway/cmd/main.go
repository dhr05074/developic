package main

import (
	"ariga.io/entcache"
	"code-connect/gateway"
	"code-connect/gateway/handler"
	"code-connect/pkg/ai"
	"code-connect/pkg/db"
	"code-connect/pkg/log"
	"code-connect/problem"
	"code-connect/problem/ent"
	"context"
	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sashabaranov/go-openai"
	"os"
)

var l = log.NewZapSugaredLogger().With("service", "gateway")

func main() {
	apiKey := os.Getenv("CHATGPT_API_KEY")
	cli := openai.NewClient(apiKey)
	openAIClient := ai.NewOpenAI(cli)

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

	problemHandler := problem.NewHandler(openAIClient, entClient)
	strictHandler := handler.NewStrictHandler(openAIClient, problemHandler)
	srvHandler := gateway.NewStrictHandler(strictHandler, []gateway.StrictMiddlewareFunc{})

	app := echo.New()
	defer func() {
		if err := app.Shutdown(context.Background()); err != nil {
			l.Fatalw("failed to close echo app", "err", err)
		}
	}()
	app.HideBanner = true

	gateway.RegisterHandlers(app, srvHandler)

	swagger, err := gateway.GetSwagger()
	if err != nil {
		panic(err)
	}

	swagger.Servers = nil

	app.Use(middleware.CORS())
	app.Use(middleware.Logger())
	app.Use(openAPIRequestValidator(swagger))

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
