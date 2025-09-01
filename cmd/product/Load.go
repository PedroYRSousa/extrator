package product

import (
	"extrator/config"
	"extrator/modules"
	"extrator/product/endpoint"
	"os"
)

func (product *S_Product) _load() error {
	if product == nil {
		panic("TODO, melhorar | ERROR _load product")
	}

	entries, err := os.ReadDir(product.Path)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		path := modules.JoinPaths(product.Path, entry.Name())
		endpoint := endpoint.New()
		endpoint.Name = entry.Name()
		endpoint.Path = path

		err = modules.YamlToStruct(endpoint.Path, &endpoint)
		if err != nil {
			return err
		}

		product.Endpoints = append(product.Endpoints, endpoint)
	}

	return nil
}

func Load(conf config.S_Config) ([]S_Product, error) {
	products := []S_Product{}

	entries, err := os.ReadDir(conf.Products.Path)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		path := modules.JoinPaths(conf.Products.Path, entry.Name())
		product := New()
		product.Name = entry.Name()
		product.Path = path

		err = product._load()
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}
