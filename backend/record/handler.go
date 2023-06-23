package record

import (
	"code-connect/ent"
	"code-connect/ent/predicate"
	"code-connect/ent/problem"
	"code-connect/ent/record"
	"code-connect/gateway"
	"code-connect/pkg/store"
	"code-connect/schema/message"
	"context"
	nanoid "github.com/matoous/go-nanoid/v2"
	"go.uber.org/zap"
)

type Handler struct {
	paramClient store.KV
	entClient   *ent.Client
	log         *zap.SugaredLogger
	submitCh    chan message.ProblemMessage
}

type NewHandlerParams struct {
	ParamClient store.KV
	EntClient   *ent.Client
	SubmitCh    chan message.ProblemMessage
}

func NewHandler(params NewHandlerParams) *Handler {
	return &Handler{
		paramClient: params.ParamClient,
		entClient:   params.EntClient,
		submitCh:    params.SubmitCh,
	}
}

func (s *Handler) SubmitSolution(ctx context.Context, request gateway.SubmitSolutionRequestObject) (gateway.SubmitSolutionResponseObject, error) {
	tx, err := s.entClient.Tx(ctx)
	if err != nil {
		s.log.Errorw("트랜잭션 생성 실패", "error", err)
		return gateway.SubmitSolutiondefaultJSONResponse{
			Body: gateway.Error{
				Code:    gateway.ServerError,
				Message: gateway.ServerErrorMessage,
			},
			StatusCode: 500,
		}, err
	}

	p, err := tx.Problem.Query().Where(problem.UUID(request.Body.ProblemId)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return gateway.SubmitSolution404JSONResponse{
				Code:    gateway.ProblemNotFound,
				Message: gateway.ProblemNotFoundMessage,
			}, nil
		}

		s.log.Errorw("문제 조회 실패", "error", err)
		return gateway.SubmitSolutiondefaultJSONResponse{
			Body: gateway.Error{
				Code:    gateway.ServerError,
				Message: gateway.ServerErrorMessage,
			},
			StatusCode: 500,
		}, err
	}

	id := nanoid.Must(8)

	createQuery := tx.Record.Create().SetUUID(id).SetProblem(p).SetCode(request.Body.Code)
	username, ok := store.UsernameFromContext(ctx)
	if ok && username != "" {
		createQuery.SetUserUUID(username)
	}

	err = createQuery.Exec(ctx)
	if err != nil {
		return gateway.SubmitSolutiondefaultJSONResponse{
			Body: gateway.Error{
				Code:    gateway.ServerError,
				Message: gateway.ServerErrorMessage,
			},
			StatusCode: 500,
		}, err
	}

	if err := tx.Commit(); err != nil {
		return gateway.SubmitSolutiondefaultJSONResponse{
			Body: gateway.Error{
				Code:    gateway.ServerError,
				Message: gateway.ServerErrorMessage,
			},
			StatusCode: 500,
		}, err
	}

	s.submitCh <- message.ProblemMessage{
		ID: id,
	}

	return gateway.SubmitSolution202JSONResponse{
		N202SubmitJSONResponse: gateway.N202SubmitJSONResponse{
			RecordId: id,
		},
	}, nil
}

func (s *Handler) GetRecords(ctx context.Context, request gateway.GetRecordsRequestObject) (gateway.GetRecordsResponseObject, error) {
	page := 1
	if request.Params.Page != nil {
		page = int(*request.Params.Page)
	}

	limit := 10
	if request.Params.Limit != nil {
		limit = int(*request.Params.Limit)
	}

	predicates := []predicate.Record{record.HasProblem()}
	if username, ok := store.UsernameFromContext(ctx); ok && username != "" {
		predicates = append(predicates, record.UserUUID(username))
	}

	records, err := s.entClient.Record.Query().Where(predicates...).WithProblem().Offset((page - 1) * limit).Limit(limit).All(ctx)
	if err != nil {
		return gateway.GetRecordsdefaultJSONResponse{
			Body: gateway.Error{
				Code:    gateway.ServerError,
				Message: gateway.ServerErrorMessage,
			},
			StatusCode: 500,
		}, err
	}

	recs := make([]gateway.Record, len(records))
	for i, re := range records {
		recs[i] = gateway.Record{
			Code:         re.Code,
			Efficiency:   gateway.Score(re.Efficiency),
			Id:           re.UUID,
			ProblemId:    re.Edges.Problem.UUID,
			ProblemTitle: re.Edges.Problem.Title,
			Readability:  gateway.Score(re.Readability),
			Robustness:   gateway.Score(re.Robustness),
		}
	}

	return gateway.GetRecords200JSONResponse{
		N200GetRecordsJSONResponse: gateway.N200GetRecordsJSONResponse{
			Records: recs,
		},
	}, nil
}

func (s *Handler) GetRecord(ctx context.Context, request gateway.GetRecordRequestObject) (gateway.GetRecordResponseObject, error) {
	rec, err := s.entClient.Record.Query().Where(record.UUID(request.Id)).WithProblem().Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return gateway.GetRecord404Response{}, nil
		}

		s.log.Errorw("레코드 조회 실패", "error", err)
		return gateway.GetRecorddefaultJSONResponse{
			Body: gateway.Error{
				Code:    gateway.ServerError,
				Message: gateway.ServerErrorMessage,
			},
			StatusCode: 500,
		}, err
	}

	return gateway.GetRecord200JSONResponse{
		N200GetRecordJSONResponse: gateway.N200GetRecordJSONResponse{
			Code:         rec.Code,
			Efficiency:   gateway.Score(rec.Efficiency),
			Id:           rec.UUID,
			ProblemId:    rec.Edges.Problem.UUID,
			ProblemTitle: rec.Edges.Problem.Title,
			Readability:  gateway.Score(rec.Readability),
			Robustness:   gateway.Score(rec.Robustness),
		},
	}, nil
}
