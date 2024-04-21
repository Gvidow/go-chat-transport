package segment

import (
	"context"

	"github.com/gvidow/go-chat-transport/internal/pkg/entity"
	"github.com/gvidow/go-chat-transport/internal/pkg/repository/segment"
	"github.com/gvidow/go-chat-transport/pkg/errors"
)

type SegmentStacker interface {
	SaveToQueue(context.Context, *entity.Segment) error
}

type usecaseSegment struct {
	repo segment.Repository
}

func NewSegmentUsecase(repo segment.Repository) SegmentStacker {
	return usecaseSegment{repo}
}

func (u usecaseSegment) SaveToQueue(ctx context.Context, segment *entity.Segment) error {
	return errors.WrapError("save to queue", u.repo.Put(ctx, segment))
}
