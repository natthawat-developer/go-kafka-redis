package main

import (
	"fmt"
	"time"
	"go-kafka-redis/config"
	"go-kafka-redis/pkg/kafka"
	"go-kafka-redis/pkg/logger"
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

	// สร้าง Kafka Producer
	producer, err := kafka.NewProducer(config.AppConfig.Kafka.BootstrapServers)
	if err != nil {
		// หากสร้าง Kafka producer ไม่สำเร็จ
		logger.ErrorLogger.Printf("Failed to create producer: %s", err)
		return
	}
	defer producer.Close()

	topic := "test_topic"

	// ส่งข้อความ 10 ข้อความ
	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("Message %d", i)
		err = producer.ProduceMessage(topic, message)
		if err != nil {
			// หากส่งข้อความไม่สำเร็จ
			logger.ErrorLogger.Printf("Failed to produce message: %s", err)
		} else {
			// หากส่งข้อความสำเร็จ
			logger.InfoLogger.Printf("Produced: %s", message)
		}
		time.Sleep(1 * time.Second)
	}

	// ฟังก์ชัน Flush
	producer.Flush()
}
