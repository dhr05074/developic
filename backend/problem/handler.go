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
	"fmt"
	nanoid "github.com/matoous/go-nanoid/v2"
	"go.uber.org/zap"
	"strings"
	"time"
)

const (
	problemNotReadyCode = "ProblemNotReady"
	problemNotFoundCode = "ProblemNotFound"
)

const (
	problemNotReadyMessage = "문제가 아직 준비되지 않았습니다. 잠시 후 다시 시도해주세요."
	problemNotFoundMessage = "문제가 존재하지 않습니다. 문제 ID를 확인해주세요."
	serverErrorMessage     = "일시적인 서버 오류입니다. 잠시 후에 다시 시도해주세요."
)

const (
	generatingProblemPromptKey = "/prompts/problem/generating"
	languageTemplateKey        = "{LANGUAGE}"
	eloScoreTemplateKey        = "{ELO_SCORE}"
)

const (
	retrieveCodePrompt = "Give me the code"
)

var serverErrResp = gateway.GetProblemdefaultJSONResponse{
	Body: gateway.Error{
		Message: serverErrorMessage,
	},
	StatusCode: 500,
}

type Handler struct {
	paramClient store.KV
	entClient   *ent.Client
	log         *zap.SugaredLogger
	reqeuestCh  chan message.ProblemMessage
}

func NewHandler(
	paramClient store.KV,
	entClient *ent.Client,
	reqCh chan message.ProblemMessage,
) *Handler {
	l := log.NewZap().With("handler", "problem")

	return &Handler{
		paramClient: paramClient,
		entClient:   entClient,
		log:         l,
		reqeuestCh:  reqCh,
	}
}

type GPTOutput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Code        string `json:"code"`
}

func (h *Handler) RequestProblem(ctx context.Context, request gateway.RequestProblemRequestObject) (gateway.RequestProblemResponseObject, error) {
	problemID := nanoid.Must(8)

	// UUID만 담긴 문제 객체를 미리 생성한다.
	// 문제가 아예 없는 경우와, 문제 생성 요청이 들어간 상태를 구분하기 위해서다.
	if err := h.createProblem(ctx, problemID, request); err != nil {
		h.log.Errorw("문제 객체 생성 실패", "error", err)
		return gateway.RequestProblemdefaultJSONResponse{
			Body: gateway.Error{
				Message: serverErrorMessage,
			},
			StatusCode: 500,
		}, nil
	}

	// 문제를 백그라운드에서 생성한다.
	go func() {
		anotherCtx := context.TODO()

		output, err := h.requestProblem(anotherCtx, request)
		if err != nil {
			h.log.Errorw("GPT를 통한 문제 출제 실패", "error", err)
			return
		}

		if err := h.saveProblem(anotherCtx, problemID, output); err != nil {
			h.log.Errorw("문제 저장 실패", "error", err)
			return
		}

		h.reqeuestCh <- message.ProblemMessage{
			ID: problemID,
		}
	}()

	// 사용자에게는 문제 ID를 바로 반환하고, 백그라운드에서 문제를 생성한다.
	return gateway.RequestProblem202JSONResponse{
		ProblemId: problemID,
	}, nil
}

func (h *Handler) requestProblem(ctx context.Context, request gateway.RequestProblemRequestObject) (GPTOutput, error) {
	now := time.Now()

	prompt, err := h.paramClient.Get(ctx, generatingProblemPromptKey)
	if err != nil {
		h.log.Errorw("프롬포트 조회 실패", "key", generatingProblemPromptKey, "error", err)
		return GPTOutput{}, err
	}

	gptClient, err := ai.NewDefaultOpenAI()
	if err != nil {
		h.log.Errorw("GPT 클라이언트 생성 실패", "error", err)
		return GPTOutput{}, err
	}

	prompt = h.injectData(prompt, request)
	gptClient.AddPrompt(prompt)

	result, err := gptClient.Complete(ctx)
	if err != nil {
		h.log.Errorw("GPT의 프롬포트 처리 실패", "type", "gpt", "error", err)
		return GPTOutput{}, err
	}

	var output GPTOutput
	if err := json.Unmarshal([]byte(result), &output); err != nil {
		h.log.Errorw("결과 언마샬 실패", "error", err)
		return GPTOutput{}, err
	}

	gptClient.AddPrompt(retrieveCodePrompt)
	output.Code, err = gptClient.Complete(ctx)
	if err != nil {
		h.log.Errorw("GPT의 프롬포트 처리 실패", "type", "gpt", "error", err)
		return GPTOutput{}, err
	}

	h.log.Infow("문제 생성 완료", "elapsed", time.Since(now).String())

	output.Code = h.extractCode(output.Code)
	output.Code = h.encodeCode(output.Code)

	return output, nil
}

func (h *Handler) injectData(prompt string, req gateway.RequestProblemRequestObject) string {
	prompt = strings.ReplaceAll(prompt, languageTemplateKey, string(req.Body.Language))
	prompt = strings.ReplaceAll(prompt, eloScoreTemplateKey, fmt.Sprintf("%d", req.Body.EloScore))

	return prompt
}

func (h *Handler) extractCode(code string) string {
	return str.ExtractCodeBlocksFromMarkdown(code)[0]
}

func (h *Handler) encodeCode(code string) string {
	encoder := base64.StdEncoding
	return encoder.EncodeToString([]byte(code))
}

func (h *Handler) createProblem(ctx context.Context, uuid string, request gateway.RequestProblemRequestObject) error {
	var difficulty *int
	if request.Body.EloScore != nil {
		elo := int(*request.Body.EloScore)
		difficulty = &elo
	}

	return h.entClient.Problem.Create().SetUUID(uuid).SetLanguage(request.Body.Language).SetNillableDifficulty(difficulty).Exec(ctx)
}

func (h *Handler) saveProblem(ctx context.Context, uuid string, output GPTOutput) error {
	tx, err := h.entClient.Tx(ctx)
	if err != nil {
		h.log.Errorw("failed to create transaction", "category", "db", "error", err)
		return err
	}

	p, err := tx.Problem.Query().ForUpdate().Where(problem.UUID(uuid)).Only(ctx)
	if err != nil {
		h.log.Errorw("failed to get problem", "category", "db", "error", err)
		return err
	}

	err = p.Update().SetTitle(output.Title).SetDescription(output.Description).SetCode(output.Code).Exec(ctx)
	if err != nil {
		h.log.Errorw("failed to update problem", "category", "db", "error", err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		h.log.Errorw("failed to commit transaction", "category", "db", "error", err)
		return err
	}

	return nil
}

func (h *Handler) GetProblem(ctx context.Context, request gateway.GetProblemRequestObject) (gateway.GetProblemResponseObject, error) {
	p, err := h.entClient.Problem.Query().Where(problem.UUID(request.Id)).Only(entcache.WithTTL(ctx, 1*time.Second))
	if err != nil {
		if ent.IsNotFound(err) {
			return gateway.GetProblem404JSONResponse{
				Code:    problemNotFoundCode,
				Message: problemNotFoundMessage,
			}, nil
		}

		h.log.Errorw("failed to get problem", "category", "db", "error", err)
		return serverErrResp, nil
	}

	// 문제 생성 요청은 들어갔지만 아직 문제가 생성되지 않은 경우
	if p.Code == "" || p.Title == "" {
		return gateway.GetProblem409JSONResponse{
			Code:    problemNotReadyCode,
			Message: problemNotReadyMessage,
		}, nil
	}

	return gateway.GetProblem200JSONResponse{
		N200GetProblemJSONResponse: gateway.N200GetProblemJSONResponse{
			Code:        p.Code,
			Id:          p.UUID,
			Title:       p.Title,
			Description: p.Description,
		},
	}, nil
}
