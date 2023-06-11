package score

import "code-connect/pkg/aws"

type ScoreWorker struct {
	ssmClient *aws.SSMClient
}
