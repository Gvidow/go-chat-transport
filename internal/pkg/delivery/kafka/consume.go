package kafka

import (
	"encoding/json"

	"github.com/IBM/sarama"
	"github.com/gvidow/go-chat-transport/internal/pkg/entity"
	"github.com/gvidow/go-chat-transport/pkg/errors"
)

func (c *consumerHub) Consume(pc sarama.PartitionConsumer) {
	var err error

	for mes := range pc.Messages() {
		err = c.serveMessage(mes)
		if err != nil {
			c.log.Sugar().Errorf("serve message with key %s: %v", mes.Key, err)
		}
	}
}

func (c *consumerHub) serveMessage(m *sarama.ConsumerMessage) error {
	data := m.Value
	segment := &entity.Segment{}

	err := json.Unmarshal(data, segment)
	if err != nil {
		return errors.WrapFail(err, "unmarshal consumer message")
	}

	_ = segment

	return nil
}
