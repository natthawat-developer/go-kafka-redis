package kafka

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer struct {
	client *kafka.Consumer
}

func NewConsumer(bootstrapServers, groupID, autoOffsetReset string) (*Consumer, error) {
	client, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
		"group.id":          groupID,
		"auto.offset.reset": autoOffsetReset,
	})
	if err != nil {
		return nil, err
	}

	return &Consumer{client: client}, nil
}

func (c *Consumer) SubscribeTopic(topic string) {
	err := c.client.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %s", err)
	}
}

func (c *Consumer) ReadMessage() (*kafka.Message, error) {
	return c.client.ReadMessage(-1)
}

func (c *Consumer) Close() {
	c.client.Close()
}
