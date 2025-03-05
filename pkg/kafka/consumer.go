package kafka

import (
	"go-kafka-redis/pkg/logger"
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
		// ใช้ logger ในการบันทึกข้อผิดพลาดแทนการใช้ log.Fatal
		logger.ErrorLogger.Printf("Failed to create consumer: %s", err)
		return nil, err
	}

	return &Consumer{client: client}, nil
}

func (c *Consumer) SubscribeTopic(topic string) {
	err := c.client.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		// ใช้ logger ในการบันทึกข้อผิดพลาดแทนการใช้ log.Fatalf
		logger.ErrorLogger.Printf("Failed to subscribe to topic: %s", err)
	}
}

func (c *Consumer) ReadMessage() (*kafka.Message, error) {
	msg, err := c.client.ReadMessage(-1)
	if err != nil {
		// ใช้ logger ในการบันทึกข้อผิดพลาดหากมีปัญหาการอ่านข้อความ
		logger.ErrorLogger.Printf("Failed to read message: %s", err)
		return nil, err
	}

	// ใช้ logger ในการบันทึกข้อความที่อ่านได้ (ถ้าต้องการ)
	logger.InfoLogger.Printf("Received message: %s", string(msg.Value))

	return msg, nil
}

func (c *Consumer) Close() {
	c.client.Close()
	// ใช้ logger ในการบันทึกเมื่อปิด consumer
	logger.InfoLogger.Println("Consumer closed")
}
