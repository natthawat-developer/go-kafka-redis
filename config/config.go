package config

import (
	"go-kafka-redis/pkg/logger" 
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Kafka struct {
		BootstrapServers string `yaml:"bootstrap_servers"`
		GroupID          string `yaml:"group_id"`
		AutoOffsetReset  string `yaml:"auto_offset_reset"`
	} `yaml:"kafka"`
	Redis struct {
		Addr string `yaml:"addr"`
	} `yaml:"redis"`
}

var AppConfig Config

// เปลี่ยนให้ฟังก์ชัน LoadConfig คืนค่า error
func LoadConfig(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		// ใช้ logger แทน log.Fatalf
		logger.ErrorLogger.Printf("Failed to read config file: %s", err)
		return err // คืนค่า error
	}

	err = yaml.Unmarshal(data, &AppConfig)
	if err != nil {
		// ใช้ logger แทน log.Fatalf
		logger.ErrorLogger.Printf("Failed to parse config file: %s", err)
		return err // คืนค่า error
	}

	// ถ้าการโหลด config สำเร็จ
	logger.InfoLogger.Println("Config loaded successfully")
	return nil // คืนค่า nil ถ้าไม่มีข้อผิดพลาด
}
