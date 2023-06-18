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

func NewHandler(entClient *ent.Client) *Handler {
	return &Handler{entClient: entClient}
}

func (h *Handler) GetMe(ctx context.Context, req gateway.GetMeRequestObject) (gateway.GetMeResponseObject, error) {
	userID, ok := store.UsernameFromContext(ctx)
	if !ok {
		return gateway.GetMedefaultJSONResponse{
			Body:       gateway.Error{},
			StatusCode: 0,
		}, nil
	}

	count, err := h.entClient.User.Query().Where(user2.UUID(userID)).Count(ctx)
	if err != nil {
		return nil, err
	}

	var (
		user *ent.User
	)
	// 유저 신규 등록
	if count == 0 {
		user, err = h.entClient.User.Create().SetUUID(userID).SetNickname(userID).Save(ctx)
		if err != nil {
			return nil, err
		}
	} else {
		user, err = h.entClient.User.Query().Where(user2.UUID(userID)).Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return gateway.GetMe404JSONResponse{
					Code:    userNotFoundErrorCode,
					Message: userNotFoundErrorMessage,
				}, err
			}
			return nil, err
		}
	}

	return gateway.GetMe200JSONResponse{
		EloScore: gateway.ELOScore(user.EloScore),
		Nickname: user.UUID,
	}, nil
}
