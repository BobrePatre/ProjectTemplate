package config

import (
	"net"
)

type RedisConfig struct {
	Host     string `env:"REDIS_HOST" env-default:"localhost"`
	Port     string `env:"REDIS_PORT" env-default:"6379"`
	Password string `env:"REDIS_PASS" env-required:"true"`
	Database int    `env:"REDIS_DB" env-default:"0"`
}

func NewRedisConfig() (*RedisConfig, error) {
	var cfg RedisConfig
	err := Load(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (r RedisConfig) Address() string {
	return net.JoinHostPort(r.Host, r.Port)
}
