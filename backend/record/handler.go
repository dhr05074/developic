package record

import (
	"code-connect/ent"
	"code-connect/ent/problem"
	"code-connect/gateway"
	"code-connect/pkg/ai"
	"code-connect/pkg/store"
	"code-connect/schema/message"
	"context"
	nanoid "github.com/matoous/go-nanoid/v2"
	"go.uber.org/zap"
)

type Handler struct {
	paramClient store.KV
	gptClient   ai.GPTClient
	entClient   *ent.Client
	log         *zap.SugaredLogger
	submitCh    chan message.ProblemMessage
}

type NewHandlerParams struct {
	ParamClient store.KV
	GptClient   ai.GPTClient
	EntClient   *ent.Client
	SubmitCh    chan message.ProblemMessage
}

func NewHandler(params NewHandlerParams) *Handler {
	return &Handler{
		paramClient: params.ParamClient,
		gptClient:   params.GptClient,
		entClient:   params.EntClient,
		submitCh:    params.SubmitCh,
	}
}

func (s *Handler) SubmitSolution(ctx context.Context, request gateway.SubmitSolutionRequestObject) (gateway.SubmitSolutionResponseObject, error) {
	tx, err := s.entClient.Tx(ctx)
	if err != nil {
		return gateway.SubmitSolutiondefaultJSONResponse{
			Body: gateway.Error{
				Message: err.Error(),
			},
			StatusCode: 500,
		}, err
	}

	p, err := tx.Problem.Query().Where(problem.UUID(request.Body.ProblemId)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return gateway.SubmitSolutiondefaultJSONResponse{
				Body: gateway.Error{
					Message: err.Error(),
				},
				StatusCode: 404,
			}, err
		}
		return gateway.SubmitSolutiondefaultJSONResponse{
			Body: gateway.Error{
				Message: err.Error(),
			},
			StatusCode: 500,
		}, err
	}

	id := nanoid.Must(8)

	err = tx.Record.Create().SetUUID(id).AddProblem(p).SetCode(request.Body.Code).Exec(ctx)
	if err != nil {
		return gateway.SubmitSolutiondefaultJSONResponse{
			Body: gateway.Error{
				Message: err.Error(),
			},
			StatusCode: 500,
		}, err
	}

	if err := tx.Commit(); err != nil {
		return gateway.SubmitSolutiondefaultJSONResponse{
			Body: gateway.Error{
				Message: err.Error(),
			},
			StatusCode: 500,
		}, err
	}

	s.submitCh <- message.ProblemMessage{
		ID: id,
	}

	return gateway.SubmitSolution202JSONResponse{}, nil
}

func (s *Handler) GetRecords(ctx context.Context, request gateway.GetRecordsRequestObject) (gateway.GetRecordsResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Handler) GetRecord(ctx context.Context, request gateway.GetRecordRequestObject) (gateway.GetRecordResponseObject, error) {
	//TODO implement me
	panic("implement me")
}
