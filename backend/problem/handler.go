package problem

import (
	"code-connect/gateway"
	"code-connect/pkg/aws"
	"context"
)

type Handler struct {
	paramClient *aws.SSMClient
}

func (h *Handler) RequestProblem(ctx context.Context, request gateway.RequestProblemRequestObject) (gateway.RequestProblemResponseObject, error) {
	panic("implement me")
}
