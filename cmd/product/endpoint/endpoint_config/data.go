package endpoint_config

import (
	"extrator/product/endpoint/endpoint_config/body"
	"net/http"
)

const (
	RESPONSE_FORMAT_JSON     = "json"
	RESPONSE_FORMAT_XML      = "xml"
	RESPONSE_FORMAT_ZIP_JSON = "zip/json"
	RESPONSE_FORMAT_ZIP_XML  = "zip/xml"

	METHOD_GET  = http.MethodGet
	METHOD_POST = http.MethodPost

	RESPONSE_FORMAT_DEFAULT = RESPONSE_FORMAT_JSON
	METHOD_DEFAULT          = METHOD_GET
)

var (
	METHOD_ALLOWED          = []string{METHOD_GET, METHOD_POST}
	RESPONSE_FORMAT_ALLOWED = []string{RESPONSE_FORMAT_JSON, RESPONSE_FORMAT_XML, RESPONSE_FORMAT_ZIP_JSON, RESPONSE_FORMAT_ZIP_XML}
)

type S_EndpointConfig struct {
	URL string `yaml:"url" validate:"required,url,printascii"`

	Method         *string            `yaml:"_method,omitempty" validate:"uppercase,printascii"`
	ResponseFormat *string            `yaml:"_response_format,omitempty" validate:"lowercase,printascii"`
	QueryParams    *map[string]string `yaml:"_query_params,omitempty" validate:"printascii"`
	Headers        *map[string]string `yaml:"_headers,omitempty" validate:"printascii"`
	Body           *body.S_Body       `yaml:"_body,omitempty" validate:"printascii"`
}
