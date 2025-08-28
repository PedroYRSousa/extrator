package config

const (
	DEFAULT_PRODUCTS_PATH = "../products"
	DEFAULT_CONFIG_PATH   = "../config.yaml"
)

type S_ConfigProducts struct {
	Path string `yaml:"path" validate:"required,dir,printascii"`
}

type S_Config struct {
	Products *S_ConfigProducts `yaml:"_products,omitempty"`
}
