package endpoint

import (
	"extrator/internal/product/endpoint_config"
)

type endpointResponseFormat string

const (
	ENDPOINT_RESPONSE_FORMAT_JSON endpointResponseFormat = "json"

	ENDPOINT_LIMIT_EXTRACT_DEFAULT = -1
)

var (
	// Por hora somente esses formatos são suportados para extração de dados
	RESPONSES_FORMAT_AVAILABLE = []endpointResponseFormat{ENDPOINT_RESPONSE_FORMAT_JSON}
)

type S_Endpoint struct {
	URI            string                 `yaml:"uri"`
	ResponseFormat endpointResponseFormat `yaml:"response_format"`

	// Opcionais
	LimitExtract   *int                              `yaml:"limit_extract,omitempty"`
	EndpointConfig *endpoint_config.S_EndpointConfig `yaml:"endpoint_config,omitempty"`
}
