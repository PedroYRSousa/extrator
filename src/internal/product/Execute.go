package product

import (
	"extrator/internal/configs"
	"extrator/pkg/async"
	"log"
)

func (p *S_Product) Execute(conf *configs.S_Configs) {
	asyncEndpoints := async.AsyncLimiter(conf.EndpointsParallelExecutionCount)

	for _, endpoint := range p.Endpoints {
		_endpoint := endpoint
		if _endpoint.ToAbort {
			continue
		}

		asyncEndpoints.Go(func() {
			err := _endpoint.Execute()
			log.Println("Aborting endpoint: ", _endpoint.Path, err)
			_endpoint.ToAbort = true
		})
	}

	asyncEndpoints.Wait()
}
