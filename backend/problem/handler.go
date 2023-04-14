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
	nanoid "github.com/matoous/go-nanoid/v2"
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
	CodeAlphabets = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func NewHandler(gptClient ai.GPTClient, entClient *ent.Client) *Handler {
	return &Handler{gptClient: gptClient, entClient: entClient}
}

func (h *Handler) Get(ctx context.Context, req GetProblemRequest) (*GetProblemResponse, error) {
	tx, err := h.entClient.Tx(ctx)
	if err != nil {
		l.Errorw("error while creating transaction", "error", err)
		return nil, DatabaseError
	}

	p, err := tx.Problem.Query().Where(problem.UUID(req.ID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, NotFound
		}
		return nil, DatabaseError
	}

	if err := tx.Commit(); err != nil {
		l.Errorw("error while committing transaction", "error", err)
		return nil, DatabaseError
	}

	return &GetProblemResponse{
		Statement: p.Statement,
		Language:  p.Language,
	}, nil
}

func (h *Handler) Create(_ context.Context, req CreateProblemRequest) (*CreateProblemResponse, error) {
	id := nanoid.MustGenerate(CodeAlphabets, 10)

	prompts := []string{
		fmt.Sprintf("Create a coding test in %s to evaluate if candidate is a developer who can use data structures efficiently.", req.Language),
		fmt.Sprintf("The difficulty level is %d out of 100. The higher the difficulty level, the harder the problem.", req.Difficulty),
		fmt.Sprintf("Test should include problem statement, example usage, constraints, evaluation criteria only. Test should not give any hints or suggestions. Test should not include heading. If there is a difficult concept in the question, add a brief explanation. Give the output as Markdown format."),
	}

	go func(prompts []string) {
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

		tx, err := h.entClient.Tx(newCtx)
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

		_, err = tx.Problem.Create().SetStatement(encodedAnswer).SetUUID(id).SetDifficulty(req.Difficulty).SetLanguage(req.Language).Save(newCtx)
		if err != nil {
			l.Errorw("error while creating problem", "error", err)
			return
		}

		if err = tx.Commit(); err != nil {
			l.Errorw("error while committing transaction", "error", err)
			return
		}
	}(prompts)

	return &CreateProblemResponse{
		ID: id,
	}, nil
}
