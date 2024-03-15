package config

type AppConfig struct {
	ENV string `env:"ENV" env-required:"true"`
}

func NewAppConfig() (*AppConfig, error) {
	var cfg AppConfig
	err := Load(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
