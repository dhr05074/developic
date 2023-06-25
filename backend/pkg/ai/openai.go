package ai

import (
	"code-connect/schema"
	"context"
	"errors"
	"github.com/sashabaranov/go-openai"
	"os"
)

var (
	errAPIKeyNotFound = errors.New("openai api key not found")
)

const (
	gptTemperature = 0.8
)

type OpenAI struct {
	openaiClient *openai.Client
	messages     []openai.ChatCompletionMessage
}

func (o *OpenAI) AddPrompt(prompt string) {
	o.messages = append(
		o.messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: prompt,
		},
	)
}

func (o *OpenAI) Clone() GPTClient {
	return NewOpenAI(o.openaiClient)
}

func (o *OpenAI) ClearContext() {
	o.messages = []openai.ChatCompletionMessage{}
}

func (o *OpenAI) Complete(ctx context.Context) (answer string, err error) {
	response, err := o.openaiClient.CreateChatCompletion(
		ctx, openai.ChatCompletionRequest{
			Model:       openai.GPT3Dot5Turbo0613,
			Temperature: gptTemperature,
			Messages:    o.messages,
		},
	)
	if err != nil {
		return "", err
	}

	o.messages = append(o.messages, response.Choices[0].Message)

	return response.Choices[0].Message.Content, nil
}

func NewOpenAI(openaiClient *openai.Client) *OpenAI {
	return &OpenAI{openaiClient: openaiClient}
}

func newOpenAIClientFromEnv() (*openai.Client, error) {
	apiKey, ok := os.LookupEnv(schema.ChatGPTAPIKeyEnvKey)
	if !ok {
		return nil, errAPIKeyNotFound
	}

	return newOpenAIClient(apiKey), nil
}

func newOpenAIClient(apiKey string) *openai.Client {
	return openai.NewClient(apiKey)
}

func NewDefaultOpenAIClientGenerator() (GPTClient, error) {
	return newDefaultOpenAIClient()
}

func newDefaultOpenAIClient() (*OpenAI, error) {
	openaiClient, err := newOpenAIClientFromEnv()
	if err != nil {
		return nil, err
	}

	return NewOpenAI(openaiClient), nil
}
