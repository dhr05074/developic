package main

import (
	"ariga.io/entcache"
	"code-connect/ent"
	"code-connect/gateway"
	"code-connect/gateway/handler"
	customMiddleware "code-connect/gateway/middleware"
	"code-connect/pkg/ai"
	"code-connect/pkg/aws"
	"code-connect/pkg/db"
	"code-connect/pkg/log"
	"code-connect/pkg/store"
	"code-connect/problem"
	"code-connect/record"
	"code-connect/schema/message"
	"code-connect/worker/score"
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/getkin/kin-openapi/openapi3filter"
	_ "github.com/go-sql-driver/mysql"
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
	app.Use(customMiddleware.AttachUsername)

	ctx := context.Background()
	kvStore := mustInitKVStore(ctx)
	gptClient := mustInitGPTClient()
	entClient := mustInitEntClient(ctx)

	reqCh := make(chan message.ProblemMessage)
	subCh := make(chan message.ProblemMessage)

	w := score.NewScoreWorker(score.NewScoreWorkerParams{
		ParamClient: kvStore,
		EntClient:   entClient,
		GptClient:   gptClient,
		ProblemCh:   reqCh,
		SubmitCh:    subCh,
	})

	go func() {
		if err := w.Run(ctx); err != nil {
			l.Fatalw("failed to run score worker", "err", err)
		}
	}()

	problemHandler := problem.NewHandler(kvStore, gptClient, entClient, reqCh)
	recordHandler := record.NewHandler(record.NewHandlerParams{
		ParamClient: kvStore,
		EntClient:   entClient,
		SubmitCh:    subCh,
	})

	strictHandler := handler.NewStrictHandler(problemHandler, recordHandler)
	serverInterface := gateway.NewStrictHandler(strictHandler, []gateway.StrictMiddlewareFunc{})
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

	swagger.Servers = nil

	return oapimiddleware.OapiRequestValidatorWithOptions(swagger, &oapimiddleware.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: customMiddleware.ValidateAuth,
		},
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
	if !ok || apiToken == "" {
		l.Fatal("CHATGPT_API_KEY is not set")
	}
	cli := ai.NewOpenAI(openai.NewClient(apiToken))
	return cli
}

func mustInitEntClient(ctx context.Context) *ent.Client {
	drv := db.NewCachedEntDriver()
	client := ent.NewClient(ent.Driver(drv))
	if err := client.Schema.Create(entcache.Skip(ctx)); err != nil {
		l.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
