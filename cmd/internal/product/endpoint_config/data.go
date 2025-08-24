package endpointconfig

const (
	ENDPOINT_RETRY_ATTEMPTS_DEFAULT         = 10
	ENDPOINT_RETRY_DELAY_IN_SECONDS_DEFAULT = 30

	ENDPOINT_ENDPOINT_CONFIG_WAITING_TIME_IN_SECONDS_DEFAULT       = 60
	ENDPOINT_ENDPOINT_CONFIG_WAITING_TIME_ERROR_IN_SECONDS_DEFAULT = 30
	ENDPOINT_ENDPOINT_CONFIG_TIMEOUT_IN_SECONDS_DEFAULT            = 30
)

var (
	ENDPOINT_ENDPOINT_CONFIG_STATUS_CODE_WAITING_DEFAULT  = []int{429}
	ENDPOINT_ENDPOINT_CONFIG_STATUS_CODE_SUCCESS_DEFAULT  = []int{200}
	ENDPOINT_ENDPOINT_CONFIG_STATUS_CODE_SKIPPING_DEFAULT = []int{}
)

type S_Retry struct {
	// Opcionais
	Attempts       *int `yaml:"attempts,omitempty"`
	DelayInSeconds *int `yaml:"delay_in_seconds,omitempty"`
	BackoffFactor  *int `yaml:"backoff_factor,omitempty"`
}

type S_EndpointConfig struct {
	// Opcionais
	StatusCodeWaiting         *[]int             `yaml:"status_code_waiting,omitempty"`
	StatusCodeSuccess         *[]int             `yaml:"status_code_success,omitempty"`
	StatusCodeSkip            *[]int             `yaml:"status_code_skip,omitempty"`
	WaitingTimeInSeconds      *int               `yaml:"waiting_time_in_seconds,omitempty"`
	WaitingTimeErrorInSeconds *int               `yaml:"waiting_time_error_in_seconds,omitempty"`
	TimeoutInSeconds          *int               `yaml:"timeout_in_seconds,omitempty"`
	Retry                     *S_Retry           `yaml:"retry,omitempty"`
	Headers                   *map[string]string `yaml:"headers,omitempty"`
	Body                      *map[string]string `yaml:"body,omitempty"`
	QueryParams               *map[string]string `yaml:"query_params,omitempty"`
}
