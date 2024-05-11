package kafka

import (
	"github.com/IBM/sarama"
	"github.com/gvidow/go-chat-transport/pkg/errors"
)

type manager struct {
	addrs  []string
	client sarama.Client
	config *sarama.Config
	topic  string
}

func NewKafkaManager(addrs []string, topic string, config *sarama.Config) (*manager, error) {
	m := &manager{
		addrs:  addrs,
		config: config,
		topic:  topic,
	}

	client, err := sarama.NewClient(addrs, config)
	if err != nil {
		return nil, errors.WrapError(err, "create kafka client")
	}

	m.client = client
	return m, nil
}

func (m *manager) Topic() string {
	return m.topic
}

func (m *manager) Partitions() ([]int32, error) {
	p, err := m.client.Partitions(m.topic)
	if err != nil {
		return nil, errors.WrapError(err, "manager: get partitions")
	}
	return p, nil
}

func (m *manager) NewSyncProducer() (sarama.SyncProducer, error) {
	p, err := sarama.NewSyncProducer(m.addrs, m.config)
	if err != nil {
		return nil, errors.WrapError(err, "manager: new sync producer")
	}
	return p, nil
}

func (m *manager) NewConsumer() (sarama.Consumer, error) {
	c, err := sarama.NewConsumer(m.addrs, m.config)
	if err != nil {
		return nil, errors.WrapError(err, "manager: new consumer")
	}
	return c, nil
}
