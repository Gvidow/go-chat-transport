package main

import (
	"fmt"
	"log"
	"time"

	"github.com/IBM/sarama"
)

var _addrs = []string{":9092"}

const _topic = "chat"

var (
	_ = createTopic
	_ = delTopic
)

func main() {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	cl, err := sarama.NewClient(_addrs, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer cl.Close()
	t, err := cl.Topics()
	fmt.Println(t, err)
	p, err := cl.Partitions(_topic)
	fmt.Println(p, err)
	// fmt.Println(createTopic(cl, _topic, 16))
}

func createTopic(cl sarama.Client, topic string, numPartition int) error {
	res, err := cl.LeastLoadedBroker().CreateTopics(&sarama.CreateTopicsRequest{
		Timeout: time.Second,
		TopicDetails: map[string]*sarama.TopicDetail{
			"chat": {
				NumPartitions:     int32(numPartition),
				ReplicationFactor: -1,
			},
		},
	})
	fmt.Println(res)
	return err
}

func delTopic(cl sarama.Client, topic string) error {
	_, err := cl.LeastLoadedBroker().DeleteTopics(&sarama.DeleteTopicsRequest{
		Topics:  []string{topic},
		Timeout: time.Second,
	})
	return err
}
