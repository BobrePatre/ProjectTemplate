package config

import (
	"github.com/go-playground/validator/v10"
)

type AppConfig struct {
	ENV string `env:"ENV" json:"env" validate:"required"`
}

func NewAppConfig(validate *validator.Validate) (*AppConfig, error) {
	var cfg struct {
		Config AppConfig `env:"APP" json:"app"`
	}
	if err := Load(&cfg, validate); err != nil {
		return nil, err
	}

	return &cfg.Config, nil
}
