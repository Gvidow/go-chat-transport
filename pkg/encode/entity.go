package encode

import (
	"encoding/json"

	"github.com/gvidow/go-chat-transport/pkg/errors"
)

type entityEncoder struct {
	data []byte
}

func NewEntityEncoder(entity any) (entityEncoder, error) {
	data, err := json.Marshal(entity)
	if err != nil {
		return entityEncoder{}, errors.WrapFail(err, "marshal entity")
	}
	return entityEncoder{
		data: data,
	}, nil
}

func (e entityEncoder) Encode() ([]byte, error) {
	return e.data, nil
}

func (e entityEncoder) Length() int {
	return len(e.data)
}
