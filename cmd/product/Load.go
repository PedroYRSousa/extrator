package product

import (
	"extrator/config"
	"extrator/modules"
	"extrator/product/endpoint"
	"os"
)

func Load(conf *config.S_Config) ([]S_Product, error) {
	products := []S_Product{}

	entries, err := os.ReadDir(conf.Products.Path)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			product := S_Product{
				Name: entry.Name(),
				Path: modules.JoinPaths(conf.Products.Path, entry.Name()),
			}

			product.Endpoints, err = endpoint.Load(conf, product.Path)
			if err != nil {
				return nil, err
			}

			err = product.check()
			if err != nil {
				return nil, err
			}

			products = append(products, product)
		}
	}

	return products, nil
}
