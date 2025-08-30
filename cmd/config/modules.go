package config

import (
	"extrator/modules"

	"github.com/go-playground/validator/v10"
)

func (config *S_Config) format() {
	if config == nil {
		panic("TODO, melhorar | ERROR cheformatck config")
	}

	modules.FormatString(config)
}

func (config *S_Config) transform() error {
	if config == nil {
		panic("TODO, melhorar | ERROR transform config")
	}

	err := modules.TransformString(config)
	if err != nil {
		return err
	}

	return nil
}

func (config *S_Config) validate() error {
	if config == nil {
		panic("TODO, melhorar | ERROR validate config")
	}

	err := validator.New().Struct(config)
	if err != nil {
		return err
	}

	return nil
}

func (config *S_Config) check() error {
	if config == nil {
		panic("TODO, melhorar | ERROR check config")
	}

	config.format()

	err := config.validate()
	if err != nil {
		return err
	}

	err = config.transform()
	if err != nil {
		return err
	}

	return nil
}
