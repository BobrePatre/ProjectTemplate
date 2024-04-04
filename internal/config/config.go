package config

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
)

func Load(cfg interface{}, validate *validator.Validate) error {

	err := cleanenv.ReadConfig("config.json", cfg)
	if err != nil {
		return fmt.Errorf("error with reading config: %v", err)
	}

	err = validate.Struct(cfg)
	if err != nil {
		return fmt.Errorf("error with validationg config: %v", err)
	}

	return nil
}
