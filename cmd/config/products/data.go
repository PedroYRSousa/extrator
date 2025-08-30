package config_products

const (
	DEFAULT_PRODUCTS_PATH                         = "../products"
	DEFAULT_PRODUCT_ASYNC_COUNT              uint = 3
	DEFAULT_ENDPOINT_PER_PRODUCT_ASYNC_COUNT uint = 10
	DEFAULT_MAX_ATTEMPTS                     uint = 10
	DEFAULT_DELAY_ATTEMPTS_IN_SECONDS        uint = 30
)

type S_ConfigProducts struct {
	Path                         string `yaml:"path" validate:"required,dirpath,printascii"`
	ProductAsyncCount            uint   `yaml:"_product_async_count,omitempty"`
	EndpointPerProductAsyncCount uint   `yaml:"_endpoint_per_product_async_count,omitempty"`
	MaxAttempts                  uint   `yaml:"_max_attempts,omitempty"`
	DelayAttemptsInSeconds       uint   `yaml:"_delay_attempts_in_seconds,omitempty"`
}
