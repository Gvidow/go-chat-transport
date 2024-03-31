package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	zap.Logger
}

func NewLogger(opts ...Option) (*Logger, error) {
	cfg := zap.NewProductionConfig()
	for _, opt := range opts {
		opt.apply(&cfg)
	}

	log, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return &Logger{*log}, nil
}
