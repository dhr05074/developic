package handler

import (
	"code-connect/gateway"
	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
)

type ServerHandler struct {
	openaiClient *openai.Client
}

func (s ServerHandler) CreateProblem(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s ServerHandler) GetProblem(ctx echo.Context, problemId gateway.ProblemID) error {
	//TODO implement me
	panic("implement me")
}
