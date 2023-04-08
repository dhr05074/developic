package log

import "go.uber.org/zap"

func NewZapSugaredLogger() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	return logger.Sugar()
}
