package repository

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/gvidow/go-chat-transport/pkg/errors"
)

func MakeRequestBody(segment any) (io.Reader, error) {
	data, err := json.Marshal(segment)
	if err != nil {
		return nil, errors.WrapFail(err, "marshal segment")
	}

	return bytes.NewReader(data), nil
}
