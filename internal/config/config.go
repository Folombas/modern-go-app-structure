package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"database"`
	App struct {
		Port int    `yaml:"port"`
		Env  string `yaml:"env"`
	} `yaml:"app"`
}

func LoadConfig(path string) (*Config, error) {
	cfg := &Config{}
	
	// Для Docker окружения используем переменные окружения
	if os.Getenv("APP_ENV") == "docker" {
		cfg.Database.Host = os.Getenv("DB_HOST")
		cfg.Database.Port = 5432
		cfg.Database.User = os.Getenv("DB_USER")
		cfg.Database.Password = os.Getenv("DB_PASSWORD")
		cfg.Database.DBName = os.Getenv("DB_NAME")
		cfg.Database.SSLMode = "disable"
		cfg.App.Port = 8080
		cfg.App.Env = "docker"
		return cfg, nil
	}
	
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()
	
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		return nil, fmt.Errorf("failed to decode config: %w", err)
	}
	
	return cfg, nil
}