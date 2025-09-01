package config_products

const (
	DEFAULT_PRODUCTS_PATH                           = "../products"
	DEFAULT_PRODUCT_ASYNC_COUNT              uint64 = 3
	DEFAULT_ENDPOINT_PER_PRODUCT_ASYNC_COUNT uint64 = 10
	DEFAULT_MAX_ATTEMPTS                     uint64 = 10
	DEFAULT_DELAY_ATTEMPTS_IN_SECONDS        uint64 = 30
)

type S_ConfigProducts struct {
	Path                         string `yaml:"path" validate:"required,dirpath,printascii"`
	ProductAsyncCount            uint64 `yaml:"_product_async_count,omitempty"`
	EndpointPerProductAsyncCount uint64 `yaml:"_endpoint_per_product_async_count,omitempty"`
	MaxAttempts                  uint64 `yaml:"_max_attempts,omitempty"`
	DelayAttemptsInSeconds       uint64 `yaml:"_delay_attempts_in_seconds,omitempty"`
}
