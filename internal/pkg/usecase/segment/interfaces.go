package segment

import (
	"context"
	"time"

	"github.com/gvidow/go-chat-transport/internal/pkg/entity"
)

type SegmentStacker interface {
	SaveToQueue(context.Context, *entity.Segment) error
}

type SegmentProcessor interface {
	SegmentProcessing(context.Context, *entity.Segment) error
	StartSegmentProcessing(ctx context.Context, delta time.Duration) error
	Stop() error
}

type Builder interface {
	Build() *entity.MessageWithErrorFlag
	AddSegment(*entity.Segment)
}
