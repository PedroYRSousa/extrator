package config

func newConfigProducts() *S_ConfigProducts {
	var productAsyncCount = DEFAULT_PRODUCT_ASYNC_COUNT
	var endpointPerProductAsyncCount = DEFAULT_ENDPOINT_PER_PRODUCT_ASYNC_COUNT
	var maxAttempts = DEFAULT_MAX_ATTEMPTS
	var delayAttemptsInSeconds = DEFAULT_DELAY_ATTEMPTS_IN_SECONDS

	return &S_ConfigProducts{
		Path:                         DEFAULT_PRODUCTS_PATH,
		ProductAsyncCount:            &productAsyncCount,
		EndpointPerProductAsyncCount: &endpointPerProductAsyncCount,
		MaxAttempts:                  &maxAttempts,
		DelayAttemptsInSeconds:       &delayAttemptsInSeconds,
	}
}

func newConfig() *S_Config {
	return &S_Config{
		Products: newConfigProducts(),
	}
}
