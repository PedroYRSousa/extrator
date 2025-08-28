package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func Load() (*S_Config, error) {
	configBytes, err := os.ReadFile(DEFAULT_CONFIG_PATH)
	if err != nil {
		return nil, err
	}

	var conf S_Config
	err = yaml.Unmarshal(configBytes, &conf)
	if err != nil {
		return nil, err
	}

	log.Println(conf)

	err = conf.check()
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
