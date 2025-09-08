package endpointconfig

func New() *S_EndpointConfig {
	return &S_EndpointConfig{
		Method:      DEFAULT_METHOD,
		Headers:     &[]map[string]string{},
		QueryParams: &[]map[string]string{},
		Body:        nil,
	}
}
