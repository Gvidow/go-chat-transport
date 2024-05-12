package kafka

import (
	"context"
	"sync"

	"github.com/IBM/sarama"
	"github.com/gvidow/go-chat-transport/internal/pkg/usecase/segment"
	"github.com/gvidow/go-chat-transport/pkg/errors"
)

type manager interface {
	Topic() string
	Partitions() ([]int32, error)
	NewConsumer() (sarama.Consumer, error)
}

type consumerHub struct {
	ctx context.Context
	cfg *Config
	wg  *sync.WaitGroup

	consumer sarama.Consumer

	usecase segment.SegmentProcessor
}

func NewConsumerHub(ctx context.Context, cfg *Config, u segment.SegmentProcessor) (*consumerHub, error) {
	consumer, err := cfg.Manager.NewConsumer()
	if err != nil {
		return nil, errors.WrapError(err, "new consumer for consumer hub")
	}
	return &consumerHub{
		ctx:      ctx,
		cfg:      cfg,
		usecase:  u,
		wg:       &sync.WaitGroup{},
		consumer: consumer,
	}, nil
}

func (c *consumerHub) Serve() error {
	errChan := make(chan error)

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		topic := c.cfg.Manager.Topic()
		partitions, err := c.cfg.Manager.Partitions()
		errChan <- errors.WrapError(err, "start serve partitions")
		for _, partition := range partitions {
			c.wg.Add(1)
			go func(partition int32) {
				defer c.wg.Done()
				pc, err := c.consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
				if err != nil {
					c.cfg.Log.Error(err.Error())
					return
				}
				c.Consume(pc)
			}(partition)
		}

	}()

	err := <-errChan
	if err != nil {
		return errors.WrapError(err, "serve consumer hub")
	}

	return errors.WrapError(
		c.usecase.StartSegmentProcessing(c.ctx, c.cfg.TimeoutSenderMessage),
		"start segment processing",
	)
}

func (c *consumerHub) Shutdown() {
	c.consumer.Close()
	c.usecase.Stop()
	c.wg.Wait()
}
