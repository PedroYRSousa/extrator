package endpoint

import (
	"extrator/modules"
)

func (endpoint *S_Endpoint) format() *S_Endpoint {
	if endpoint == nil {
		panic("TODO, Melhora isso aqui")
	}

	modules.FormatString(endpoint)

	return endpoint
}
