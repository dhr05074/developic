package score

import (
	"code-connect/ent"
	"code-connect/ent/problem"
	"code-connect/ent/record"
	"code-connect/pkg/ai"
	"code-connect/pkg/aws"
	"code-connect/pkg/log"
	"code-connect/schema/message"
	"context"
	"encoding/base64"
)

type ScoreWorker struct {
	ssmClient *aws.SSMClient
	entClient *ent.Client
	problemCh chan message.ProblemMessage
	submitCh  chan message.ProblemMessage
	gptClient *ai.OpenAI
	checkList map[string]string
}

type NewScoreWorkerParams struct {
	SSMClient *aws.SSMClient
	EntClient *ent.Client
	GptClient *ai.OpenAI
	ProblemCh chan message.ProblemMessage
	SubmitCh  chan message.ProblemMessage
}

func NewScoreWorker(params NewScoreWorkerParams) *ScoreWorker {
	return &ScoreWorker{
		ssmClient: params.SSMClient,
		entClient: params.EntClient,
		problemCh: params.ProblemCh,
		submitCh:  params.SubmitCh,
		gptClient: params.GptClient,
		checkList: make(map[string]string),
	}
}

func (s *ScoreWorker) initialize(ctx context.Context) error {
	readability, err := s.ssmClient.GetParameter(ctx, "/prompts/problem/evaluating/readability")
	if err != nil {
		return err
	}

	s.checkList["readability"] = readability

	efficiency, err := s.ssmClient.GetParameter(ctx, "/prompts/problem/evaluating/efficiency")
	if err != nil {
		return err
	}

	s.checkList["efficiency"] = efficiency

	modularity, err := s.ssmClient.GetParameter(ctx, "/prompts/problem/evaluating/modularity")
	if err != nil {
		return err
	}

	s.checkList["modularity"] = modularity

	testability, err := s.ssmClient.GetParameter(ctx, "/prompts/problem/evaluating/testability")
	if err != nil {
		return err
	}

	s.checkList["testability"] = testability

	maintainability, err := s.ssmClient.GetParameter(ctx, "/prompts/problem/evaluating/maintainability")
	if err != nil {
		return err
	}

	s.checkList["maintainability"] = maintainability

	return nil
}

func (s *ScoreWorker) handlingCode(ctx context.Context, code string) {
	basePrompt := "Given a piece of code and a checklist, please generate a list of all the checklist items that the code meets.\n\nCode : " + code +
		"\n\nCheckList :"

	for _, v := range s.checkList {
		prompt := basePrompt + v
		s.gptClient.AddPrompt(prompt)

		answer, err := s.gptClient.Complete(ctx)
		if err != nil {
			log.NewZap().Error(err.Error())
			break
		}

		log.NewZap().Info(answer)
	}
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
			s.handlingCode(ctx, string(code))

			//p.Update().SetModularity().SetMaintainablity().SetReadability().SetTestability().SetEfficiency().Save(ctx)

		case msg := <-s.submitCh:
			r, err := s.entClient.Record.Query().Where(record.UUID(msg.ID)).Only(ctx)
			if err != nil {
				log.NewZap().Error(err.Error())
				continue
			}

			code, _ := base64.StdEncoding.DecodeString(r.Code)
			s.handlingCode(ctx, string(code))

		case <-ctx.Done():
			return nil
		}
	}

}
