package config

import (
	"net"
)

type HttpConfig struct {
	Port string `env:"HTTP_PORT" env-default:"8080"`
}

func NewHTTPConfig() (*HttpConfig, error) {
	var cfg HttpConfig
	err := Load(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (cfg HttpConfig) Address() string {
	return net.JoinHostPort("localhost", cfg.Port)
}
