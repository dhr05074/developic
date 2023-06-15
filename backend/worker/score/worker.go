package score

import (
	"code-connect/ent"
	"code-connect/ent/problem"
	"code-connect/pkg/ai"
	"code-connect/pkg/log"
	"code-connect/pkg/store"
	"code-connect/schema/message"
	"context"
	"encoding/base64"
	"encoding/json"
	"strings"
)

const (
	scorePromptKey = "/prompts/score"
	replaceKey     = "{CODE}"
)

type ScoreWorker struct {
	entClient   *ent.Client
	paramClient store.KV
	problemCh   chan message.ProblemMessage
	submitCh    chan message.ProblemMessage
	gptClient   ai.GPTClient
	propt       string
}

type ScoreResult struct {
	Efficiency  int `json:"efficiency"`
	Robustness  int `json:"robustness"`
	Readability int `json:"readability"`
}

type NewScoreWorkerParams struct {
	ParamClient store.KV
	EntClient   *ent.Client
	GptClient   ai.GPTClient
	ProblemCh   chan message.ProblemMessage
	SubmitCh    chan message.ProblemMessage
}

func NewScoreWorker(params NewScoreWorkerParams) *ScoreWorker {
	return &ScoreWorker{
		paramClient: params.ParamClient,
		entClient:   params.EntClient,
		problemCh:   params.ProblemCh,
		submitCh:    params.SubmitCh,
		gptClient:   params.GptClient,
	}
}

func (s *ScoreWorker) initialize(ctx context.Context) error {
	prompt, err := s.paramClient.Get(ctx, scorePromptKey)
	if err != nil {
		return err
	}

	s.propt = prompt

	return nil
}

func (s *ScoreWorker) handlingCode(ctx context.Context, code string) (*ScoreResult, error) {
	prompt := strings.Replace(s.propt, replaceKey, code, 1)

	s.gptClient.AddPrompt(prompt)

	result, err := s.gptClient.Complete(ctx)
	if err != nil {
		log.NewZap().Error(err.Error())
		return nil, err
	}

	var scoreResult ScoreResult
	err = json.Unmarshal([]byte(result), &scoreResult)
	if err != nil {
		log.NewZap().Error(err.Error())
		return nil, err
	}

	return &scoreResult, nil
}

func (s *ScoreWorker) Run(ctx context.Context) (err error) {
	err = s.initialize(ctx)
	if err != nil {
		return err
	}

	for {
		select {
		case msg := <-s.problemCh:
			p, err := s.entClient.Problem.Query().Where(problem.UUID(msg.ID)).Only(ctx)
			if err != nil {
				log.NewZap().Error(err.Error())
				continue
			}

			code, _ := base64.StdEncoding.DecodeString(p.Code)
			result, err := s.handlingCode(ctx, string(code))
			if err != nil {
				continue
			}

			if err := p.Update().
				SetRobustness(result.Robustness).
				SetEfficiency(result.Efficiency).
				SetReadability(result.Readability).
				Exec(ctx); err != nil {
				log.NewZap().Error(err.Error())
				continue
			}
			log.NewZap().Info("Scored Problem Successfully")

		case <-s.submitCh:
			//r, err := s.entClient.Record.Query().Where(record.UUID(msg.ID)).Only(ctx)
			//if err != nil {
			//	log.NewZap().Error(err.Error())
			//	continue
			//}
			//
			//code, _ := base64.StdEncoding.DecodeString(r.Code)
			//s.handlingCode(ctx, string(code))

		case <-ctx.Done():
			return nil
		}
	}

}
