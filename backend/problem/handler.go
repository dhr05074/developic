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
	"regexp"
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

		prompt := fmt.Sprintf("Define a Context: Pinpoint a specific, real-world instance that exemplifies the functioning of a distinct application or system. This instance should be impactful for developers, offering them a practical understanding of the application or system.\n\nGenerate Unoptimized Code: Craft the code in %s, that corresponds to your defined instance. This code should be working, but it shouldn't be optimized or efficient. It should display a range of typical issues that signal a need for refactoring, including code repetition, lengthy methods, oversized classes, and poorly selected identifiers. Make sure the code is devoid of syntax errors or bugs that hinder its operation, as the aim here is to enhance the structure, not to rectify broken code.\n\nIntegrate steps 1 and 2 to create a refactoring task. The difficulty level of this task should correspond to difficulty level %d on a scale of 1 to 100. Your response should be provided exclusively in the JSON format. The JSON object should be structured as follows: {title, background, target_code, estimated_time}. The 'estimated_time' should be stipulated in minutes.", req.Body.Difficulty, req.Body.Language)
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
			l.Errorw("error while unmarshalling json", "error", err)
			return
		}

		dur := time.Since(startTime)
		l.Infof("problem created in %f seconds", dur.Seconds())

		// If the testCode contents are returned as they are, there will be an escape problem, so encode with Base64 before saving.
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

		pr, err := tx.Problem.Create().SetUUID(problemID).SetRequestID(requestID).SetTitle(output.Title).SetBackground(output.Background).SetCode(output.TargetCode).SetLanguage(req.Body.Language).SetEstimatedTime(output.EstimatedTime).Save(newCtx)
		if err != nil {
			l.Errorw("error while creating problem", "error", err)
			return
		}

		if err = tx.Commit(); err != nil {
			l.Errorw("error while committing transaction", "error", err)
			return
		}

		// Creating test testCode for the testCode.
		newGPTClient.AddPrompt("Create a test code for the previous code. It should have 5 test cases.")
		testCode, err := newGPTClient.Complete(newCtx)
		if err != nil {
			l.Errorw("error while completing prompt", "error", err)
			return
		}

		l.Info(testCode)

		_, testCode = extractCode(testCode)

		// If the testCode contents are returned as they are, there will be an escape problem, so encode with Base64 before saving.
		testCode = base64.StdEncoding.EncodeToString([]byte(testCode))
		if err != nil {
			l.Errorw("error while decoding problem content", "error", err)
			return
		}

		tx, err = entClient.Tx(newCtx)
		if err != nil {
			l.Errorw("error while starting transaction", "error", err)
			return
		}

		err = tx.Problem.UpdateOne(pr).SetTestCode(testCode).Exec(newCtx)
		if err != nil {
			l.Errorw("error while updating problem", "error", err)
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

func extractCode(markdown string) (string, string) {
	regex := regexp.MustCompile("`{3}(.+?)\n(.+?)\n`{3}")
	match := regex.FindStringSubmatch(markdown)
	if len(match) >= 3 {
		return match[1], match[2]
	}
	return "", ""
}
