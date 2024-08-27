package fistats

import (
	"github.com/krabiworld/fistats/fistorage"
)

type Config struct {
	Storage fistorage.Storage
}

func configDefault(config ...Config) Config {
	cfg := config[0]

	if cfg.Storage == nil {
		cfg.Storage = fistorage.NewMemory()
	}

	return cfg
}
