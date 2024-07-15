package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Interval string `yaml:"interval"`
	Poster   struct {
		Options struct {
			BotTokenEnv string `yaml:"bot_token_env"`
			ChannelName string `yaml:"channel_name"`
		} `yaml:"options"`
	} `yaml:"poster"`
}

func readConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("can't open config file: %w", err)
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, fmt.Errorf("can't parse config file: %w", err)
	}

	return config, nil
}
