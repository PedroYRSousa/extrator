package configs

const (
	DEFAULT_PATH_CONFIGS  = "../configs/config.yaml"
	DEFAULT_PATH_PRODUCTS = "../configs/products"
)

type S_Configs struct {
	PathProducts string `yaml:"path_products" validate:"required"`
}
