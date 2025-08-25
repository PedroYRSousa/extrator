package product

import (
	"log"
	"net/http/httputil"
)

func (p *S_ProductEndpoint) Start() error {
	log.Println("Start", p.ProductName, p.Name)

	request, err := p.Mount()
	if err != nil {
		return err
	}

	data, err := httputil.DumpRequest(request, true)
	log.Println(string(data))
	log.Println(err)

	return nil
}
