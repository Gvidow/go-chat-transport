package segment

import (
	"context"
	errorsStd "errors"
	"fmt"
	"sync"
	"time"

	"github.com/gvidow/go-chat-transport/internal/pkg/entity"
	"github.com/gvidow/go-chat-transport/internal/pkg/repository/message"
	"github.com/gvidow/go-chat-transport/internal/pkg/repository/segment"
	"github.com/gvidow/go-chat-transport/internal/pkg/repository/user"
	"github.com/gvidow/go-chat-transport/internal/pkg/usecase/message/builder"
	"github.com/gvidow/go-chat-transport/pkg/errors"
	log "github.com/gvidow/go-chat-transport/pkg/logger"
)

var ErrContextProcessingCancel = errorsStd.New("cancel context")
var ErrProcessorNotRunning = errorsStd.New("processor is not running")

type usecaseSegment struct {
	log      *log.Logger
	repo     segment.Repository
	userRepo user.Repository
	msgRepo  message.Repository
	ctx      context.Context
	cancel   func()
	ticker   *time.Ticker
	story    map[uint64]Builder
	mu       sync.Mutex
}

var _ interface {
	SegmentProcessor
	SegmentStacker
} = (*usecaseSegment)(nil)

func NewSegmentUsecase(log *log.Logger, repo segment.Repository, userRepo user.Repository, msgRepo message.Repository) *usecaseSegment {
	return &usecaseSegment{
		log:      log,
		repo:     repo,
		userRepo: userRepo,
		msgRepo:  msgRepo,
		mu:       sync.Mutex{},
	}
}

func (u *usecaseSegment) SaveToQueue(ctx context.Context, segment *entity.Segment) error {
	return errors.WrapError(u.repo.Put(ctx, segment), "save to queue")
}

func (u *usecaseSegment) SegmentProcessing(ctx context.Context, segment *entity.Segment) error {
	username, err := u.userRepo.GetUsernameById(ctx, segment.Time)
	if err != nil {
		return errors.WrapFail(err, "get username for segment")
	}

	u.mu.Lock()
	defer u.mu.Unlock()
	if u.story != nil {
		if b, ok := u.story[segment.Time]; ok {
			b.AddSegment(segment)
		} else {
			b = builder.NewBuilder(segment.Time, username, segment.Size)
			b.AddSegment(segment)
			u.story[segment.Time] = b
		}
		return nil
	}
	return ErrProcessorNotRunning
}

func (u *usecaseSegment) StartSegmentProcessing(ctx context.Context, delta time.Duration) error {
	u.ctx, u.cancel = context.WithCancel(ctx)
	u.ticker = time.NewTicker(delta)
	u.story = make(map[uint64]Builder)

	var err error

	for {
		select {
		case <-u.ticker.C:
			u.mu.Lock()
			err = errors.WrapError(u.processBatchSegment(ctx, u.story), "segment processing: batch process")
			u.story = make(map[uint64]Builder)
			u.mu.Unlock()

			if err != nil {
				u.log.Error(err.Error())
			}
		case <-u.ctx.Done():
			return ErrContextProcessingCancel
		}
	}
}

func (u *usecaseSegment) Stop() error {
	if u.cancel != nil {
		u.ticker.Stop()
		u.cancel()
		u.cancel = nil
		u.story = nil
	}
	return nil
}

func (u *usecaseSegment) processBatchSegment(ctx context.Context, batch map[uint64]Builder) error {
	var err error
	for _, messageBuilder := range batch {
		msg := messageBuilder.Build()
		if err = u.msgRepo.Send(ctx, msg); err != nil {
			return errors.WrapError(err, fmt.Sprintf("send message with error flag: %v", msg))
		}

		u.log.Sugar().Infof("success send message with flag: %v", msg)
	}
	return nil
}
