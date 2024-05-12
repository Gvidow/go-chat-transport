package segment

import (
	"context"
	"net/http"

	"github.com/gvidow/go-chat-transport/internal/pkg/entity"
	"github.com/gvidow/go-chat-transport/internal/pkg/repository"
	"github.com/gvidow/go-chat-transport/pkg/errors"
)

type segmentTransfer struct {
	Client *http.Client
	addr   string
}

func NewSegmentTransfer(address string) *segmentTransfer {
	return &segmentTransfer{
		Client: &http.Client{},
		addr:   address,
	}
}

func (s *segmentTransfer) DoTransfer(ctx context.Context, segment *entity.Segment) error {
	body, err := repository.MakeRequestBody(segment)
	if err != nil {
		return errors.WrapFail(err, "make body for request")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.addr, body)
	if err != nil {
		return errors.WrapFail(err, "make request")
	}

	res, err := s.Client.Do(req)
	if err != nil {
		return errors.WrapError(err, "do request")
	}

	return errors.MakeHTTPError(res, "transfer segment")
}
