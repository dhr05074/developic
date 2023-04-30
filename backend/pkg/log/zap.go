package log

import "go.uber.org/zap"

func NewZap() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	return logger.Sugar()
}
