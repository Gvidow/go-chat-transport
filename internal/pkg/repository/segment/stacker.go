package segment

import (
	"context"

	"github.com/IBM/sarama"

	"github.com/gvidow/go-chat-transport/internal/pkg/entity"
	"github.com/gvidow/go-chat-transport/pkg/encode"
	"github.com/gvidow/go-chat-transport/pkg/errors"
)

type producerGetter interface {
	Topic() string
	NewSyncProducer() (sarama.SyncProducer, error)
}

type segmentStacker struct {
	producer sarama.SyncProducer
	topic    string
}

func NewSegmentStacker(pg producerGetter) (*segmentStacker, error) {
	syncProducer, err := pg.NewSyncProducer()
	if err != nil {
		return nil, errors.WrapError(err, "get sync producer")
	}
	return &segmentStacker{
		producer: syncProducer,
		topic:    pg.Topic(),
	}, nil
}

func (s *segmentStacker) Put(ctx context.Context, segment *entity.Segment) error {
	data, err := encode.NewEntityEncoder(segment)
	if err != nil {
		return errors.WrapFail(err, "make data encoder")
	}
	mes := &sarama.ProducerMessage{
		Topic: s.topic,
		Key:   encode.Uint64Encoder(segment.Time),
		Value: data,
	}

	_, _, err = s.producer.SendMessage(mes)
	return errors.WrapFail(err, "send producer message")
}
