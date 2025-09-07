package endpoint

import (
	endpointconfig "extrator/internal/product/endpoint/endpointConfig"
)

type S_Endpoint struct {
	Name           string
	EndpointConfig endpointconfig.S_EndpointConfig
}
