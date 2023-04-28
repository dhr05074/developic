package problem

import (
	"code-connect/gateway"
	"code-connect/pkg/ai"
	"code-connect/pkg/log"
	"code-connect/problem/ent"
	problem2 "code-connect/problem/ent/problem"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/go-resty/resty/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"os"
)

const codeAlphabets = "0123456789abcdefghijklmnopqrstuvwxyz"

type Handler struct {
	gptClient ai.GPTClient
	entClient *ent.Client
}

func NewHandler(gptClient ai.GPTClient, entClient *ent.Client) *Handler {
	return &Handler{gptClient: gptClient, entClient: entClient}
}

var l = log.NewZapSugaredLogger()

type TranslateRequest struct {
	Text   string `json:"text"`
	Source string `json:"source"`
	Target string `json:"target"`
}

type TranslateResponse struct {
	Text string `json:"text"`
}

func (h *Handler) translateRequirements(ctx context.Context, requirements, source, target string) (result string, err error) {
	client := resty.New().SetBaseURL("https://deepl-translator.p.rapidapi.com").R().SetContext(ctx)
	client.SetHeader("X-RapidAPI-Key", os.Getenv("DEEPL_API_KEY"))
	client.SetHeader("X-RapidAPI-Host", os.Getenv("DEEPL_API_HOST"))

	body := TranslateRequest{
		Text:   requirements,
		Source: source,
		Target: target,
	}

	var resp TranslateResponse
	_, err = client.SetBody(body).SetResult(&resp).Post("/translate")
	if err != nil {
		return "", err
	}

	return resp.Text, nil
}

func (h *Handler) CreateProblem(ctx context.Context, req gateway.CreateProblemRequestObject) (*gateway.CreateProblem202JSONResponse, error) {
	requestID := h.generateID()

	go func(entClient *ent.Client, requestID string) {
		newCtx := context.TODO()
		newGPTClient := h.gptClient.NewClientWithEmptyContext()

		prompts := make([]string, 0, 2)
		prompts = append(prompts, fmt.Sprintf("Please generate a code refactoring challenge problem that incorporates the characteristics of dirty code mentioned below. The challenge should be tailored to the following user preferences:"))
		prompts = append(prompts, fmt.Sprintf("Difficulty score: %d", req.Body.Difficulty))
		prompts = append(prompts, fmt.Sprintf("Programming language: %s", req.Body.Language))

		characteristics := h.bulletPointList([]string{
			"Lack of structure and organization",
			"Inconsistency in naming conventions, indentation, and formatting",
			"No modularity, with mixed functionalities and concerns",
			"Code duplication and redundancy",
			"Poor documentation and missing comments",
			"Overly complex logic and nested control structures",
			"Hardcoded values",
			"Use of global variables and side effects",
			"Inadequate error handling",
			"Inefficient algorithms or suboptimal performance"})
		prompts = append(prompts, "Characteristics of dirty code:")
		prompts = append(prompts, characteristics)
		prompts = append(prompts, "Note: The challenge problem should be provided without any answer code. Just give the code which is include main function and instruction and title of the problem.")

		res, err := newGPTClient.CompleteWithContext(newCtx, prompts)
		if err != nil {
			l.Errorw("error while completing prompt", "error", err)
			return
		}

		l.Infof("problem created: %s\n", res)

		tx, err := entClient.Tx(newCtx)
		if err != nil {
			l.Errorw("error while creating transaction", "error", err)
			return
		}

		defer func() {
			if err == nil {
				return
			}

			if err := tx.Rollback(); err != nil {
				l.Errorw("error while rolling back transaction", "error", err)
			}
		}()

		content := base64.StdEncoding.EncodeToString([]byte(res))
		if err != nil {
			l.Errorw("error while decoding problem content", "error", err)
			return
		}

		problemID := h.generateID()
		err = tx.Problem.Create().SetUUID(problemID).SetRequestID(requestID).SetContent(content).Exec(newCtx)
		if err != nil {
			l.Errorw("error while creating problem", "error", err)
			return
		}

		if err = tx.Commit(); err != nil {
			l.Errorw("error while committing transaction", "error", err)
			return
		}

	}(h.entClient, requestID)

	return &gateway.CreateProblem202JSONResponse{
		RequestId: requestID,
	}, nil
}

func (h *Handler) GetProblem(ctx context.Context, req gateway.GetProblemRequestObject) (*gateway.GetProblem200JSONResponse, error) {
	tx, err := h.entClient.Tx(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err == nil {
			return
		}

		if err := tx.Rollback(); err != nil {
			l.Errorw("error while rolling back transaction", "error", err)
		}
	}()

	problem, err := tx.Problem.Query().ForUpdate().Where(problem2.RequestID(req.RequestId)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, err
		}

		return nil, err
	}

	if err = tx.Commit(); err != nil {
		l.Errorw("error while committing transaction", "error", err)
		return nil, err
	}

	return &gateway.GetProblem200JSONResponse{
		Problem: gateway.Problem{
			Content:   problem.Content,
			ProblemId: problem.UUID,
			Title:     problem.Title,
		},
	}, nil
}

func (h *Handler) EvaluateSolution(ctx context.Context, req gateway.EvaluateSolutionRequestObject) (*gateway.EvaluateSolution200JSONResponse, error) {
	before, err := base64.StdEncoding.DecodeString(req.Body.Before)
	if err != nil {
		return nil, err
	}

	after, err := base64.StdEncoding.DecodeString(req.Body.After)
	if err != nil {
		return nil, err
	}

	prompts := make([]string, 0, 2)
	prompts = append(prompts, fmt.Sprintf("Compare the before and after code and give me a score out of 100 on how well the refactoring went. Just give me the score only."))
	prompts = append(prompts, fmt.Sprintf("Before:\n%s\nAfter:\n%s", before, after))
	prompts = append(prompts, "Score:")

	res, err := h.gptClient.NewClientWithEmptyContext().CompleteWithContext(ctx, prompts)
	if err != nil {
		return nil, err
	}

	return &gateway.EvaluateSolution200JSONResponse{
		Score: res,
	}, nil
}

func (h *Handler) bulletPointList(arr []string) string {
	res := ""
	for _, v := range arr {
		res += "- " + v + "\n"
	}
	return res
}

func (h *Handler) generateID() string {
	return gonanoid.MustGenerate(codeAlphabets, 7)
}
