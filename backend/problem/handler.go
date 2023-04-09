package problem

import (
	"code-connect/pkg/ai"
	"code-connect/pkg/log"
	"code-connect/problem/ent"
	"code-connect/problem/ent/problem"
	"context"
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
	NotFound = errors.New("problem not found")
)

func NewProblemHandler(gptClient ai.GPTClient, entClient *ent.Client) *Handler {
	return &Handler{gptClient: gptClient, entClient: entClient}
}

func (h *Handler) Get(ctx context.Context, req GetProblemRequest) (*GetProblemResponse, error) {
	p, err := h.entClient.Problem.Query().Where(problem.UUID(req.ID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, NotFound
		}
		return nil, err
	}

	return &GetProblemResponse{
		Statement: p.Statement,
		Language:  p.Language,
	}, nil
}

func (h *Handler) Create(ctx context.Context, req CreateProblemRequest) (*CreateProblemResponse, error) {
	prompts := []string{
		fmt.Sprintf("Create a coding test in %s to evaluate if candidate is a developer who can write data structures efficiently. The difficulty level is %d out of 100. The higher the difficulty level, the harder the question. Test should include problem statement, example usage, constraints, evaluation criteria only. Test should not give any hints or suggestions. Test should not include heading. Give the output as Markdown format.", req.Language, req.Difficulty),
	}

	id := gonanoid.MustGenerate("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ", 10)

	go func(id string, difficulty int, language string, prompts []string) {
		var (
			tx  *ent.Tx
			err error
		)

		tx, err = h.entClient.Tx(context.TODO())
		if err != nil {
			l.Errorw("error while creating transaction", "error", err)
			return
		}
		defer func() {
			if err == nil {
				return
			}

			if err := tx.Rollback(); err != nil {
				l.Errorw("error while closing transaction", "error", err)
			}
		}()

		var p *ent.Problem
		p, err = tx.Problem.Create().SetUUID(id).SetDifficulty(req.Difficulty).SetLanguage(req.Language).Save(ctx)
		if err != nil {
			l.Errorw("error while creating problem", "error", err)
			return
		}

		var answer string
		answer, err = h.gptClient.CompleteWithContext(ctx, prompts)
		if err != nil {
			l.Errorw("error while completing prompt", "error", err)
		}

		err = tx.Problem.UpdateOne(p).SetStatement(answer).Exec(ctx)
		if err != nil {
			l.Errorw("error while updating problem", "error", err)
			return
		}
	}(id, req.Difficulty, req.Language, prompts)

	return &CreateProblemResponse{
		ID: id,
	}, nil
}
