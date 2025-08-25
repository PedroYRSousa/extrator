package endpoint

func (e *S_Endpoint) Transform() error {
	err := e.EndpointConfig.Transform()
	if err != nil {
		return err
	}

	return nil
}
