package endpoint

import (
	"extrator/product/endpoint/body"
	"net/http"
)

const (
	DEFAULT_METHOD               = http.MethodGet
	DEFAULT_RESPONSE_FORMAT      = "json"
	DEFAULT_EXTRACTION_JSON_PATH = "$"
)

var (
	METHODS_ALLOWED         = []string{http.MethodGet, http.MethodPost}
	RESPONSE_FORMAT_ALLOWED = []string{"json"}
)

type S_Endpoint struct {
	Request            *http.Request
	Name               string             `validate:"required,printascii"`
	Path               string             `validate:"required,file,printascii"`
	Version            string             `yaml:"version" validate:"required,printascii"`
	Description        string             `yaml:"description" validate:"required,printascii,min=15"`
	URL                string             `yaml:"url" validate:"required,printascii,url"`
	Method             *string            `yaml:"_method,omitempty" validate:"printascii"`
	ResponseFormat     *string            `yaml:"_response_format,omitempty" validate:"printascii"`
	ExtractionJSONPath *string            `yaml:"_extraction_json_path,omitempty" validate:"printascii"`
	QueryParams        *map[string]string `yaml:"_query_params,omitempty" validate:"printascii"`
	Headers            *map[string]string `yaml:"_headers,omitempty" validate:"printascii"`
	Body               *body.S_Body       `yaml:"_body,omitempty"`
}
