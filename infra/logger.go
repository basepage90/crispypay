package infra

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewLogger(lc fx.Lifecycle) (*zap.Logger, func(), error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() { _ = logger.Sync() }
	return logger, cleanup, nil
}
