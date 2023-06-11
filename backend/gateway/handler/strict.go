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

func (s *StrictHandler) RequestProblem(ctx context.Context, request gateway.RequestProblemRequestObject) (gateway.RequestProblemResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s *StrictHandler) GetProblem(ctx context.Context, request gateway.GetProblemRequestObject) (gateway.GetProblemResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s *StrictHandler) GetRecords(ctx context.Context, request gateway.GetRecordsRequestObject) (gateway.GetRecordsResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s *StrictHandler) GetRecord(ctx context.Context, request gateway.GetRecordRequestObject) (gateway.GetRecordResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s *StrictHandler) SubmitSolution(ctx context.Context, request gateway.SubmitSolutionRequestObject) (gateway.SubmitSolutionResponseObject, error) {
	//TODO implement me
	panic("implement me")
}
