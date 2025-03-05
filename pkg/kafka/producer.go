package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer struct {
	client *kafka.Producer
}

func NewProducer(bootstrapServers string) (*Producer, error) {
	client, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": bootstrapServers})
	if err != nil {
		return nil, err
	}
	return &Producer{client: client}, nil
}

func (p *Producer) ProduceMessage(topic, message string) error {
	return p.client.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)
}

func (p *Producer) Flush() {
	p.client.Flush(15 * 1000)
}

func (p *Producer) Close() {
	p.client.Close()
}
