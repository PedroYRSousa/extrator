package product

import (
	"strings"
)

func (c *S_ProductEndpoint) format() {
	c.Name = strings.TrimSpace(c.Name)
	c.Version = strings.TrimSpace(c.Version)
	c.Description = strings.TrimSpace(c.Description)

	c.Endpoint.Format()
	c.Pagination.Format()
	c.Auth.Format()
	for index := range c.Tables {
		c.Tables[index].Format()
	}
}
