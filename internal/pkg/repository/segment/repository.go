package segment

import (
	"context"

	"github.com/gvidow/go-chat-transport/internal/pkg/entity"
)

type Transfer interface {
	DoTransfer(context.Context, *entity.Segment) error
}

type Stacker interface {
	Put(context.Context, *entity.Segment) error
}

type Repository interface {
	Stacker
	Transfer
}

type repo struct {
	Stacker
	Transfer
}

func NewRepository(s Stacker, t Transfer) Repository {
	return repo{s, t}
}
