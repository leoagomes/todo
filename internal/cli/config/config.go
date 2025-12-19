package config

import (
	"errors"
	"os"

	"github.com/goccy/go-yaml"
)

type Config struct {
	TasksDirectory string `yaml:"tasks-directory"`
	MetaDirectory  string `yaml:"meta-directory"`
}

var Default = Config{
	TasksDirectory: "tasks",
	MetaDirectory:  ".todo",
}

func LoadFile(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func LoadFileOrDefault(path string) (*Config, error) {
	cfg, err := LoadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return &Default, nil
		}
		return nil, err
	}
	return cfg, nil
}
