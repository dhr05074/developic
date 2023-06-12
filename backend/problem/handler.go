package problem

import (
	"code-connect/ent"
	"code-connect/ent/problem"
	"code-connect/gateway"
	"code-connect/pkg/ai"
	"code-connect/pkg/store"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	nanoid "github.com/matoous/go-nanoid/v2"
	"strings"
)

type Handler struct {
	paramClient store.KV
	gptClient   ai.GPTClient
	entClient   *ent.Client
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

func (h *Handler) RequestProblem(ctx context.Context, request gateway.RequestProblemRequestObject) (gateway.RequestProblemResponseObject, error) {
	problemID := nanoid.Must(8)

	// UUID만 담긴 문제 객체를 미리 생성한다.
	// 문제가 아예 없는 경우와, 문제 생성 요청이 들어간 상태를 구분하기 위해서다.
	if err := h.createProblem(ctx, problemID); err != nil {
		return gateway.RequestProblemdefaultJSONResponse{
			Body: struct {
				ErrorCode string `json:"error_code"`
			}{},
			StatusCode: 500,
		}, nil
	}

	go func() {
		anotherCtx := context.TODO()

		output, err := h.requestProblem(anotherCtx, request)
		if err != nil {
			return
		}

		if err := h.saveProblem(anotherCtx, problemID, output); err != nil {
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
		return Output{}, err
	}

	prompt = h.injectData(prompt, request)
	h.gptClient.AddPrompt(prompt)

	result, err := h.gptClient.Complete(ctx)
	if err != nil {
		return Output{}, err
	}

	var output Output
	if err := json.Unmarshal([]byte(result), &output); err != nil {
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

func (h *Handler) createProblem(ctx context.Context, uuid string) error {
	return h.entClient.Problem.Create().SetUUID(uuid).Exec(ctx)
}

func (h *Handler) saveProblem(ctx context.Context, uuid string, output Output) error {
	return h.entClient.Problem.Update().Where(problem.UUID(uuid)).SetTitle(output.Title).SetCode(output.Code).Exec(ctx)
}
