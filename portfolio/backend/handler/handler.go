package handler

import (
	"context"
	"encoding/json"
	gonanoid "github.com/matoous/go-nanoid"
	"portfolio/ent"
	"portfolio/ent/report"
	"portfolio/gateway"
	"portfolio/schema"
	"regexp"
)

type Strict struct {
	ai       *OpenAI
	db       *ent.Client
	putQueue chan PutQueueRequest
	// Key: requestID, Value: status
	queue map[string]gateway.ReportStatus
}

type PutQueueRequest struct {
	RequestID string               `json:"request_id"`
	Status    gateway.ReportStatus `json:"status"`
}

func NewStrict(ai *OpenAI, db *ent.Client) *Strict {
	queue := make(map[string]gateway.ReportStatus)
	putQueue := make(chan PutQueueRequest, 10)
	go func() {
		for request := range putQueue {
			queue[request.RequestID] = request.Status
		}
	}()

	return &Strict{ai: ai, db: db, putQueue: putQueue, queue: queue}
}

func (s *Strict) GetReportByRequestID(ctx context.Context, request gateway.GetReportByRequestIDRequestObject) (gateway.GetReportByRequestIDResponseObject, error) {
	status, ok := s.queue[request.RequestId]
	if !ok {
		return gateway.GetReportByRequestID404JSONResponse{
			Message: "request not found",
		}, nil
	}

	if status != gateway.Done {
		return gateway.GetReportByRequestID200JSONResponse{
			ProjectFeedbacks:         nil,
			ProjectRecommendations:   nil,
			Status:                   status,
			TechStackFeedbacks:       nil,
			TechStackRecommendations: nil,
		}, nil
	}

	r, err := s.db.Report.Query().Where(report.RequestID(request.RequestId)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return gateway.GetReportByRequestID404JSONResponse{
				Message: "report not found",
			}, nil
		}

		return nil, err
	}

	return gateway.GetReportByRequestID200JSONResponse{
		ProjectFeedbacks:         r.ProjectFeedbacks,
		ProjectRecommendations:   r.ProjectRecommendations,
		Status:                   gateway.ReportStatus(r.Status),
		TechStackFeedbacks:       r.TechStackFeedbacks,
		TechStackRecommendations: r.TechStackRecommendations,
	}, nil
}

func (s *Strict) SubmitPortfolio(ctx context.Context, request gateway.SubmitPortfolioRequestObject) (gateway.SubmitPortfolioResponseObject, error) {
	var projects = make([]schema.Project, len(request.Body.Projects))
	for i, project := range request.Body.Projects {
		projects[i] = schema.Project{
			Title:       project.Title,
			Description: project.Description,
		}
	}

	params := Parameters{
		Job:              request.Body.Job,
		CareerYears:      uint(request.Body.CareerYears),
		TechStacks:       request.Body.TechStacks,
		Projects:         projects,
		PreferredCompany: request.Body.PreferredCompany,
	}

	s.ai.AddPrompt(makePrompt(params))

	requestID := gonanoid.MustGenerate("0123456789abcdefghijklmnopqrstuvwxyz", 10)
	s.putQueue <- PutQueueRequest{
		RequestID: requestID,
		Status:    gateway.Pending,
	}

	go func() {
		var err error
		defer func() {
			if err != nil {
				s.putQueue <- PutQueueRequest{
					RequestID: requestID,
					Status:    gateway.Failed,
				}
			}
		}()

		s.putQueue <- PutQueueRequest{
			RequestID: requestID,
			Status:    gateway.Processing,
		}

		result, err := s.ai.Complete(ctx, 0.8)
		if err != nil {
			return
		}

		_, jsonCode := s.extractCodeInMarkdownDocument(result)
		if jsonCode == "" {
			return
		}

		var output schema.Report
		if err = json.Unmarshal([]byte(jsonCode), &output); err != nil {
			return
		}

		err = s.db.Report.Create().SetRequestID(requestID).SetProjectFeedbacks(output.ProjectFeedbacks).SetTechStackFeedbacks(output.TechStackFeedbacks).SetProjectRecommendations(output.ProjectRecommendations).SetTechStackRecommendations(output.TechStackRecommendations).Exec(ctx)
		if err != nil {
			return
		}

		s.putQueue <- PutQueueRequest{
			RequestID: requestID,
			Status:    gateway.Done,
		}
	}()

	return gateway.SubmitPortfolio202JSONResponse{
		RequestId: requestID,
	}, nil
}

func (s *Strict) extractCodeInMarkdownDocument(markdown string) (language string, code string) {
	regex := regexp.MustCompile("`{3}(.+?)\n(.+?)\n`{3}")
	match := regex.FindStringSubmatch(markdown)
	if len(match) >= 3 {
		return match[1], match[2]
	}
	return "", ""
}
