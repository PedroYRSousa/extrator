package product

func (c *S_ProductEndpoint) transform() error {
	c.Endpoint.Transform()
	c.Pagination.Transform()
	c.Auth.Transform()

	for _, table := range c.Tables {
		err := table.Transform()
		if err != nil {
			return err
		}
	}

	return nil
}
