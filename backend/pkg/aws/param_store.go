package aws

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

type SSMClient struct {
	client *ssm.Client
}

func NewSSMClient(client *ssm.Client) *SSMClient {
	return &SSMClient{client: client}
}

func (c *SSMClient) GetParameter(ctx context.Context, name string) (string, error) {
	result, err := c.client.GetParameter(ctx, &ssm.GetParameterInput{
		Name: &name,
	})
	if err != nil {
		return "", err
	}

	if result.Parameter.Value == nil {
		return "", errors.New("value is nil")
	}

	return *result.Parameter.Value, nil
}
