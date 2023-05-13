package problem

import (
	"code-connect/gateway"
	"code-connect/pkg/ai"
	"code-connect/pkg/log"
	"code-connect/problem/ent"
	problem2 "code-connect/problem/ent/problem"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"time"
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

type Output struct {
	Title         string `json:"title"`
	Background    string `json:"background"`
	TargetCode    string `json:"target_code"`
	EstimatedTime int    `json:"estimated_time"`
}

func (h *Handler) CreateProblem(_ context.Context, req gateway.CreateProblemRequestObject) (*gateway.CreateProblem202JSONResponse, error) {
	// When searching for problems created later, you can search by Request ID.
	requestID := h.generateID()

	go func(entClient *ent.Client, requestID string) {
		newCtx := context.TODO()
		newGPTClient := h.gptClient.NewContext()

		prompt := fmt.Sprintf("Imagine you are a backend developer at an IT company and create two key components of a code refactoring challenge. \n\n1. create a specific background scenario for the code to be refactored.\n\n2. create a piece of code that follows the previous scenario and would take a new developer %d minutes to fix. Use %s as your programming language.\n\n[Result format].\n{\"title\":\"\",\n\"background:\"\",\n\"target_code\":\"\",\n\"estimated_time\":30\n}", req.Body.EstimatedTime, req.Body.Language)
		newGPTClient.AddPrompt(prompt)

		startTime := time.Now()

		msg, err := newGPTClient.Complete(newCtx)
		if err != nil {
			l.Errorw("error while completing prompt", "error", err)
			return
		}

		l.Info(msg)

		var output Output
		if err = json.Unmarshal([]byte(msg), &output); err != nil {
			l.Errorw("error while unmarshalling output", "error", err)
			return
		}

		dur := time.Since(startTime)
		l.Infof("problem created in %f seconds", dur.Seconds())

		// If the code contents are returned as they are, there will be an escape problem, so encode with Base64 before saving.
		output.TargetCode = base64.StdEncoding.EncodeToString([]byte(output.TargetCode))
		if err != nil {
			l.Errorw("error while decoding problem content", "error", err)
			return
		}
		problemID := h.generateID()

		tx, err := entClient.Tx(newCtx)
		if err != nil {
			l.Errorw("error while starting transaction", "error", err)
			return
		}

		err = tx.Problem.Create().SetUUID(problemID).SetRequestID(requestID).SetTitle(output.Title).SetBackground(output.Background).SetCode(output.TargetCode).SetLanguage(req.Body.Language).SetNillableEstimatedTime(req.Body.EstimatedTime).Exec(newCtx)
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

func (h *Handler) GetProblem(ctx context.Context, req gateway.GetProblemRequestObject) (gateway.GetProblemResponseObject, error) {
	timeout, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	tx, err := h.entClient.Tx(timeout)
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

	problem, err := tx.Problem.Query().ForUpdate().Where(problem2.RequestID(req.RequestId)).Only(timeout)
	if err != nil {
		if ent.IsNotFound(err) {
			return gateway.GetProblem404JSONResponse{Message: "problem not found. check your request id. or wait a minute and try again."}, nil
		}

		return nil, err
	}

	if err = tx.Commit(); err != nil {
		l.Errorw("error while committing transaction", "error", err)
		return nil, err
	}

	return &gateway.GetProblem200JSONResponse{
		Problem: gateway.Problem{
			Background:    problem.Background,
			Code:          problem.Code,
			EstimatedTime: problem.EstimatedTime,
			Title:         problem.Title,
			ProblemId:     problem.UUID,
		},
	}, nil
}

func (h *Handler) generateID() string {
	return gonanoid.MustGenerate(codeAlphabets, 7)
}
