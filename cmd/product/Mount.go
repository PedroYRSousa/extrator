package product

import (
	"extrator/config"
)

func (product *S_Product) Mount(conf *config.S_Config) error {
	for index := range product.Endpoints {
		err := product.Endpoints[index].Mount(conf)
		if err != nil {
			return err
		}
	}

	return nil
}
