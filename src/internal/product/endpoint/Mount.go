package endpoint

func (e *S_Endpoint) Mount() error {
	err := e.EndpointConfig.Mount()
	if err != nil {
		return err
	}

	return nil
}
