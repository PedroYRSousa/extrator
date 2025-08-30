package config

import (
	"extrator/modules"
)

func Load() (*S_Config, error) {
	var conf S_Config = newConfig()
	err := modules.YamlToStruct(DEFAULT_CONFIG_PATH, &conf)
	if err != nil {
		return nil, err
	}

	err = conf.check()
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
