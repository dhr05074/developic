package ai

import "context"

type GPTClient interface {
	CompleteWithContext(ctx context.Context, prompts []string) (answer string, err error)
	ClearContext() error
}
