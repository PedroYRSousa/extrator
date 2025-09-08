package product

import (
	"errors"
	"extrator/internal/configs"
	"extrator/internal/product/endpoint"
	"extrator/pkg/validator"
	"fmt"
	"os"
)

func (p *S_Product) load() error {
	if p == nil {
		panic("Internal error | product.load")
	}

	endpoints, err := endpoint.Loads(p.Path)
	if err != nil {
		return err
	}

	p.Endpoints = endpoints

	if err := validator.New().Struct(p); err != nil {
		return errors.New(validator.FormatErrors(err, p))
	}

	return nil
}

func Loads(conf *configs.S_Configs) ([]*S_Product, error) {
	products := []*S_Product{}

	entriesProducts, err := os.ReadDir(conf.PathProducts)
	if err != nil {
		return nil, err
	}

	// Ler a pasta de configuração
	for _, entryProduct := range entriesProducts {
		if !entryProduct.IsDir() {
			continue
		}

		pathProduct := fmt.Sprintf("%s/%s", conf.PathProducts, entryProduct.Name())
		_product := New(entryProduct.Name(), pathProduct)
		err := _product.load()
		if err != nil {
			return nil, err
		}

		products = append(products, _product)
	}

	return products, nil
}
