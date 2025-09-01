package product

import (
	"github.com/go-playground/validator/v10"
)

func (product *S_Product) validate() error {
	if product == nil {
		panic("TODO, melhorar | ERROR validate product")
	}

	validator := validator.New()

	err := validator.Struct(product)
	if err != nil {
		return err
	}

	for _, endpoint := range product.Endpoints {
		err := validator.Struct(endpoint)
		if err != nil {
			return err
		}
	}

	return nil
}

func (product *S_Product) check() error {
	if product == nil {
		panic("TODO, melhorar | ERROR check product")
	}

	err := product.validate()
	if err != nil {
		return err
	}

	return nil
}
