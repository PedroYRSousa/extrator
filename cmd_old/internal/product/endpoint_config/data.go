package endpoint_config

import "net/http"

type endpointMethod string
type paginationMode string
type paginationDirection string
type paginationLocation string

const (
	ENDPOINT_RETRY_ATTEMPTS_DEFAULT         uint = 10
	ENDPOINT_RETRY_DELAY_IN_SECONDS_DEFAULT uint = 30

	ENDPOINT_ENDPOINT_CONFIG_METHOD_DEFAULT                        endpointMethod = http.MethodGet
	ENDPOINT_ENDPOINT_CONFIG_WAITING_TIME_IN_SECONDS_DEFAULT       uint           = 60
	ENDPOINT_ENDPOINT_CONFIG_WAITING_TIME_ERROR_IN_SECONDS_DEFAULT uint           = 30
	ENDPOINT_ENDPOINT_CONFIG_TIMEOUT_IN_SECONDS_DEFAULT            uint           = 30

	PAGINATION_MODE_NONE        paginationMode = "none"
	PAGINATION_MODE_OFFSET      paginationMode = "offset"
	PAGINATION_MODE_PAGE        paginationMode = "page"
	PAGINATION_MODE_PROPERTY    paginationMode = "property"
	PAGINATION_MODE_LINK_HEADER paginationMode = "link_header"

	PAGINATION_DIRECTION_NEXT     paginationDirection = "next"
	PAGINATION_DIRECTION_PREVIOUS paginationDirection = "previous"

	PAGINATION_LOCATION_HEADER      paginationLocation = "header"
	PAGINATION_LOCATION_BODY        paginationLocation = "body"
	PAGINATION_LOCATION_QUERY_PARAM paginationLocation = "query_param"
)

var (
	ENDPOINT_ENDPOINT_CONFIG_STATUS_CODE_WAITING_DEFAULT  = []uint{429}
	ENDPOINT_ENDPOINT_CONFIG_STATUS_CODE_SUCCESS_DEFAULT  = []uint{200, 201}
	ENDPOINT_ENDPOINT_CONFIG_STATUS_CODE_SKIPPING_DEFAULT = []uint{}
	METHODS_AVAILABLE                                     = []endpointMethod{http.MethodGet, http.MethodPost}
	AVAILABLE_LOCATIONS                                   = []paginationLocation{PAGINATION_LOCATION_HEADER, PAGINATION_LOCATION_BODY, PAGINATION_LOCATION_QUERY_PARAM}
	AVAILABLE_DIRECTIONS                                  = []paginationDirection{PAGINATION_DIRECTION_NEXT, PAGINATION_DIRECTION_PREVIOUS}
	AVAILABLE_MODES                                       = []paginationMode{PAGINATION_MODE_NONE, PAGINATION_MODE_OFFSET, PAGINATION_MODE_PAGE, PAGINATION_MODE_PROPERTY, PAGINATION_MODE_LINK_HEADER}
)

type S_PaginationOffset struct {
	Direction paginationDirection `yaml:"direction"`
	Location  paginationLocation  `yaml:"location"`
	Offset    string              `yaml:"offset"`
	Limit     string              `yaml:"limit"`
}

type S_PaginationPage struct {
	Direction paginationDirection `yaml:"direction"`
	Location  paginationLocation  `yaml:"location"`
	Page      string              `yaml:"page"`
	PageSize  int                 `yaml:"page_size"`
}

type S_PaginationProperty struct {
	Property string `yaml:"property"`
}

type S_PaginationLinkHeader struct {
	Header string `yaml:"header"`
}

type S_Pagination struct {
	Mode paginationMode `yaml:"mode"`

	// Opcionais
	Offset     *S_PaginationOffset     `yaml:"offset,omitempty"`
	Page       *S_PaginationPage       `yaml:"page,omitempty"`
	Property   *S_PaginationProperty   `yaml:"property,omitempty"`
	LinkHeader *S_PaginationLinkHeader `yaml:"link_header,omitempty"`
}

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
	Pagination                *S_Pagination      `yaml:"pagination,omitempty"`
}
