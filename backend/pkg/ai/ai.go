package ai

import "context"

type GPTClient interface {
	Clone() GPTClient
	Complete(ctx context.Context) (answer string, err error)
	ClearContext()
	AddPrompt(prompt string)
}

type GPTClientGenerator func() (GPTClient, error)
