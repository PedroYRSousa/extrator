package config

func newConfigProducts() *S_ConfigProducts {
	return &S_ConfigProducts{
		Path: DEFAULT_PRODUCTS_PATH,
	}
}

func newConfig() *S_Config {
	return &S_Config{
		Products: newConfigProducts(),
	}
}
