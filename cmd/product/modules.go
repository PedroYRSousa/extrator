package product

import (
	"extrator/modules"

	"github.com/go-playground/validator/v10"
)

func (product *S_Product) format() {
	if product == nil {
		panic("TODO, melhorar | ERROR format product")
	}

	modules.FormatString(product)
}

func (product *S_Product) transform() error {
	if product == nil {
		panic("TODO, melhorar | ERROR transform product")
	}

	err := modules.TransformString(product)
	if err != nil {
		return err
	}

	return nil
}

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

	product.format()

	err := product.validate()
	if err != nil {
		return err
	}

	err = product.transform()
	if err != nil {
		return err
	}

	return nil
}
