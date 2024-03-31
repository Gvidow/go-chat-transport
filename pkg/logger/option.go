package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Option interface {
	apply(*zap.Config)
}

type option func(*zap.Config)

func (opt option) apply(cfg *zap.Config) {
	opt(cfg)
}

func WithTimeLayout(layout string) Option {
	return option(func(cfg *zap.Config) {
		cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(layout)
	})
}

func WithRFC3339Time() Option {
	return option(func(cfg *zap.Config) {
		cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	})
}
