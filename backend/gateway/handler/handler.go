package handler

import (
	"code-connect/gateway"
	"code-connect/problem"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
)

type ServerHandler struct {
	openaiClient   *openai.Client
	problemHandler *problem.Handler
}

func (s ServerHandler) CreateProblem(ctx echo.Context) error {
	var request problem.CreateProblemRequest
	if err := ctx.Bind(&request); err != nil {
		return err
	}

	result, err := s.problemHandler.Create(ctx.Request().Context(), request)
	if err != nil {
		return err
	}

	return ctx.JSON(202, problem.CreateProblemResponse{ID: result.ID})
}

func (s ServerHandler) GetProblem(ctx echo.Context, problemId gateway.ProblemID) error {
	result, err := s.problemHandler.Get(ctx.Request().Context(), problem.GetProblemRequest{ID: problemId})
	if err != nil {
		if errors.Is(err, problem.NotFound) {
			return ctx.JSON(404, echo.Map{
				"message": fmt.Sprintf("problem is not found: %s", problemId),
			})
		}

		return err
	}

	response := gateway.GetProblemResponse{
		Content:  result.Statement,
		Id:       problemId,
		Language: result.Language,
	}

	return ctx.JSON(200, response)
}
