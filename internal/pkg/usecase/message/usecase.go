package message

import (
	"context"

	"github.com/gvidow/go-chat-transport/internal/pkg/entity"
	"github.com/gvidow/go-chat-transport/internal/pkg/repository/segment"
	"github.com/gvidow/go-chat-transport/pkg/errors"
	"github.com/gvidow/go-chat-transport/pkg/math"
)

type Sender interface {
	Send(context.Context, *entity.Message) error
}

type usecaseMessage struct {
	repoSegment segment.Repository
	segmentSize uint32
}

func NewUsecaseMessage(repo segment.Repository, opts ...option) Sender {
	usecase := &usecaseMessage{
		repoSegment: repo,
		segmentSize: 0,
	}

	for _, opt := range opts {
		opt(usecase)
	}

	return usecase
}

func (u *usecaseMessage) Send(ctx context.Context, mes *entity.Message) error {
	segments := u.SplitIntoSegments(mes)
	for ind := range segments {
		err := u.repoSegment.Transfer(ctx, &segments[ind])
		if err != nil {
			return errors.WrapError("send message by segments", err)
		}
	}
	return nil
}

func (u *usecaseMessage) SplitIntoSegments(mes *entity.Message) []entity.Segment {
	if u.segmentSize == 0 {
		return []entity.Segment{{
			Time: mes.Time,
			Size: 1,
			Num:  0,
			Data: mes.Content,
		}}
	}

	countSegment := (len(mes.Content) + int(u.segmentSize) - 1) / int(u.segmentSize)
	res := make([]entity.Segment, 0, countSegment)

	for i := 0; i < countSegment; i++ {
		res = append(res, entity.Segment{
			Time: mes.Time,
			Size: countSegment,
			Num:  i,
			Data: mes.Content[i*int(u.segmentSize) : math.MinInt(len(mes.Content), i*int(u.segmentSize)+int(u.segmentSize))],
		})
	}

	return res
}
