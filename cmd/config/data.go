package config

const (
	DEFAULT_PRODUCTS_PATH = "../products"
	DEFAULT_CONFIG_PATH   = "../config.yaml"
)

type S_ConfigProducts struct {
	// Obrigatórios
	Path string `yaml:"path"`
	// Opcionais
}

type S_Config struct {
	// Obrigatórios
	// Opcionais
	Products *S_ConfigProducts `yaml:"_products,omitempty"`
}
