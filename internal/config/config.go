package config

import "github.com/ilyakaznacheev/cleanenv"

func Load(cfg interface{}) error {
	err := cleanenv.ReadConfig(".env", cfg)
	err = cleanenv.ReadEnv(cfg)
	return err
}
