package product

import (
	"fmt"
	"log"
	"net/http/httputil"
)

func (p *S_ProductEndpoint) Start() error {
	log.Println("Start", p.ProductName, p.Name)

	request, err := p.Mount()
	if err != nil {
		return err
	}

	data, _ := httputil.DumpRequest(request, true)
	fmt.Println(string(data))

	return nil
}
