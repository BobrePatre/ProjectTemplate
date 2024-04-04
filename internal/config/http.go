package config

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"net"
)

type HttpConfig struct {
	Port int `json:"port" env:"PORT" env-default:"8080"`
}

func NewHTTPConfig(validate *validator.Validate) (*HttpConfig, error) {

	var cfg struct {
		Config HttpConfig `env-prefix:"HTTP_" json:"http"`
	}
	if err := Load(&cfg, validate); err != nil {
		return nil, err
	}

	return &cfg.Config, nil
}

func (cfg *HttpConfig) Address() string {
	return net.JoinHostPort("localhost", fmt.Sprint(cfg.Port))
}
