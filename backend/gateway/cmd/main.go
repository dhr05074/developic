package main

import (
	"code-connect/ent"
	"code-connect/gateway"
	"code-connect/gateway/handler"
	"code-connect/pkg/ai"
	"code-connect/pkg/aws"
	"code-connect/pkg/db"
	"code-connect/pkg/log"
	"code-connect/pkg/store"
	"code-connect/problem"
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sashabaranov/go-openai"
	"os"
)

var l = log.NewZap().With("service", "gateway")

func main() {
	app := echo.New()

	app.Use(middleware.CORS())
	app.Use(middleware.Logger())
	app.Use(mustGetSwaggerValidator())

	ctx := context.Background()
	kvStore := mustInitKVStore(ctx)
	gptClient := mustInitGPTClient()
	entClient := mustInitEntClient(ctx)

	problemHandler := problem.NewHandler(kvStore, gptClient, entClient)

	strictHandler := handler.NewStrictHandler(problemHandler)
	serverInterface := gateway.NewStrictHandler(strictHandler, []gateway.StrictHandlerFunc{})
	gateway.RegisterHandlers(app, serverInterface)

	l.Infow("starting server", "port", 3000)
	if err := app.Start(":3000"); err != nil {
		panic(err)
	}
}

func mustGetSwaggerValidator() echo.MiddlewareFunc {
	swagger, err := gateway.GetSwagger()
	if err != nil {
		l.Fatalf("failed to get swagger: %v", err)
	}

	return oapimiddleware.OapiRequestValidatorWithOptions(swagger, &oapimiddleware.Options{
		Options: openapi3filter.Options{},
	})
}

func mustInitKVStore(ctx context.Context) store.KV {
	cfg, err := aws.Config(ctx)
	if err != nil {
		l.Fatalf("failed to get aws config: %v", err)
	}
	ssmClient := ssm.NewFromConfig(cfg)

	return aws.NewSSMClient(ssmClient)
}

func mustInitGPTClient() ai.GPTClient {
	apiToken, ok := os.LookupEnv("CHATGPT_API_KEY")
	if !ok {
		l.Fatal("CHATGPT_API_KEY is not set")
	}
	cli := ai.NewOpenAI(openai.NewClient(apiToken))
	return cli
}

func mustInitEntClient(ctx context.Context) *ent.Client {
	drv := db.NewCachedEntDriver()
	client := ent.NewClient(ent.Driver(drv))
	if err := client.Schema.Create(ctx); err != nil {
		l.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
