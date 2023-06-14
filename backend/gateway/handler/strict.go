package handler

import (
	"code-connect/gateway"
	"code-connect/problem"
	"context"
)

type StrictHandler struct {
	problemHandler *problem.Handler
}

func (s *StrictHandler) GetMe(ctx context.Context, request gateway.GetMeRequestObject) (gateway.GetMeResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func NewStrictHandler(problemHandler *problem.Handler) *StrictHandler {
	return &StrictHandler{problemHandler: problemHandler}
}

func (s *StrictHandler) RequestProblem(ctx context.Context, request gateway.RequestProblemRequestObject) (gateway.RequestProblemResponseObject, error) {
	return s.problemHandler.RequestProblem(ctx, request)
}

func (s *StrictHandler) GetProblem(ctx context.Context, request gateway.GetProblemRequestObject) (gateway.GetProblemResponseObject, error) {
	return s.problemHandler.GetProblem(ctx, request)
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
