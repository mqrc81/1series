package api

import (
	"go.uber.org/zap"
)

func NewLogger() (*zap.SugaredLogger, error) {
	logger, err := zap.NewProduction(zap.AddStacktrace(zap.ErrorLevel))
	if err != nil {
		return nil, err
	}

	return logger.Sugar(), nil
}
