package product

import (
	"errors"
)

func (c *S_ProductEndpoint) validate() error {
	if c.Name == "" {
		return errors.New("invalid config | Check name | name cannot be empty")
	}

	if c.Version == "" {
		return errors.New("invalid config | Check version | version cannot be empty")
	}

	if c.Description == "" {
		return errors.New("invalid config | Check description | description cannot be empty")
	}

	err := c.Endpoint.Validate()
	if err != nil {
		return err
	}

	err = c.Pagination.Validate()
	if err != nil {
		return err
	}

	err = c.Auth.Validate()
	if err != nil {
		return err
	}

	if len(c.Tables) <= 0 {
		return errors.New("invalid config | Check tables | tables cannot be empty")
	}
	for _, table := range c.Tables {
		err := table.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}
