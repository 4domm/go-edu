package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Port int `yaml:"port"`
}

func MustLoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
