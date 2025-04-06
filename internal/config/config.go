package config

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Server        `yaml:"server"`
	Database      `yaml:"database"`
	SessionSecret string `yaml:"session_secret"`
	YandexGPT     `yaml:"yandex_gpt"`
	Prompts       `yaml:"prompts"`
	Mongo         `yaml:"mongo"`
	ExerciseUrl   string `yaml:"exercise_url"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	Timezone string `yaml:"timezone"`
}

type YandexGPT struct {
	CatalogToken string `yaml:"catalog_token"`
	ApiToken     string `yaml:"api_token"`
	URL          string `yaml:"url"`
}

type Prompts struct {
	Plans string `yaml:"plans"`
	Chat  string `yaml:"chat"`
}

type Mongo struct {
	URI    string `yaml:"uri"`
	DBName string `yaml:"db"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	yamlPath := os.Getenv("CONFIG_PATH")
	if yamlPath == "" {
		return nil, errors.New("config file environment variable not set")
	}
	yamlFile, err := os.ReadFile(yamlPath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error read config file: %s", err))
	}
	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error parse config file: %s", err))
	}
	return cfg, nil
}
