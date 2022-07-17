package config

import (
	"errors"

	"github.com/morfin60/parallel-handlers/internal/helpers/environment"
)

type Config struct {
	Address       string `json:"address"`
	RequestsLimit int32  `json:"requests_limit"`
	PoolSize      int32  `json:"pool_size"`
	Debug         bool   `json:"debug"`
}

func New() (*Config, error) {
	config := &Config{}

	config.Address = environment.GetString("PARALLEL_HANDLERS_ADDRESS", ":8888")
	config.RequestsLimit = environment.GetInt32("PARALLEL_HANDLERS_REQUESTS_LIMIT", 100)
	config.PoolSize = environment.GetInt32("PARALLEL_HANDLERS_POOL_SIZE", 4)
	config.Debug = environment.GetBool("PARALLEL_HANDLERS_DEBUG", false)

	if config.Address == "" {
		return nil, errors.New("Incorrect address specified(should be address:port)")
	}

	return config, nil
}
