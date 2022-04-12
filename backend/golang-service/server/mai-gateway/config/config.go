package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"manso.live/backend/golang-service/pkg/util/configutil"
)

type Config struct {
	Log struct {
		Level string
	}

	Http struct {
		Port string
	}
}

func NewConfig(file string) (*Config, error) {
	if err := configutil.InitConfigForEnv(file); err != nil {
		return nil, errors.Wrap(err, "failed to configutil configuration from .env")
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal configuration to struct")
	}

	return &cfg, nil
}
