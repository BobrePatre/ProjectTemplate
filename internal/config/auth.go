package config

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type WebAuthConfig struct {
	PublicJwkUri      string        `env:"PUBLIC_JWK_URI" json:"publicJwkUri" validate:"required"`
	RefreshJwkTimeout time.Duration `env:"REFRESH_JWK_TIMEOUT" json:"refreshJwkTimeout" env-default:"3h"`
	ClientId          string        `env:"CLIENT_ID" json:"clientId" validate:"required"`
}

func NewAuthConfig(validate *validator.Validate) (*WebAuthConfig, error) {
	var cfg struct {
		Config WebAuthConfig `env-prefix:"AUTH_" json:"auth"`
	}
	if err := Load(&cfg, validate); err != nil {
		return nil, err
	}

	return &cfg.Config, nil
}
