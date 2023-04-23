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

		decodedRequirements, err := base64.StdEncoding.DecodeString(req.Body.Requirements)
		if err != nil {
			l.Errorw("error while decoding requirements", "error", err)
			return
		}
		l.Infof("decoded requirements: %s\n", decodedRequirements)

		requirements, err := h.translateRequirements(newCtx, string(decodedRequirements), "KO", "EN")
		if err != nil {
			l.Errorw("error while translating requirements", "error", err)
			return
		}
		l.Infof("translated requirements: %s\n", requirements)

		prompts := make([]string, 0, 2)
		prompts = append(prompts, requirements)
		prompts = append(prompts, "Based on the data above, fill in the blanks below.\n\nTechnical Stack:\nProgramming languages, frameworks, and tools: \nAdditional technologies: \n  Role and Experience Level:\n  Target role: \nDesired experience level: \nKey Skills and Concepts: \nDomain-specific Knowledge: \nReal-world Problem-solving: \nTime Constraints: \nEvaluation Criteria: \nCollaboration and Communication Skills: \n")

		res, err := newGPTClient.CompleteWithContext(newCtx, prompts)
		if err != nil {
			l.Errorw("error while completing prompt", "error", err)
			return
		}

		l.Infof("analysis created: %s\n", res)

		prompts = make([]string, 0, 1)
		prompts = append(prompts, fmt.Sprintf("We request the creation of a custom coding challenge, targeting the role mentioned eariler. The challenge should be in Markdown format, using '##' for headings and '-' for details. Please develop three tasks focusing on key skills, domain-specific knowledge, and real-world problem-solving scenarios, incorporating the specified programming languages, frameworks, and tools. Adhere to the time limit and ensure the tasks align with the evaluation criteria provided by the company."))

		res, err = newGPTClient.CompleteWithContext(newCtx, prompts)
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

		prompts = make([]string, 0, 1)
		prompts = append(prompts, "Translate it to Korean")

		res, err = newGPTClient.CompleteWithContext(newCtx, prompts)
		if err != nil {
			l.Errorw("error while completing prompt", "error", err)
			return
		}

		l.Infof("problem content translated to Korean: %s\n", res)

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
