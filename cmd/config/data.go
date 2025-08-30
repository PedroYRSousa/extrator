package config

import config_products "extrator/config/products"

const (
	DEFAULT_CONFIG_PATH = "../config.yaml"
)

type S_Config struct {
	Products config_products.S_ConfigProducts `yaml:"_products,omitempty"`
}
