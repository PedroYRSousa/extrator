package endpoint

import "extrator/product/endpoint/endpoint_config"

func New() S_Endpoint {
	return S_Endpoint{
		EndpointConfig: endpoint_config.New(),
	}
}
