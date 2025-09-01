package config

import config_products "extrator/config/products"

func New() *S_Config {
	return &S_Config{
		Products: config_products.New(),
	}
}
