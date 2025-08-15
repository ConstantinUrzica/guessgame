package main

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Path   string `json:"path"`
	DBPath string `json:"dbPath"`
}

func Load(path string) (*Config, error) {
	var cfg Config
	if path != "" {
		if err := cleanenv.ReadConfig(path, &cfg); err != nil {
			return nil, fmt.Errorf("read config: %w", err)
		}
	}
	return &cfg, nil
}
