package aws

import (
	"code-connect/pkg/log"
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"go.uber.org/zap"
)

var (
	errValueIsNil = errors.New("value is nil")
)

type SSMClient struct {
	client *ssm.Client
	logger *zap.SugaredLogger
}

func (c *SSMClient) Set(_ context.Context, _, _ string) error {
	panic("implement me")
}

func (c *SSMClient) Get(ctx context.Context, name string) (string, error) {
	result, err := c.client.GetParameter(
		ctx, &ssm.GetParameterInput{
			Name: &name,
		},
	)
	if err != nil {
		c.logger.Errorw("AWS SSM으로부터 파라미터를 가져오는데 실패했습니다.", "name", name, "error", err)
		return "", err
	}

	paramValue := result.Parameter.Value
	if paramValue == nil {
		return "", errValueIsNil
	}

	return *paramValue, nil
}

func NewDefaultSSMKeyValueStore(ctx context.Context) (*SSMClient, error) {
	cfg, err := LoadConfig(ctx)
	if err != nil {
		return nil, err
	}

	return newSSMClient(ssm.NewFromConfig(cfg)), nil
}

func newSSMClient(client *ssm.Client) *SSMClient {
	logger := log.NewZap().With("client", "aws.ssm")
	return &SSMClient{client: client, logger: logger}
}
