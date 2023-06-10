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

func (s StrictHandler) PostProblems(ctx context.Context, request gateway.PostProblemsRequestObject) (gateway.PostProblemsResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s StrictHandler) GetProblemsId(ctx context.Context, request gateway.GetProblemsIdRequestObject) (gateway.GetProblemsIdResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s StrictHandler) GetRecords(ctx context.Context, request gateway.GetRecordsRequestObject) (gateway.GetRecordsResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s StrictHandler) GetRecordsId(ctx context.Context, request gateway.GetRecordsIdRequestObject) (gateway.GetRecordsIdResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s StrictHandler) PostSubmit(ctx context.Context, request gateway.PostSubmitRequestObject) (gateway.PostSubmitResponseObject, error) {
	//TODO implement me
	panic("implement me")
}
