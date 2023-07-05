package problem

import (
	"ariga.io/entcache"
	"code-connect/ent"
	"code-connect/ent/problem"
	"code-connect/gateway"
	"code-connect/pkg/ai"
	"code-connect/pkg/log"
	"code-connect/pkg/store"
	"code-connect/pkg/str"
	"code-connect/schema/message"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alitto/pond"
	nanoid "github.com/matoous/go-nanoid/v2"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	generatingProblemPromptKey = "/prompts/problem/generating"
	languageTemplateKey        = "{LANGUAGE}"
	eloScoreTemplateKey        = "{ELO_SCORE}"
)

const (
	retrieveCodePrompt = "Give me the code"
)

const (
	poolMaxWorkers  = 10
	poolMaxCapacity = 1000
)

const (
	problemIDLength      = 8
	maxRequestCountPerIP = 5
)

var (
	ErrExceededMaxRequestCount = errors.New("exceeded max request count")
)

var pool = pond.New(poolMaxWorkers, poolMaxCapacity)

type Handler struct {
	paramClient       store.KeyValue
	redisClient       store.KeyValue
	entClient         *ent.Client
	logger            *zap.SugaredLogger
	reqeuestCh        chan message.ProblemMessage
	generateGPTClient ai.GPTClientGenerator
}

func NewHandler(
	paramClient store.KeyValue,
	redisClient store.KeyValue,
	entClient *ent.Client,
	gptClientGenerator ai.GPTClientGenerator,
	reqCh chan message.ProblemMessage,
) *Handler {
	l := log.NewZap().With("handler", "problem")

	return &Handler{
		paramClient:       paramClient,
		redisClient:       redisClient,
		entClient:         entClient,
		generateGPTClient: gptClientGenerator,
		logger:            l,
		reqeuestCh:        reqCh,
	}
}

type gptOutput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Code        string `json:"code"`
}

func (h *Handler) RequestProblem(ctx context.Context, request gateway.RequestProblemRequestObject) (gateway.RequestProblemResponseObject, error) {
	var internalServerErrorResp = gateway.RequestProblemdefaultJSONResponse{
		Body: gateway.Error{
			Code:    gateway.ServerError,
			Message: gateway.ServerErrorMessage,
		},
		StatusCode: http.StatusInternalServerError,
	}

	userIP, ok := store.IPFromContext(ctx)
	if !ok {
		h.logger.Errorw("IP 주소를 가져올 수 없습니다.")
		return internalServerErrorResp, nil
	}

	requestCount, err := h.requestCount(ctx, userIP)
	if err != nil {
		return internalServerErrorResp, nil
	}

	err = h.checkRequestCount(userIP, requestCount)
	if err != nil {
		return gateway.RequestProblem429JSONResponse{
			Code:    gateway.TooManyRequests,
			Message: gateway.TooManyRequestsMessage,
		}, nil
	}

	problemID := nanoid.Must(problemIDLength)

	// UUID만 담긴 문제 객체를 미리 생성한다.
	// 문제가 아예 없는 경우와, 문제 생성 요청이 들어간 상태를 구분하기 위해서다.
	if err := h.createProblemObject(ctx, problemID, request); err != nil {
		h.logger.Errorw("문제 객체 생성 실패", "error", err)
		return internalServerErrorResp, nil
	}

	err = h.increaseRequestCount(ctx, userIP, requestCount)
	if err != nil {
		h.logger.Errorw("요청 횟수 증가 실패", "error", err)
		return internalServerErrorResp, nil
	}

	// 문제를 백그라운드에서 생성한다.
	// 과도한 goroutine 생성 방지를 위해 Worker pool을 활용한다.
	pool.Submit(
		func() {
			var (
				output   gptOutput
				err      error
				emptyCtx = context.TODO()
			)

			output, err = h.requestProblem(emptyCtx, request)
			if err != nil {
				h.logger.Errorw("GPT를 통한 문제 출제 실패", "error", err)
				return
			}

			if err = h.saveProblem(emptyCtx, problemID, output); err != nil {
				h.logger.Errorw("문제 저장 실패", "error", err)
				return
			}

			h.reqeuestCh <- message.ProblemMessage{
				ID: problemID,
			}
		},
	)

	// 사용자에게는 문제 ID를 바로 반환하고, 백그라운드에서 문제를 생성한다.
	return gateway.RequestProblem202JSONResponse{
		ProblemId: problemID,
	}, nil
}

func (h *Handler) checkRequestCount(userIP string, requestCount int) error {
	if requestCount > maxRequestCountPerIP {
		h.logger.Errorw("IP 주소당 요청 횟수 초과", "ip", userIP)
		return ErrExceededMaxRequestCount
	}

	return nil
}

func (h *Handler) requestCount(ctx context.Context, userIP string) (count int, err error) {
	requestCountString, err := h.redisClient.Get(ctx, userIP)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(requestCountString)
}

func (h *Handler) increaseRequestCount(ctx context.Context, userIP string, requestCount int) error {
	return h.redisClient.Set(ctx, userIP, strconv.Itoa(requestCount+1))
}

func (h *Handler) createProblemObject(ctx context.Context, uuid string, request gateway.RequestProblemRequestObject) error {
	var difficulty *int
	if request.Body.EloScore != nil {
		difficultyValue := int(*request.Body.EloScore)
		difficulty = &difficultyValue
	}

	return h.entClient.Problem.Create().SetUUID(uuid).SetLanguage(request.Body.Language).SetNillableDifficulty(difficulty).Exec(ctx)
}

func (h *Handler) requestProblem(ctx context.Context, request gateway.RequestProblemRequestObject) (gptOutput, error) {
	now := time.Now()

	prompt, err := h.paramClient.Get(ctx, generatingProblemPromptKey)
	if err != nil {
		h.logger.Errorw("프롬포트 조회 실패", "key", generatingProblemPromptKey, "error", err)
		return gptOutput{}, err
	}

	gptClient, err := h.generateGPTClient()
	if err != nil {
		h.logger.Errorw("GPT 클라이언트 생성 실패", "error", err)
		return gptOutput{}, err
	}

	prompt = h.injectParameterToPrompt(prompt, request)
	gptClient.AddPrompt(prompt)

	result, err := gptClient.Complete(ctx)
	if err != nil {
		h.logger.Errorw("GPT의 프롬포트 처리 실패", "type", "gpt", "error", err)
		return gptOutput{}, err
	}

	var output gptOutput
	if err := json.Unmarshal([]byte(result), &output); err != nil {
		h.logger.Errorw("결과 JSON 언마샬 실패", "error", err)
		return gptOutput{}, err
	}

	output.Code, err = h.generateScratchCode(ctx, gptClient)
	if err != nil {
		h.logger.Errorw("스크래치 코드 생성 실패", "error", err)
		return gptOutput{}, err
	}

	h.logger.Infow("문제 생성 완료", "elapsed", time.Since(now).String())

	return output, nil
}

func (h *Handler) injectParameterToPrompt(prompt string, req gateway.RequestProblemRequestObject) string {
	prompt = strings.ReplaceAll(prompt, languageTemplateKey, string(req.Body.Language))
	prompt = strings.ReplaceAll(prompt, eloScoreTemplateKey, fmt.Sprintf("%d", req.Body.EloScore))

	return prompt
}

func (h *Handler) generateScratchCode(ctx context.Context, gptClient ai.GPTClient) (string, error) {
	gptClient.AddPrompt(retrieveCodePrompt)

	code, err := gptClient.Complete(ctx)
	if err != nil {
		h.logger.Errorw("GPT의 프롬포트 처리 실패", "type", "gpt", "error", err)
		return "", err
	}

	code = h.encodeCode(h.extractCode(code))

	return code, nil
}

func (h *Handler) extractCode(code string) string {
	return str.ExtractCodeBlocksFromMarkdown(code)[0]
}

func (h *Handler) encodeCode(code string) string {
	encoder := base64.StdEncoding
	return encoder.EncodeToString([]byte(code))
}

func (h *Handler) saveProblem(ctx context.Context, uuid string, output gptOutput) error {
	p, err := h.entClient.Problem.Query().Where(problem.UUID(uuid)).Only(ctx)
	if err != nil {
		h.logger.Errorw("failed to get problem", "category", "db", "error", err)
		return err
	}

	err = p.Update().SetTitle(output.Title).SetDescription(output.Description).SetCode(output.Code).Exec(ctx)
	if err != nil {
		h.logger.Errorw("failed to update problem", "category", "db", "error", err)
		return err
	}

	return nil
}

func (h *Handler) GetProblem(ctx context.Context, request gateway.GetProblemRequestObject) (gateway.GetProblemResponseObject, error) {
	const cacheTTL = 1 * time.Second

	queriedProblem, err := h.entClient.Problem.Query().Where(problem.UUID(request.Id)).Only(
		entcache.WithTTL(
			ctx, cacheTTL,
		),
	)
	if err != nil {
		if ent.IsNotFound(err) {
			return gateway.GetProblem404JSONResponse{
				Code:    gateway.ProblemNotFound,
				Message: gateway.ProblemNotFoundMessage,
			}, nil
		}

		h.logger.Errorw("failed to get problem", "category", "db", "error", err)
		return gateway.GetProblemdefaultJSONResponse{
			Body: gateway.Error{
				Code:    gateway.ServerError,
				Message: gateway.ServerErrorMessage,
			},
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	if h.isProblemNotCompleted(queriedProblem) {
		return gateway.GetProblem409JSONResponse{
			Code:    gateway.ProblemNotReady,
			Message: gateway.ProblemNotReadyMessage,
		}, nil
	}

	return gateway.GetProblem200JSONResponse{
		N200GetProblemJSONResponse: gateway.N200GetProblemJSONResponse{
			Code:        queriedProblem.Code,
			Id:          queriedProblem.UUID,
			Title:       queriedProblem.Title,
			Description: queriedProblem.Description,
		},
	}, nil
}

func (h *Handler) isProblemNotCompleted(p *ent.Problem) bool {
	return p.Code == "" || p.Title == ""
}
