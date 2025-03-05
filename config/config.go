package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"log"
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

func LoadConfig(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read config file: %s", err)
	}

	err = yaml.Unmarshal(data, &AppConfig)
	if err != nil {
		log.Fatalf("Failed to parse config file: %s", err)
	}
}