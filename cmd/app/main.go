package main

import (
	"context"
	"log"

	"github.com/gvidow/go-chat-transport/internal/app"
	"github.com/gvidow/go-chat-transport/pkg/logger"
)

func main() {
	baseCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger, err := logger.NewLogger(logger.WithRFC3339Time())
	if err != nil {
		log.Printf("new logger: %v\n", err)
		return
	}
	defer logger.Sync()

	if err := app.Main(baseCtx, logger); err != nil {
		logger.Sugar().Errorf("main executed with error: %v", err)
	}
}
