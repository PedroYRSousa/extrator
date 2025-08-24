package endpoint

type endpointResponseFormat string

const (
	ENDPOINT_RESPONSE_FORMAT_JSON endpointResponseFormat = "json"
)

var responsesFormatAvailable = []endpointResponseFormat{ENDPOINT_RESPONSE_FORMAT_JSON}

type S_Retry struct {
	Attempts       int  `yaml:"attempts"`
	DelayInSeconds int  `yaml:"delay_in_seconds"`
	BackoffFactor  *int `yaml:"backoff_factor,omitempty"`
}

type S_EndpointConfig struct {
	StatusCodeWaiting         []int `yaml:"status_code_waiting"`
	StatusCodeSuccess         []int `yaml:"status_code_success"`
	StatusCodeSkip            []int `yaml:"status_code_skip"`
	WaitingTimeInSeconds      int   `yaml:"waiting_time_in_seconds"`
	WaitingTimeErrorInSeconds int   `yaml:"waiting_time_in_error_in_seconds"`

	// Opcionais
	TimeoutInSeconds *int               `yaml:"timeout_in_seconds,omitempty"`
	Retry            *S_Retry           `yaml:"retry,omitempty"`
	Headers          *map[string]string `yaml:"headers,omitempty"`
	Body             *map[string]string `yaml:"body,omitempty"`
	QueryParams      *map[string]string `yaml:"query_params,omitempty"`
}

type S_Endpoint struct {
	URI            string                 `yaml:"uri"`
	Method         string                 `yaml:"method"`
	ResponseFormat endpointResponseFormat `yaml:"response_format"`
	EndpointConfig S_EndpointConfig       `yaml:"endpoint_config"`

	// Opcionais
	LimitExtract *int `yaml:"limit_extract,omitempty"`
}
