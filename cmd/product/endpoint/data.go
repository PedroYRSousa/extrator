package endpoint

import "net/http"

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
	Name               string  `validate:"required,printascii"`
	Path               string  `validate:"required,file,printascii"`
	Version            string  `yaml:"version" validate:"required,printascii"`
	Description        string  `yaml:"description" validate:"required,printascii,min=15"`
	URL                string  `yaml:"url" validate:"required,printascii,url"`
	Method             *string `yaml:"_method" validate:"printascii"`
	ResponseFormat     *string `yaml:"_response_format" validate:"printascii"`
	ExtractionJSONPath *string `yaml:"_extraction_json_path" validate:"printascii"`
}
