package kafka

import (
	"time"

	log "github.com/gvidow/go-chat-transport/pkg/logger"
)

type Config struct {
	Log                  *log.Logger
	TimeoutSenderMessage time.Duration
	Manager              manager
}
