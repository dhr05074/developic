package ai

import "context"

type GPTClient interface {
	NewContext() GPTClient
	Complete(ctx context.Context) (answer string, err error)
	ClearContext()
	AddPrompt(prompt string)
}
