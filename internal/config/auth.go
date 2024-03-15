package config

import (
	"time"
)

type WebAuthConfig struct {
	PublicJwkUri      string        `env:"PUBLIC_JWK_URI" env-required:"true"`
	RefreshJwkTimeout time.Duration `env:"REFRESH_JWK_TIMEOUT" env-default:"3h"`
	ClientId          string        `env:"CLIENT_ID" env-required:"true"`
}

func NewAuthConfig() (*WebAuthConfig, error) {
	var cfg WebAuthConfig
	err := Load(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
