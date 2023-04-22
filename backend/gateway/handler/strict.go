package handler

import (
	"code-connect/gateway"
	"code-connect/pkg/ai"
	"code-connect/problem"
	"context"
	"encoding/base64"
)

type StrictHandler struct {
	gptClient      ai.GPTClient
	problemHandler *problem.Handler
}

func NewStrictHandler(gptClient ai.GPTClient, problemHandler *problem.Handler) *StrictHandler {
	return &StrictHandler{gptClient: gptClient, problemHandler: problemHandler}
}

func (s StrictHandler) CreateProblem(ctx context.Context, request gateway.CreateProblemRequestObject) (gateway.CreateProblemResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s StrictHandler) GetProblem(ctx context.Context, request gateway.GetProblemRequestObject) (gateway.GetProblemResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s StrictHandler) CreateScenario(ctx context.Context, request gateway.CreateScenarioRequestObject) (gateway.CreateScenarioResponseObject, error) {
	decodedQualification, err := base64.StdEncoding.DecodeString(request.Body.Qualifications)
	if err != nil {
		return gateway.CreateScenariodefaultResponse{}, err
	}

	decodedPreferences, err := base64.StdEncoding.DecodeString(request.Body.Preferences)
	if err != nil {
		return gateway.CreateScenariodefaultResponse{}, err
	}

	req := problem.CreateScenarioRequest{
		Qualification: string(decodedQualification),
		Preferences:   string(decodedPreferences),
	}
	resp, err := s.problemHandler.CreateScenario(ctx, req)
	if err != nil {
		return gateway.CreateScenariodefaultResponse{}, err
	}

	return gateway.CreateScenario202JSONResponse{RequestId: resp.RequestID}, nil
}

func (s StrictHandler) GetScenariosByRequestID(ctx context.Context, request gateway.GetScenariosByRequestIDRequestObject) (gateway.GetScenariosByRequestIDResponseObject, error) {
	resp, err := s.problemHandler.ScenariosByRequestID(ctx, request.RequestId)
	if err != nil {
		return gateway.GetScenariosByRequestIDdefaultResponse{}, err
	}

	scenarios := make([]gateway.Scenario, 0, len(resp))
	for _, scenario := range resp {
		scenarios = append(scenarios, gateway.Scenario{
			Id:      scenario.ID,
			Title:   scenario.Title,
			Content: scenario.Content,
		})
	}

	return gateway.GetScenariosByRequestID200JSONResponse{Scenarios: scenarios}, nil
}
