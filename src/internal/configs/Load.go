package configs

import (
	"errors"
	"extrator/pkg/validator"
	"os"

	"gopkg.in/yaml.v3"
)

func Load() (*S_Configs, error) {
	configs := New()

	fileData, err := os.ReadFile(DEFAULT_PATH_CONFIGS)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(fileData, configs); err != nil {
		return nil, err
	}

	if err := validator.New().Struct(configs); err != nil {
		return nil, errors.New(validator.FormatErrors(err, configs))
	}

	return configs, nil
}
