package endpoint

import (
	"strings"
)

func (e *S_Endpoint) Format() {
	e.URI = strings.TrimSpace(strings.ToLower(e.URI))
	e.Method = strings.TrimSpace(strings.ToUpper(e.Method))
	e.ResponseFormat = endpointResponseFormat(strings.TrimSpace(strings.ToLower(string(e.ResponseFormat))))
	e.EndpointConfig.format()
}

func (r *S_Retry) format() {
	if r.BackoffFactor == nil {
		r.BackoffFactor = new(int)
		*r.BackoffFactor = 0
	}
}

func (ec *S_EndpointConfig) format() {
	ec.Retry.format()
}
