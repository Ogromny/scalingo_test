package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Config struct {
	Port     int `default:"5000"`
	Username string
	Token    string // Oauth or PAT
}

func NewConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, errors.Wrapf(err, "fail to build config from env")
	}
	return &cfg, nil
}
