package ai

import "context"

type GPTClient interface {
	NewClientWithEmptyContext() GPTClient
	CompleteWithContext(ctx context.Context, prompts []string) (answer string, err error)
	ClearContext()
}
