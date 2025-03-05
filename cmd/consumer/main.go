package main

import (
	"fmt"
	"log"
	"go-kafka-redis/config"
	"go-kafka-redis/pkg/kafka"
	"go-kafka-redis/pkg/redis"
)

func main() {
	// โหลดค่า config
	config.LoadConfig("config/config.yaml")

	// เชื่อมต่อ Redis
	redis.InitRedis(config.AppConfig.Redis.Addr)
	defer redis.CloseRedis()

	// สร้าง Kafka Consumer
	consumer, err := kafka.NewConsumer(
		config.AppConfig.Kafka.BootstrapServers,
		config.AppConfig.Kafka.GroupID,
		config.AppConfig.Kafka.AutoOffsetReset,
	)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %s", err)
	}
	defer consumer.Close()

	// Subscribe to topic
	topic := "test_topic"
	consumer.SubscribeTopic(topic)

	// รับข้อความจาก Kafka และบันทึกลง Redis
	fmt.Println("Waiting for messages...")
	for {
		msg, err := consumer.ReadMessage()
		if err == nil {
			fmt.Printf("Received: %s\n", string(msg.Value))
			redis.SaveToRedis(string(msg.Key), string(msg.Value))
		} else {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
