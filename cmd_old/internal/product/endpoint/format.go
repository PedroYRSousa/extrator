package endpoint

import (
	"strings"
)

func (e *S_Endpoint) format() {
	if e == nil {
		panic("Endpoint is required")
	}

	e.URI = strings.TrimSpace(strings.ToLower(e.URI))
	e.ResponseFormat = endpointResponseFormat(strings.TrimSpace(strings.ToLower(string(e.ResponseFormat))))

	if e.LimitExtract == nil {
		e.LimitExtract = new(int)
		*e.LimitExtract = ENDPOINT_LIMIT_EXTRACT_DEFAULT
	}
}
