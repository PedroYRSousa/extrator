package bearer

import endpointconfig "extrator/internal/product/endpoint/endpointConfig"

type S_Bearer struct {
	Name           string
	Prefix         string
	Token          *string
	IsDynamic      bool
	EndpointConfig *endpointconfig.S_EndpointConfig
}
