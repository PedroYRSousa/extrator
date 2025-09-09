package configs

const (
	DEFAULT_PATH_CONFIGS                       = "../configs/config.yaml"
	DEFAULT_PATH_PRODUCTS                      = "../configs/products"
	DEFAULT_PRODUCTS_PARALLEL_EXECUTION_COUNT  = 3
	DEFAULT_ENDPOINTS_PARALLEL_EXECUTION_COUNT = 3
)

type S_Configs struct {
	PathProducts                    string `yaml:"path_products" validate:"required"`
	ProductsParallelExecutionCount  int    `yaml:"_products_parallel_execution_count" validate:"gt=0"`
	EndpointsParallelExecutionCount int    `yaml:"_endpoints_parallel_execution_count" validate:"gt=0"`
}
