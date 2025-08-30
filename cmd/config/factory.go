package config

import config_products "extrator/config/products"

func newConfig() S_Config {
	return S_Config{
		Products: config_products.NewConfigProducts(),
	}
}
