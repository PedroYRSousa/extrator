package product

import (
	"extrator/internal/configs"
	"extrator/pkg/async"
	"log"
)

func (p *S_Product) Mount(conf *configs.S_Configs) {
	asyncEndpoints := async.AsyncLimiter(conf.EndpointsParallelExecutionCount)

	for _, endpoint := range p.Endpoints {
		_endpoint := endpoint

		asyncEndpoints.Go(func() {
			err := _endpoint.Mount()
			if err != nil {
				log.Println("Aborting endpoint: ", _endpoint.Path, err)
				_endpoint.ToAbort = true
				return
			}

			_endpoint.ToAbort = false
		})
	}

	asyncEndpoints.Wait()
}
