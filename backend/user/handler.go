package user

import (
	"code-connect/ent"
	user2 "code-connect/ent/user"
	"code-connect/gateway"
	"code-connect/pkg/store"
	"context"
)

const (
	userNotFoundErrorCode    = "UserNotFound"
	userNotFoundErrorMessage = "유저가 존재하지 않습니다."
)

type Handler struct {
	entClient *ent.Client
}

func (h *Handler) GetMe(ctx context.Context, req gateway.GetMeRequestObject) (gateway.GetMeResponseObject, error) {
	userID, ok := store.UsernameFromContext(ctx)
	if !ok {
		return gateway.GetMedefaultJSONResponse{
			Body:       gateway.Error{},
			StatusCode: 0,
		}, nil
	}

	user, err := h.entClient.User.Query().Where(user2.UUID(userID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return gateway.GetMe404JSONResponse{
				Code:    userNotFoundErrorCode,
				Message: userNotFoundErrorMessage,
			}, err
		}
		return nil, err
	}

	return gateway.GetMe200JSONResponse{
		EloScore: gateway.ELOScore(user.EloScore),
		Nickname: user.UUID,
	}, nil
}
