package http

import (
	"context"
	"time"

	"github.com/gvidow/go-chat-transport/internal/pkg/usecase/message"
	"github.com/gvidow/go-chat-transport/internal/pkg/usecase/segment"
	"github.com/gvidow/go-chat-transport/pkg/logger"
)

type Handler struct {
	log            *logger.Logger
	segmentUsecase segment.SegmentStacker
	messageUsecase message.Sender
	baseCtx        context.Context
}

func NewHandler(log *logger.Logger, s segment.SegmentStacker, m message.Sender) *Handler {
	return &Handler{
		log:            log,
		segmentUsecase: s,
		messageUsecase: m,
		baseCtx:        context.Background(),
	}
}

var (
	_timeoutSendMessage = 3 * time.Minute
	_timeoutPutSegment  = time.Minute
)
