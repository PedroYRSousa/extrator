package endpoint

import endpointconfig "extrator/internal/product/endpoint_config"

type endpointResponseFormat string

const (
	ENDPOINT_RESPONSE_FORMAT_JSON endpointResponseFormat = "json"

	ENDPOINT_LIMIT_EXTRACT_DEFAULT = -1
)

var (
	responsesFormatAvailable = []endpointResponseFormat{ENDPOINT_RESPONSE_FORMAT_JSON}
)

type S_Endpoint struct {
	URI            string                 `yaml:"uri"`
	Method         string                 `yaml:"method"`
	ResponseFormat endpointResponseFormat `yaml:"response_format"`

	// Opcionais
	EndpointConfig *endpointconfig.S_EndpointConfig `yaml:"endpoint_config,omitempty"`
	LimitExtract   *int                             `yaml:"limit_extract,omitempty"`
}
