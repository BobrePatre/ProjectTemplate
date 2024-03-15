package config

import (
	"net"
)

type GRPCConfig struct {
	Port string `env:"GRPC_PORT" env-default:"50051"`
}

func NewGRPCConfig() (*GRPCConfig, error) {
	var cfg GRPCConfig
	err := Load(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (cfg GRPCConfig) Address() string {
	return net.JoinHostPort("localhost", cfg.Port)
}
