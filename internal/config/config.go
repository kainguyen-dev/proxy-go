package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Server      ServerConfig      `yaml:"server"`
	Integration IntegrationConfig `yaml:"integration"`
	Database    DatabaseConfig    `yaml:"database"`
}

var config *Config

func Init(env string) {
	if env == "" {
		log.Fatalf("Empty env variable")
	}
	var err error
	file, err := os.ReadFile("internal/config/file/" + env + ".yaml")
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
}

func GetConfig() *Config {
	return config
}
