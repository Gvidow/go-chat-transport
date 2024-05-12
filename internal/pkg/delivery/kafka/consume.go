package kafka

import (
	"context"
	"encoding/json"

	"github.com/IBM/sarama"
	"github.com/gvidow/go-chat-transport/internal/pkg/entity"
	"github.com/gvidow/go-chat-transport/pkg/errors"
)

func (c *consumerHub) Consume(pc sarama.PartitionConsumer) {
	var err error
	log := c.cfg.Log.Sugar()

	for mes := range pc.Messages() {
		err = c.serveMessage(c.ctx, mes)
		if err != nil {
			log.Errorf("serve message with key %s: %v", mes.Key, err)
		}
	}
}

func (c *consumerHub) serveMessage(ctx context.Context, m *sarama.ConsumerMessage) error {
	data := m.Value
	segment := &entity.Segment{}

	err := json.Unmarshal(data, segment)
	if err != nil {
		return errors.WrapFail(err, "unmarshal consumer message")
	}

	return errors.WrapFail(c.usecase.SegmentProcessing(ctx, segment), "serve segment")
}
