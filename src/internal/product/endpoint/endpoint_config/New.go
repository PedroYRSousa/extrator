package endpoint_config

import "extrator/internal/product/endpoint/endpoint_config/client"

func New() *S_EndpointConfig {
	return &S_EndpointConfig{
		Method:      DEFAULT_METHOD,
		Headers:     &[]map[string]string{},
		QueryParams: &[]map[string]string{},
		Body:        nil,
		Client:      client.New(),
	}
}
