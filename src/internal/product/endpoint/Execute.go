package endpoint

func (e *S_Endpoint) Execute() error {
	err := e.EndpointConfig.Execute()
	if err != nil {
		return err
	}

	return nil
}
