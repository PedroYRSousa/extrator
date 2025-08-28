package config

import (
	"github.com/go-playground/validator"
)

func (config *S_Config) validate() error {
	validate := validator.New()

	err := validate.Struct(*config)
	if err != nil {
		return err
	}

	return nil
}
