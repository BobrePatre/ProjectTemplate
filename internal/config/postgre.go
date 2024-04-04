package config

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type PostgresqlConfig struct {
	Host     string `env:"HOST" json:"host" validate:"required"`
	Port     int    `env:"PORT" json:"port" env-default:"5432"`
	User     string `env:"USER" json:"user" env-default:"postgres"`
	Password string `env:"PASSWORD" json:"password" env-default:"postgres"`
	Database string `env:"DATABASE" json:"database" env-default:"postgres"`
}

func NewPostgresConfig(validate *validator.Validate) (*PostgresqlConfig, error) {
	var cfg struct {
		Config PostgresqlConfig `json:"postgres" env-prefix:"POSTGRES_"`
	}
	if err := Load(&cfg, validate); err != nil {
		return nil, err
	}
	return &cfg.Config, nil
}

func (cfg *PostgresqlConfig) Datasource(sslMode string) string {
	return fmt.Sprintf(
		"host=%s port=%v user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Database,
		cfg.Password,
		sslMode,
	)
}
