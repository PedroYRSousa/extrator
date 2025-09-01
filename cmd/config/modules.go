package config

import (
	"extrator/modules"
	"extrator/pkg/struct_inspector"

	"github.com/go-playground/validator/v10"
)

func (config *S_Config) validate() error {
	if config == nil {
		panic("ERROR Interno, verificar config.validate")
	}

	err := validator.New().Struct(config)
	if err != nil {
		errValidator, ok := validator.New().Struct(config).(validator.ValidationErrors)
		if !ok {
			panic("ERROR Interno, verificar config.validate")
		}

		return modules.TranslateValidatorError(errValidator)
	}

	return nil
}

func (config *S_Config) check() error {
	if config == nil {
		panic("ERROR Interno, verificar config.validate")
	}

	err := struct_inspector.Start(config)
	if err != nil {
		return err
	}

	err = config.validate()
	if err != nil {
		return err
	}

	return nil
}
