package main

import (
	"fmt"
	"log"
	"time"

	"go-kafka-redis/config"
	"go-kafka-redis/pkg/kafka"
)

func main() {
	// โหลดค่า config
	config.LoadConfig("config/config.yaml")

	// สร้าง Kafka Producer
	producer, err := kafka.NewProducer(config.AppConfig.Kafka.BootstrapServers)
	if err != nil {
		log.Fatalf("Failed to create producer: %s", err)
	}
	defer producer.Close()

	topic := "test_topic"

	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("Message %d", i)
		err = producer.ProduceMessage(topic, message)
		if err != nil {
			log.Printf("Failed to produce message: %s", err)
		} else {
			fmt.Printf("Produced: %s\n", message)
		}
		time.Sleep(1 * time.Second)
	}

	producer.Flush()
}
