package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/baking-bread/taxonomist/internal/random"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Dictionary struct {
		Adjectives []string `yaml:"adjectives"`
		Nouns      []string `yaml:"nouns"`
	} `yaml:"dictionary"`
	mu sync.RWMutex
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
	if len(c.Dictionary.Adjectives) == 0 {
		return fmt.Errorf("adjectives list is empty")
	}
	if len(c.Dictionary.Nouns) == 0 {
		return fmt.Errorf("nouns list is empty")
	}
	return nil
}

func (c *Config) GetRandomNoun() string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.Dictionary.Nouns[random.Random(len(c.Dictionary.Nouns))]
}

func (c *Config) GetRandomAdjective() string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.Dictionary.Adjectives[random.Random(len(c.Dictionary.Adjectives))]
}
