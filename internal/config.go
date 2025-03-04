package internal

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Adjectives []string `yaml:"adjectives"`
	Nouns      []string `yaml:"nouns"`
}

func LoadConfig(filename string) (*Config, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file %s does not exist", filename)
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	if err := config.validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return &config, nil
}

func (c *Config) validate() error {
	if len(c.Adjectives) == 0 {
		return fmt.Errorf("adjectives list is empty")
	}
	if len(c.Nouns) == 0 {
		return fmt.Errorf("nouns list is empty")
	}
	return nil
}

func (c *Config) GetRandomNoun() string {
	return c.Nouns[Random(len(c.Nouns))]
}

func (c *Config) GetRandomAdjective() string {
	return c.Adjectives[Random(len(c.Adjectives))]
}
