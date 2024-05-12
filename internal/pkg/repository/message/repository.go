package message

import (
	"context"
	"net/http"

	"github.com/gvidow/go-chat-transport/internal/pkg/entity"
	"github.com/gvidow/go-chat-transport/internal/pkg/repository"
	"github.com/gvidow/go-chat-transport/pkg/errors"
)

type Repository interface {
	Send(context.Context, *entity.MessageWithErrorFlag) error
}

type senderMsgWithFlag struct {
	Client *http.Client
	addr   string
}

func NewSenderMsg(addr string) *senderMsgWithFlag {
	return &senderMsgWithFlag{
		addr:   addr,
		Client: &http.Client{},
	}
}

func (s *senderMsgWithFlag) Send(ctx context.Context, msg *entity.MessageWithErrorFlag) error {
	body, err := repository.MakeRequestBody(msg)
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

	return errors.MakeHTTPError(res, "send message with error flag")
}
