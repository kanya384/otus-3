package config

import (
	"fmt"

	"github.com/vrischmann/envconfig"
)

type Config struct {
	DB struct {
		User     string
		Password string
		Port     string
		Host     string
		PoolMax  int
		Name     string
		Timeout  int
	}
}

func InitConfig(prefix string) (*Config, error) {
	conf := &Config{}
	if err := envconfig.InitWithPrefix(conf, prefix); err != nil {
		return nil, fmt.Errorf("init config error: %w", err)
	}

	return conf, nil
}
