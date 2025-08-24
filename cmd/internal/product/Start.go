package product

import "log"

func (c *S_ProductEndpoint) Start() {
	log.Println("Start", *c)
}
