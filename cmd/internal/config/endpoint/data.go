package endpoint

type S_Pagination struct {
	Direction  string  `yaml:"direction"`
	Property   string  `yaml:"property"`
	Header     string  `yaml:"header"`
	Location   string  `yaml:"location"`
	Offset     string  `yaml:"offset"`
	Limit      string  `yaml:"limit"`
	LimitValue int     `yaml:"limit_value"`
	Page       string  `yaml:"page"`
	PageSize   int     `yaml:"page_size"`
	Mode       *string `yaml:"mode,omitempty"`
}

type S_Retry struct {
	Attempts       int  `yaml:"attempts"`
	DelayInSeconds int  `yaml:"delay_in_seconds"`
	BackoffFactor  *int `yaml:"backoff_factor,omitempty"`
}

type S_QueryParam struct {
	Name  string  `yaml:"name"`
	Value *string `yaml:"value,omitempty"`
}

type S_EndpointConfig struct {
	StatusCodeWaiting         []int              `yaml:"status_code_waiting"`
	StatusCodeSuccess         []int              `yaml:"status_code_success"`
	StatusCodeSkip            []int              `yaml:"status_code_skip"`
	WaitingTimeInSeconds      int                `yaml:"waiting_time_in_seconds"`
	WaitingTimeErrorInSeconds int                `yaml:"waiting_time_in_error_in_seconds"`
	TimeoutInSeconds          *int               `yaml:"timeout_in_seconds,omitempty"`
	Retry                     *S_Retry           `yaml:"retry,omitempty"`
	Pagination                *S_Pagination      `yaml:"pagination,omitempty"`
	Headers                   *map[string]string `yaml:"headers,omitempty"`
	Body                      *map[string]string `yaml:"body,omitempty"`
	QueryParams               *[]S_QueryParam    `yaml:"query_params,omitempty"`
}

type S_Endpoint struct {
	URI            string           `yaml:"uri"`
	Method         string           `yaml:"method"`
	ResponseFormat string           `yaml:"response_format"`
	LimitExtract   *int             `yaml:"limit_extract,omitempty"`
	EndpointConfig S_EndpointConfig `yaml:"endpoint_config"`
}
