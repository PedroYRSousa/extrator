package endpoint_config

func New() S_EndpointConfig {
	method := METHOD_DEFAULT
	responseFormat := RESPONSE_FORMAT_DEFAULT

	return S_EndpointConfig{
		Method:         &method,
		ResponseFormat: &responseFormat,
		QueryParams:    &map[string]string{},
		Headers:        &map[string]string{},
		Body:           nil,
	}
}
