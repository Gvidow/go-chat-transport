package segment

import (
	"context"

	"github.com/gvidow/go-chat-transport/internal/pkg/entity"
)

type Transfer interface {
	Transfer(context.Context, *entity.Segment) error
}

type Stacker interface {
	Put(context.Context, *entity.Segment) error
}

type Repository interface {
	Stacker
	Transfer
}
