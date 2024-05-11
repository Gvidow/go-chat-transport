package segment

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gvidow/go-chat-transport/internal/pkg/entity"
	"github.com/gvidow/go-chat-transport/pkg/errors"
)

type segmentTransfer struct {
	client *http.Client
	addr   string
}

func NewSegmentTransfer(address string) *segmentTransfer {
	return &segmentTransfer{
		client: http.DefaultClient,
		addr:   address,
	}
}

func (s *segmentTransfer) Transfer(ctx context.Context, segment *entity.Segment) error {
	body, err := makeRequestBody(segment)
	if err != nil {
		return errors.WrapFail(err, "make body for request")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.addr, body)
	if err != nil {
		return errors.WrapFail(err, "make request")
	}

	res, err := s.client.Do(req)
	if err != nil {
		return errors.WrapError(err, "do request")
	}

	return errors.MakeHTTPError(res, "transfer segment")
}

func makeRequestBody(segment *entity.Segment) (io.Reader, error) {
	data, err := json.Marshal(segment)
	if err != nil {
		return nil, errors.WrapFail(err, "marshal segment")
	}

	return bytes.NewReader(data), nil
}
