package product

import (
	"extrator/internal/configs"
	"log"
)

func (p *S_Product) Mount(conf *configs.S_Configs) {
	for _, endpoint := range p.Endpoints {
		_endpoint := endpoint
		_endpoint.ToAbort = false

		err := _endpoint.Mount()
		if err != nil {
			log.Println("Aborting endpoint: ", _endpoint.Path, err)
			_endpoint.ToAbort = true
		}
	}
}
