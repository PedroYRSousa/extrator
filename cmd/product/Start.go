package product

import (
	"extrator/config"
	"extrator/modules"
	"fmt"
	"log"
)

func (product *S_Product) Start(conf *config.S_Config) error {
	endpointsAsyncCount := modules.AsyncLimiter(int(*conf.Products.EndpointPerProductAsyncCount))
	for _, endpoint := range product.Endpoints {
		_endpoint := endpoint
		endpointsAsyncCount.Go(func() {
			err := modules.ExecWithAttempts(*conf.Products.MaxAttempts, *conf.Products.DelayAttemptsInSeconds, func() error {
				err := _endpoint.Start(conf)
				if err != nil {
					return fmt.Errorf("ERROR | Start product (%s) | %+v", _endpoint.Name, err)
				}

				return nil
			})
			if err != nil {
				log.Println(err)
			}
		})

	}
	endpointsAsyncCount.Wait()

	return nil
}
