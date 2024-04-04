package config

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"net"
)

type RedisConfig struct {
	Host     string `env:"HOST" json:"host" validate:"required"`
	Port     int    `env:"PORT" json:"port" env-default:"6379"`
	Password string `env:"PASSWORD" json:"password" validate:"required"`
	Database int    `env:"DATABASE" json:"database" env-default:"0"`
}

func NewRedisConfig(validate *validator.Validate) (*RedisConfig, error) {
	var cfg struct {
		Config RedisConfig `json:"redis" env-prefix:"REDIS_"`
	}

	if err := Load(&cfg, validate); err != nil {
		return nil, err
	}

	return &cfg.Config, nil

}

func (r RedisConfig) Address() string {
	return net.JoinHostPort(r.Host, fmt.Sprint(r.Port))
}
