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
	"errors"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"regexp"
	"strconv"
	"strings"
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

		prompt := "Context: Select a real-world situation that illustrates the function of a specific application or system. This should be relevant and valuable to developers, offering them a practical understanding of the application or system.\n\nGenerate Unoptimized Code: Write code in {{programming_language}} that corresponds to the context you have chosen. The code should function correctly but exhibit typical signs of inefficiency, such as repetitive code, lengthy methods, large classes, and poorly chosen identifiers.\n\nCreate a Refactoring Task: Combine the above steps to create a refactoring task, assigning it a difficulty level of {{difficulty}} on a scale from 1 to 100. Please provide your response in JSON format with the following structure:\n\n```json\n{\n  \"title\": \"<title_of_the_task>\",\n  \"background\": \"<description_of_the_context>\",\n  \"target_code\": \"<unoptimized_code>\",\n  \"estimated_time\": <time_in_minutes>\n}\n"
		prompt = strings.ReplaceAll(prompt, "{{programming_language}}", req.Body.Language)
		difficulty := 50
		if req.Body.Difficulty != nil {
			difficulty = *req.Body.Difficulty
		}
		prompt = strings.ReplaceAll(prompt, "{{difficulty}}", strconv.Itoa(difficulty))
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

func (h *Handler) SubmitCode(ctx context.Context, req gateway.SubmitSolutionRequestObject) (gateway.SubmitSolutionResponseObject, error) {
	// Decode the submitted code.
	submittedCode, err := base64.StdEncoding.DecodeString(req.Body.Code)
	if err != nil {
		return nil, err
	}

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

	problem, err := tx.Problem.Query().ForUpdate().Where(problem2.UUID(req.RequestId)).Only(timeout)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.New("problem not found. check your request id. or wait a minute and try again.")
		}

		return nil, err
	}

	uuid := h.generateID()

	err = tx.Submission.Create().
		SetUUID(uuid).
		SetProblem(problem).
		SetCode(string(submittedCode)).
		SetSubmitterID(h.generateID()).
		SetProblemID(problem.ID).
		Exec(timeout)
	if err != nil {
		return nil, err
	}

	return gateway.SubmitSolution200JSONResponse{
		SubmissionId: uuid,
	}, nil
}
