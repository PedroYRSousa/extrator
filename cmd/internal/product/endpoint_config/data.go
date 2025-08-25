package endpoint_config

import "net/http"

type endpointMethod string

const (
	ENDPOINT_RETRY_ATTEMPTS_DEFAULT         uint = 10
	ENDPOINT_RETRY_DELAY_IN_SECONDS_DEFAULT uint = 30

	ENDPOINT_ENDPOINT_CONFIG_METHOD_DEFAULT                        endpointMethod = http.MethodGet
	ENDPOINT_ENDPOINT_CONFIG_WAITING_TIME_IN_SECONDS_DEFAULT       uint           = 60
	ENDPOINT_ENDPOINT_CONFIG_WAITING_TIME_ERROR_IN_SECONDS_DEFAULT uint           = 30
	ENDPOINT_ENDPOINT_CONFIG_TIMEOUT_IN_SECONDS_DEFAULT            uint           = 30
)

var (
	ENDPOINT_ENDPOINT_CONFIG_STATUS_CODE_WAITING_DEFAULT  = []uint{429}
	ENDPOINT_ENDPOINT_CONFIG_STATUS_CODE_SUCCESS_DEFAULT  = []uint{200, 201}
	ENDPOINT_ENDPOINT_CONFIG_STATUS_CODE_SKIPPING_DEFAULT = []uint{}
	METHODS_AVAILABLE                                     = []endpointMethod{http.MethodGet, http.MethodPost}
)

type S_Retry struct {
	// Opcionais
	Attempts       *uint `yaml:"attempts,omitempty"`
	DelayInSeconds *uint `yaml:"delay_in_seconds,omitempty"`
	BackoffFactor  *uint `yaml:"backoff_factor,omitempty"`
}

type S_EndpointConfig struct {
	// Opcionais
	Method                    *endpointMethod    `yaml:"method,omitempty"`
	StatusCodeWaiting         *[]uint            `yaml:"status_code_waiting,omitempty"`
	StatusCodeSuccess         *[]uint            `yaml:"status_code_success,omitempty"`
	StatusCodeSkip            *[]uint            `yaml:"status_code_skip,omitempty"`
	WaitingTimeInSeconds      *uint              `yaml:"waiting_time_in_seconds,omitempty"`
	WaitingTimeErrorInSeconds *uint              `yaml:"waiting_time_error_in_seconds,omitempty"`
	TimeoutInSeconds          *uint              `yaml:"timeout_in_seconds,omitempty"`
	Retry                     *S_Retry           `yaml:"retry,omitempty"`
	Headers                   *map[string]string `yaml:"headers,omitempty"`
	Body                      *map[string]string `yaml:"body,omitempty"`
	QueryParams               *map[string]string `yaml:"query_params,omitempty"`
}
