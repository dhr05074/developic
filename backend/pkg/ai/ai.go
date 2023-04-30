package ai

import "context"

type GPTClient interface {
	NewContext() GPTClient
	CompleteWithContext(ctx context.Context, prompts []string) (answer string, err error)
	ClearContext()
}
