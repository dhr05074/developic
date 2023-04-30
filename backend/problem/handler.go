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
	gonanoid "github.com/matoous/go-nanoid/v2"
)

const codeAlphabets = "0123456789abcdefghijklmnopqrstuvwxyz"

type Handler struct {
	gptClient ai.GPTClient
	entClient *ent.Client
}

func NewHandler(gptClient ai.GPTClient, entClient *ent.Client) *Handler {
	return &Handler{gptClient: gptClient, entClient: entClient}
}

var l = log.NewZap()

type Prompts struct {
	prompts []string
}

func (p *Prompts) Add(format string, args ...any) {
	p.prompts = append(p.prompts, fmt.Sprintf(format, args...))
}

func (p *Prompts) StrArray() []string {
	return p.prompts
}

func (h *Handler) wrapTransaction(ctx context.Context, cli *ent.Client, fn func(tx *ent.Tx) error) error {
	var err error

	tx, err := cli.Tx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err == nil {
			return
		}

		if err := tx.Rollback(); err != nil {
			l.Errorw("error while rolling back transaction", "error", err)
		}
	}()

	if err = fn(tx); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		l.Errorw("error while committing transaction", "error", err)
		return err
	}

	return nil
}

func (h *Handler) CreateProblem(ctx context.Context, req gateway.CreateProblemRequestObject) (*gateway.CreateProblem202JSONResponse, error) {
	// When searching for problems created later, you can search by Request ID.
	requestID := h.generateID()

	go func(entClient *ent.Client, requestID string) {
		newCtx := context.TODO()
		newGPTClient := h.gptClient.NewContext()

		var prompts Prompts
		prompts.Add("Please create a refactoring problem with a title and scratch code in the specified programming language. The code should have a main function, some additional functions, and exhibit characteristics that require refactoring based on the specified difficulty level. Do not include the problem statement or the refactored solution.")
		prompts.Add("Difficulty: %d/100", req.Body.Difficulty)
		prompts.Add("Language: %s", req.Body.Language)
		prompts.Add("Example Output: \"title\": \"<title>\", \"scratch_code\": \"<scratch_code>\"")

		res, err := newGPTClient.CompleteWithContext(newCtx, prompts.StrArray())
		if err != nil {
			l.Errorw("error while completing prompt", "error", err)
			return
		}

		l.Infof("problem created: %s\n", res)

		// If the code contents are returned as they are, there will be an escape problem, so encode with Base64 before saving.
		content := base64.StdEncoding.EncodeToString([]byte(res))
		if err != nil {
			l.Errorw("error while decoding problem content", "error", err)
			return
		}
		problemID := h.generateID()

		if err := h.wrapTransaction(context.TODO(), entClient, func(tx *ent.Tx) error {
			return tx.Problem.Create().SetUUID(problemID).SetRequestID(requestID).SetContent(content).Exec(newCtx)
		}); err != nil {
			l.Errorw("error while creating problem", "error", err)
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

	res, err := h.gptClient.NewContext().CompleteWithContext(ctx, prompts)
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
