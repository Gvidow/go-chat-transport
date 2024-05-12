package message

import (
	"context"

	"github.com/gvidow/go-chat-transport/internal/pkg/entity"
	"github.com/gvidow/go-chat-transport/internal/pkg/repository/segment"
	"github.com/gvidow/go-chat-transport/internal/pkg/repository/user"
	"github.com/gvidow/go-chat-transport/pkg/errors"
)

type Sender interface {
	Send(context.Context, *entity.Message) error
}

type usecaseMessage struct {
	repoSegment  segment.Repository
	segmentation func(*entity.Message) []entity.Segment
	userRepo     user.Repository
}

func NewUsecaseMessage(repo segment.Repository, userRepo user.Repository, opts ...option) Sender {
	usecase := &usecaseMessage{
		repoSegment: repo,
		userRepo:    userRepo,
	}

	for _, opt := range opts {
		opt(usecase)
	}

	return usecase
}

func (u *usecaseMessage) Send(ctx context.Context, mes *entity.Message) error {
	segments := u.SplitIntoSegments(mes)

	err := u.userRepo.SaveUsername(ctx, mes.Time, mes.Username)
	if err != nil {
		return errors.WrapFail(err, "save username")
	}

	for ind := range segments {
		err = u.repoSegment.DoTransfer(ctx, &segments[ind])
		if err != nil {
			return errors.WrapFail(err, "send message by segments")
		}
	}
	return nil
}

func (u *usecaseMessage) SplitIntoSegments(mes *entity.Message) []entity.Segment {
	if u.segmentation != nil {
		return u.segmentation(mes)
	}

	return []entity.Segment{{
		Time: mes.Time,
		Size: 1,
		Num:  0,
		Data: mes.Content,
	}}
}
