package endpoint

import (
	"extrator/modules"
	"strings"
)

func (endpoint *S_Endpoint) format() *S_Endpoint {
	if endpoint == nil {
		panic("TODO, Melhora isso aqui")
	}

	if endpoint.Method == nil {
		endpoint.Method = new(string)
		*endpoint.Method = DEFAULT_METHOD
	} else {
		*endpoint.Method = strings.ToUpper(*endpoint.Method)
	}

	if endpoint.ResponseFormat == nil {
		endpoint.ResponseFormat = new(string)
		*endpoint.ResponseFormat = DEFAULT_RESPONSE_FORMAT
	} else {
		*endpoint.ResponseFormat = strings.ToLower(*endpoint.ResponseFormat)
	}

	if endpoint.ExtractionJSONPath == nil {
		endpoint.ExtractionJSONPath = new(string)
		*endpoint.ExtractionJSONPath = DEFAULT_EXTRACTION_JSON_PATH
	}

	if endpoint.Headers == nil {
		endpoint.Headers = new(map[string]string)
		*endpoint.Headers = map[string]string{}
	}

	if endpoint.QueryParams == nil {
		endpoint.QueryParams = new(map[string]string)
		*endpoint.QueryParams = map[string]string{}
	}

	modules.FormatString(endpoint)

	return endpoint
}
