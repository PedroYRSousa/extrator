package config_products

func New() *S_ConfigProducts {
	return &S_ConfigProducts{
		Path:                         DEFAULT_PRODUCTS_PATH,
		ProductAsyncCount:            DEFAULT_PRODUCT_ASYNC_COUNT,
		EndpointPerProductAsyncCount: DEFAULT_ENDPOINT_PER_PRODUCT_ASYNC_COUNT,
		MaxAttempts:                  DEFAULT_MAX_ATTEMPTS,
		DelayAttemptsInSeconds:       DEFAULT_DELAY_ATTEMPTS_IN_SECONDS,
	}
}
