package problem

import (
	"code-connect/ent"
	"code-connect/ent/problem"
	"code-connect/gateway"
	"code-connect/pkg/ai"
	"code-connect/pkg/log"
	"code-connect/pkg/store"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	nanoid "github.com/matoous/go-nanoid/v2"
	"go.uber.org/zap"
	"strings"
)

type Handler struct {
	paramClient store.KV
	gptClient   ai.GPTClient
	entClient   *ent.Client
	log         *zap.SugaredLogger
}

func NewHandler(paramClient store.KV, gptClient ai.GPTClient, entClient *ent.Client) *Handler {
	l := log.NewZap().With("handler", "problem")
	return &Handler{paramClient: paramClient, gptClient: gptClient, entClient: entClient, log: l}
}

type Output struct {
	Title string `json:"title"`
	Code  string `json:"code"`
}

const (
	generatingProblemPromptKey = "/prompts/problem/generating"
	languageTemplateKey        = "{LANGUAGE}"
	eloScoreTemplateKey        = "{ELO_SCORE}"
	defaultELOScore            = 1500
)

const (
	ServerErrorMessage = "일시적인 서버 오류입니다. 잠시 후에 다시 시도해주세요."
)

func (h *Handler) RequestProblem(ctx context.Context, request gateway.RequestProblemRequestObject) (gateway.RequestProblemResponseObject, error) {
	problemID := nanoid.Must(8)

	// UUID만 담긴 문제 객체를 미리 생성한다.
	// 문제가 아예 없는 경우와, 문제 생성 요청이 들어간 상태를 구분하기 위해서다.
	if err := h.createProblem(ctx, problemID, request.Body.Language); err != nil {
		h.log.Errorw("failed to create problem", "category", "db", "error", err)
		return gateway.RequestProblemdefaultJSONResponse{
			Body: gateway.Error{
				Message: ServerErrorMessage,
			},
			StatusCode: 500,
		}, nil
	}

	// 문제를 백그라운드에서 생성한다.
	go func() {
		anotherCtx := context.TODO()

		output, err := h.requestProblem(anotherCtx, request)
		if err != nil {
			return
		}

		if err := h.saveProblem(anotherCtx, problemID, output); err != nil {
			h.log.Errorw("failed to save problem", "category", "db", "error", err)
			return
		}
	}()

	// 사용자에게는 문제 ID를 바로 반환하고, 백그라운드에서 문제를 생성한다.
	return gateway.RequestProblem202JSONResponse{
		ProblemId: problemID,
	}, nil
}

func (h *Handler) requestProblem(ctx context.Context, request gateway.RequestProblemRequestObject) (Output, error) {
	prompt, err := h.paramClient.Get(ctx, generatingProblemPromptKey)
	if err != nil {
		h.log.Errorw("failed to get prompt", "category", "kv", "error", err)
		return Output{}, err
	}

	prompt = h.injectData(prompt, request)
	h.gptClient.AddPrompt(prompt)

	result, err := h.gptClient.Complete(ctx)
	if err != nil {
		h.log.Errorw("failed to complete prompt", "category", "external_api", "api_category", "gpt", "error", err)
		return Output{}, err
	}

	var output Output
	if err := json.Unmarshal([]byte(result), &output); err != nil {
		h.log.Errorw("failed to unmarshal result", "category", "json", "error", err)
		return Output{}, err
	}

	encoder := base64.StdEncoding
	output.Code = encoder.EncodeToString([]byte(output.Code))

	return output, nil
}

func (h *Handler) injectData(prompt string, req gateway.RequestProblemRequestObject) string {
	prompt = strings.ReplaceAll(prompt, languageTemplateKey, string(req.Body.Language))
	prompt = strings.ReplaceAll(prompt, eloScoreTemplateKey, fmt.Sprintf("%d", defaultELOScore))

	return prompt
}

func (h *Handler) createProblem(ctx context.Context, uuid string, language gateway.ProgrammingLanguage) error {
	return h.entClient.Problem.Create().SetUUID(uuid).SetLanguage(language).Exec(ctx)
}

func (h *Handler) saveProblem(ctx context.Context, uuid string, output Output) error {
	return h.entClient.Problem.Update().Where(problem.UUID(uuid)).SetTitle(output.Title).SetCode(output.Code).Exec(ctx)
}

func (h *Handler) GetProblem(ctx context.Context, request gateway.GetProblemRequestObject) (gateway.GetProblemResponseObject, error) {
	p, err := h.entClient.Problem.Query().Where(problem.UUID(request.Id)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return gateway.GetProblem404Response{}, nil
		}

		h.log.Errorw("failed to get problem", "category", "db", "error", err)
		return gateway.GetProblemdefaultJSONResponse{
			Body: gateway.Error{
				Message: ServerErrorMessage,
			},
			StatusCode: 500,
		}, nil
	}

	// 문제 생성 요청은 들어갔지만 아직 문제가 생성되지 않은 경우
	if p.Code == "" || p.Title == "" {
		return gateway.GetProblem409Response{}, nil
	}

	return gateway.GetProblem200JSONResponse{
		N200GetProblemJSONResponse: gateway.N200GetProblemJSONResponse{
			Code: p.Code,
			Id:   p.UUID,
			Name: p.Title,
		},
	}, nil
}
