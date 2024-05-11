package kafka

import (
	"context"

	"github.com/gvidow/go-chat-transport/pkg/logger"
)

type consumerHub struct {
	log *logger.Logger
	ctx context.Context
}

func NewConsumerHub() *consumerHub {
	return &consumerHub{}
}
