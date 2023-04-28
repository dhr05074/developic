package handler

import (
	"code-connect/gateway"
	"code-connect/pkg/ai"
	"code-connect/problem"
	"context"
)

type StrictHandler struct {
	gptClient      ai.GPTClient
	problemHandler *problem.Handler
}

func NewStrictHandler(gptClient ai.GPTClient, problemHandler *problem.Handler) *StrictHandler {
	return &StrictHandler{gptClient: gptClient, problemHandler: problemHandler}
}

func (s *StrictHandler) CreateProblem(ctx context.Context, request gateway.CreateProblemRequestObject) (gateway.CreateProblemResponseObject, error) {
	return s.problemHandler.CreateProblem(ctx, request)
}

func (s *StrictHandler) GetProblem(ctx context.Context, request gateway.GetProblemRequestObject) (gateway.GetProblemResponseObject, error) {
	return s.problemHandler.GetProblem(ctx, request)
}

func (s *StrictHandler) EvaluateSolution(ctx context.Context, request gateway.EvaluateSolutionRequestObject) (gateway.EvaluateSolutionResponseObject, error) {
	return s.problemHandler.EvaluateSolution(ctx, request)
}
