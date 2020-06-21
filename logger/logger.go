package logger

import (
	"go.uber.org/zap"
)

// New create a new logger
func New() *zap.Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	return logger
}
