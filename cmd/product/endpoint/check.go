package endpoint

func (endpoint *S_Endpoint) check() error {
	endpoint = endpoint.format()

	err := endpoint.validate()
	if err != nil {
		return err
	}

	err = endpoint.transform()
	if err != nil {
		return err
	}

	return nil
}
