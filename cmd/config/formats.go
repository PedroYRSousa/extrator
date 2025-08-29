package config

import (
	"extrator/modules"
)

func (configProducts *S_ConfigProducts) _format() *S_ConfigProducts {
	if configProducts == nil {
		configProducts = newConfigProducts()
	}

	modules.FormatString(configProducts)

	if configProducts.ProductAsyncCount == nil {
		configProducts.ProductAsyncCount = new(uint)
		*configProducts.ProductAsyncCount = DEFAULT_PRODUCT_ASYNC_COUNT
	}

	if configProducts.EndpointPerProductAsyncCount == nil {
		configProducts.EndpointPerProductAsyncCount = new(uint)
		*configProducts.EndpointPerProductAsyncCount = DEFAULT_ENDPOINT_PER_PRODUCT_ASYNC_COUNT
	}

	if configProducts.MaxAttempts == nil {
		configProducts.MaxAttempts = new(uint)
		*configProducts.MaxAttempts = DEFAULT_MAX_ATTEMPTS
	}

	if configProducts.DelayAttemptsInSeconds == nil {
		configProducts.DelayAttemptsInSeconds = new(uint)
		*configProducts.DelayAttemptsInSeconds = DEFAULT_DELAY_ATTEMPTS_IN_SECONDS
	}

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
