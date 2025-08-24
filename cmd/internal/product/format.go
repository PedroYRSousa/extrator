package product

import (
	"strings"
)

func (c *S_ProductEndpoint) format() {
	c.Name = strings.TrimSpace(c.Name)
	c.Version = strings.TrimSpace(c.Version)
	c.Description = strings.TrimSpace(c.Description)

	// c.Auth.Format()
}
