package product

func (c *S_ProductEndpoint) transform() error {
	err := c.Endpoint.Transform()
	if err != nil {
		return err
	}

	err = c.Pagination.Transform()
	if err != nil {
		return err
	}

	err = c.Auth.Transform()
	if err != nil {
		return err
	}

	for _, table := range c.Tables {
		err := table.Transform()
		if err != nil {
			return err
		}
	}

	return nil
}
