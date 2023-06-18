package problem

import (
	"code-connect/ent/enttest"
	"code-connect/gateway"
	"code-connect/mocks"
	"context"
	nanoid "github.com/matoous/go-nanoid/v2"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type ProblemHandlerTestSuite struct {
	suite.Suite

	hdl *Handler
	kv  *mocks.KV
	gpt *mocks.GPTClient
}

func (s *ProblemHandlerTestSuite) SetupTest() {
	s.kv = &mocks.KV{}
	s.gpt = &mocks.GPTClient{}
	cli := enttest.Open(s.T(), "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	s.hdl = NewHandler(s.kv, s.gpt, cli, nil)
}

func (s *ProblemHandlerTestSuite) TestWhenConcurrentlyAccessTheSameProblem_ThenFinallyProblemCreated() {
	ctx := context.Background()

	id := nanoid.Must(8)
	s.Require().NoError(s.hdl.createProblem(ctx, id, gateway.RequestProblemRequestObject{
		Body: &gateway.RequestProblemJSONRequestBody{
			EloScore: nil,
			Language: gateway.Go,
		},
	}))

	resp, err := s.hdl.GetProblem(ctx, gateway.GetProblemRequestObject{Id: id})
	s.Require().NoError(err)
	s.Require().NoError(s.hdl.saveProblem(ctx, id, GPTOutput{
		Title: "hi",
		Code:  "21",
	}))
	resp, err = s.hdl.GetProblem(ctx, gateway.GetProblemRequestObject{Id: id})
	s.Require().NoError(err)

	_, ok := resp.(gateway.GetProblem409JSONResponse)
	s.Require().True(ok)

	s.Eventually(func() bool {
		resp, err := s.hdl.GetProblem(ctx, gateway.GetProblemRequestObject{Id: id})
		s.Require().NoError(err)

		_, ok := resp.(gateway.GetProblem200JSONResponse)
		return ok
	}, 500*time.Millisecond, 10*time.Millisecond)
}

func TestHandler_ProblemHandler(t *testing.T) {
	suite.Run(t, new(ProblemHandlerTestSuite))
}
