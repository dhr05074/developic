package handler

import (
	"context"
	"github.com/sashabaranov/go-openai"
)

type OpenAI struct {
	openaiClient *openai.Client
	messages     []openai.ChatCompletionMessage
}

func NewOpenAI(openaiClient *openai.Client) *OpenAI {
	return &OpenAI{openaiClient: openaiClient}
}

func (o *OpenAI) AddPrompt(prompt string) {
	o.messages = append(o.messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: prompt,
	})
}

func (o *OpenAI) ClearContext() {
	o.messages = []openai.ChatCompletionMessage{}
}

func (o *OpenAI) Complete(ctx context.Context, temperature ...float32) (answer string, err error) {
	var temp float32 = 0.8
	if len(temperature) > 0 {
		temp = temperature[0]
	}

	response, err := o.openaiClient.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:       openai.GPT3Dot5Turbo,
		Temperature: temp,
		Messages:    o.messages,
	})
	if err != nil {
		return "", err
	}

	o.messages = append(o.messages, response.Choices[0].Message)

	return response.Choices[0].Message.Content, nil
}
