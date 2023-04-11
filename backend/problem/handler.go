package problem

import (
	"code-connect/pkg/ai"
	"code-connect/pkg/log"
	"code-connect/problem/ent"
	"code-connect/problem/ent/problem"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

var l = log.NewZapSugaredLogger()

type CreateProblemRequest struct {
	Difficulty int    `json:"difficulty"`
	Language   string `json:"language"`
}

type CreateProblemResponse struct {
	ID string `json:"id"`
}

type GetProblemRequest struct {
	ID string `json:"id"`
}

type GetProblemResponse struct {
	Statement string `json:"statement"`
	Language  string `json:"language"`
}

type Handler struct {
	gptClient ai.GPTClient
	entClient *ent.Client
}

var (
	NotFound      = errors.New("problem not found")
	DatabaseError = errors.New("database error")
)

const (
	CodeAlphabets = "0123456789abcdefghijklmnopqrstuvwxyz"
)

func NewHandler(gptClient ai.GPTClient, entClient *ent.Client) *Handler {
	return &Handler{gptClient: gptClient, entClient: entClient}
}

func (h *Handler) Get(ctx context.Context, req GetProblemRequest) (*GetProblemResponse, error) {
	p, err := h.entClient.Problem.Query().Where(problem.UUID(req.ID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, NotFound
		}
		return nil, DatabaseError
	}

	return &GetProblemResponse{
		Statement: p.Statement,
		Language:  p.Language,
	}, nil
}

func (h *Handler) Create(ctx context.Context, req CreateProblemRequest) (*CreateProblemResponse, error) {
	id := gonanoid.MustGenerate(CodeAlphabets, 10)

	p, err := h.entClient.Problem.Create().SetUUID(id).SetDifficulty(req.Difficulty).SetLanguage(req.Language).Save(ctx)
	if err != nil {
		l.Errorw("error while creating problem", "error", err)
		return nil, DatabaseError
	}

	prompts := []string{
		fmt.Sprintf("Create a coding test in %s to evaluate if candidate is a developer who can use data structures efficiently. The difficulty level is %d out of 100. The higher the difficulty level, the harder the question. Test should include problem statement, example usage, constraints, evaluation criteria only. Test should not give any hints or suggestions. Test should not include heading. If there is a difficult concept in the question, add a brief explanation. Give the output as Markdown format.", req.Language, req.Difficulty),
	}

	go func(p *ent.Problem, prompts []string) {
		newCtx := context.TODO()

		l.Info("generating problem statement...")
		answer, err := h.gptClient.CompleteWithContext(newCtx, prompts)
		if err != nil {
			l.Errorw("error while completing prompt", "error", err)
			return
		}
		l.Info("problem statement generated")

		// encode answer to base64 to avoid issues with special characters
		encodedAnswer := base64.StdEncoding.EncodeToString([]byte(answer))
		err = h.entClient.Problem.UpdateOne(p).SetStatement(encodedAnswer).Exec(newCtx)
		if err != nil {
			l.Errorw("error while updating problem", "error", err)
			return
		}
	}(p, prompts)

	return &CreateProblemResponse{
		ID: id,
	}, nil
}
