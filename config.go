package fistats

import (
	"github.com/krabiworld/fistats/fistorage"
)

type Config struct {
	Storage fistorage.Storage
}

var ConfigDefault = Config{
	Storage: fistorage.NewMemory(),
}

func configDefault(config ...Config) Config {
	if len(config) < 1 {
		return ConfigDefault
	}

	cfg := config[0]

	if cfg.Storage == nil {
		cfg.Storage = fistorage.NewMemory()
	}

	return cfg
}
