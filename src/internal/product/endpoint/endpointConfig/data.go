package endpointconfig

import (
	"extrator/internal/product/endpoint/endpointConfig/body"
	"net/http"
)

const (
	DEFAULT_METHOD = http.MethodGet
)

type S_EndpointConfig struct {
	Request  *http.Request
	Response *http.Response

	URL         string               `yaml:"url" validate:"required,url,printascii"`
	Method      string               `yaml:"_method" validate:"required,uppercase,oneof=GET POST"`
	Headers     *[]map[string]string `yaml:"_headers" validate:"required"`
	QueryParams *[]map[string]string `yaml:"_query_params" validate:"required"`
	Body        *body.S_Body         `yaml:"_body"`
}
