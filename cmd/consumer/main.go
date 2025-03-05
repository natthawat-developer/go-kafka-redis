package main

import (
	"go-kafka-redis/config"
	"go-kafka-redis/pkg/kafka"
	"go-kafka-redis/pkg/logger"
	"go-kafka-redis/pkg/redis"
)

func main() {
	// เริ่มต้น logger
	logger.InitLogger()

	// โหลดค่า config
	err := config.LoadConfig("config/config.yaml")
	if err != nil {
		// หากโหลด config ไม่สำเร็จ
		logger.ErrorLogger.Printf("Failed to load config: %s", err)
		return
	}

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
		// หากสร้าง Kafka consumer ไม่สำเร็จ
		logger.ErrorLogger.Printf("Failed to create Kafka consumer: %s", err)
		return
	}
	defer consumer.Close()

	// Subscribe to topic
	topic := "test_topic"
	consumer.SubscribeTopic(topic)

	// รับข้อความจาก Kafka และบันทึกลง Redis
	logger.InfoLogger.Println("Waiting for messages...")
	for {
		msg, err := consumer.ReadMessage()
		if err == nil {
			// เมื่อรับข้อความจาก Kafka
			logger.InfoLogger.Printf("Received: %s", string(msg.Value))
			redis.SaveToRedis(string(msg.Key), string(msg.Value))
		} else {
			// เมื่อเกิดข้อผิดพลาดจาก Kafka
			logger.ErrorLogger.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
