package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	TasksDirectory string `mapstructure:"tasks-directory"`
	MetaDirectory  string `mapstructure:"meta-directory"`
}

// Load initializes and returns a Config using Viper
// It will load from the specified path if it exists, otherwise use defaults
func Load(configPath string) (*Config, error) {
	v := viper.New()

	// Set defaults
	v.SetDefault("tasks-directory", "tasks")
	v.SetDefault("meta-directory", ".todo")

	// Configure Viper
	if configPath != "" {
		v.SetConfigFile(configPath)
	} else {
		// Look for config in common locations
		v.SetConfigName(".todo")
		v.SetConfigType("yaml")
		v.AddConfigPath(".")
		v.AddConfigPath("$HOME")
	}

	// Read config file (ignore error if file doesn't exist)
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// Config file was found but another error occurred
			return nil, err
		}
		// Config file not found; use defaults
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
