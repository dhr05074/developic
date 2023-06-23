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
	"code-connect/schema"
	"code-connect/schema/message"
	"code-connect/user"
	"code-connect/worker/score"
	"context"
	"fmt"
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

const (
	certFilePath = "/cert/fullchain.pem"
	keyFilePath  = "/cert/privkey.pem"
)

const (
	defaultServerPort = 3000
)

func main() {
	app := echo.New()

	app.Use(middleware.CORS())
	app.Use(middleware.Logger())
	app.Use(mustGetSwaggerValidator())
	app.Use(customMiddleware.InjectUsernameToContext)

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

	problemHandler := problem.NewHandler(kvStore, entClient, reqCh)
	recordHandler := record.NewHandler(record.NewHandlerParams{
		ParamClient: kvStore,
		EntClient:   entClient,
		SubmitCh:    subCh,
	})
	userHandler := user.NewHandler(entClient)

	strictHandler := handler.NewStrictHandler(problemHandler, recordHandler, userHandler)
	serverInterface := gateway.NewStrictHandler(strictHandler, []gateway.StrictMiddlewareFunc{})
	gateway.RegisterHandlers(app, serverInterface)

	l.Infow("서버를 시작합니다.", "포트", defaultServerPort)
	if err := app.StartTLS(fmt.Sprintf(":%d", defaultServerPort), certFilePath, keyFilePath); err != nil {
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
	apiToken, ok := os.LookupEnv(schema.ChatGPTAPIKeyEnvKey)
	if !ok || apiToken == "" {
		l.Fatalf("%s 환경 변수가 설정되지 않았습니다.", schema.ChatGPTAPIKeyEnvKey)
	}
	cli := ai.NewOpenAI(openai.NewClient(apiToken))
	return cli
}

func mustInitEntClient(ctx context.Context) *ent.Client {
	drv := db.NewCachedEntDriver()
	client := ent.NewClient(ent.Driver(drv))
	if err := client.Schema.Create(entcache.Skip(ctx)); err != nil {
		l.Fatalf("DB 스키마에 맞는 리소스 생성에 실패했습니다: %v", err)
	}
	return client
}
