package logger

import (
	"go.uber.org/zap"
)

func New() *zap.Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	return logger
}
