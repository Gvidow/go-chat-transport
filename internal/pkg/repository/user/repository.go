package user

import (
	"context"
	errorsStd "errors"
	"strconv"

	"github.com/gvidow/go-chat-transport/pkg/errors"
	lru "github.com/hashicorp/golang-lru/v2"
)

var ErrNotFound = errorsStd.New("not fount username")

type Repository interface {
	GetUsernameById(context.Context, uint64) (string, error)
	SaveUsername(ctx context.Context, id uint64, username string) error
}

type usernameStory struct {
	story *lru.Cache[uint64, string]
}

func NewUsernameStory(size int) (*usernameStory, error) {
	lru, err := lru.New[uint64, string](size)
	if err != nil {
		return nil, errors.WrapError(err, "create username repository")
	}
	return &usernameStory{
		story: lru,
	}, nil
}

func (us *usernameStory) GetUsernameById(_ context.Context, id uint64) (string, error) {
	username, ok := us.story.Get(id)
	if !ok {
		return "", errors.WrapError(ErrNotFound, "get "+strconv.Itoa(int(id)))
	}
	return username, nil
}

func (us *usernameStory) SaveUsername(_ context.Context, id uint64, username string) error {
	us.story.Add(id, username)
	return nil
}
