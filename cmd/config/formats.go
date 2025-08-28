package config

import (
	"extrator/modules"
)

func (configProducts *S_ConfigProducts) _format() *S_ConfigProducts {
	if configProducts == nil {
		configProducts = newConfigProducts()
	}

	modules.FormatString(configProducts)

	return configProducts
}

func (config *S_Config) format() *S_Config {
	if config == nil {
		config = newConfig()
	}

	config.Products = config.Products._format()

	modules.FormatString(config)

	return config
}
