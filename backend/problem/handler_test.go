package problem_test

import (
	"code-connect/ent"
	"code-connect/ent/enttest"
	"code-connect/gateway"
	"code-connect/mocks"
	"code-connect/pkg/ai"
	"code-connect/problem"
	"code-connect/schema/message"
	"context"
	"github.com/stretchr/testify/suite"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func gptClientGenerator(mockClient *mocks.GPTClient) ai.GPTClientGenerator {
	return func() (ai.GPTClient, error) {
		return mockClient, nil
	}
}

type RequestProblemTestSuite struct {
	suite.Suite

	hdl           *problem.Handler
	entClient     *ent.Client
	mockGPTClient *mocks.GPTClient
}

func (s *RequestProblemTestSuite) SetupTest() {
	s.entClient = enttest.Open(s.T(), "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")

	reqCh := make(chan message.ProblemMessage)
	s.hdl = problem.NewHandler(nil, nil, s.entClient, gptClientGenerator(s.mockGPTClient), reqCh)
}

func (s *RequestProblemTestSuite) TestCreateProblemDBFailed_Return500Error() {
	s.Require().NoError(s.entClient.Close())

	req := gateway.RequestProblemRequestObject{
		Body: &gateway.RequestProblemJSONRequestBody{},
	}

	respIface, err := s.hdl.RequestProblem(context.Background(), req)
	s.Require().NoError(err)

	resp, ok := respIface.(gateway.RequestProblemdefaultJSONResponse)
	s.Require().True(ok)
	s.Require().Equal(500, resp.StatusCode)
}

func (s *RequestProblemTestSuite) TestCreateProblemELOScoreNotNilAndValid_Return200OK() {
	var score int32 = 1500
	req := gateway.RequestProblemRequestObject{
		Body: &gateway.RequestProblemJSONRequestBody{
			EloScore: &score,
			Language: gateway.Go,
		},
	}

	respIface, err := s.hdl.RequestProblem(context.Background(), req)
	s.Require().NoError(err)

	resp, ok := respIface.(gateway.RequestProblem202JSONResponse)
	s.Require().True(ok)
	s.Require().NotEmpty(resp.ProblemId)
}

func TestHandler_RequestProblem(t *testing.T) {
	suite.Run(t, new(RequestProblemTestSuite))
}
